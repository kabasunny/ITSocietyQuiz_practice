package services

import (
	"backend/src/dto"
	"backend/src/models"
	"backend/src/repositories"
	"backend/src/utils" // ValidateToken(tokenString string) (string, bool, error)
	"fmt"
	"log"
)

type IQuestionsService interface {
	FindAll() (*[]models.Questions, error)
	FindById(QuestionsId uint) (*models.Questions, error)
	GetOneDaysQuiz(tokenString string, todaysCount uint) (*[]dto.QuizData, bool, error) // 1日分のクイズを取得する
}

const (
	DailyQuestionCount = 5  // 一日当たりの問題数
	PastDaysRange      = 21 // 検索範囲の何日分
)

type QuestionsService struct {
	repository         repositories.IQuestionsRepository
	dailyQuestionCount uint // 一日当たりの問題数
	pastDaysRange      uint // 検索範囲の何日分
}

func NewQuestionsService(repository repositories.IQuestionsRepository) IQuestionsService {
	return &QuestionsService{repository: repository,
		dailyQuestionCount: DailyQuestionCount,
		pastDaysRange:      PastDaysRange,
	}
}

func (s *QuestionsService) FindAll() (*[]models.Questions, error) {
	return s.repository.FindAll()
}

func (s *QuestionsService) FindById(QuestionsId uint) (*models.Questions, error) {
	return s.repository.FindById(QuestionsId)
}

func (s *QuestionsService) GetOneDaysQuiz(tokenString string, todaysCount uint) (*[]dto.QuizData, bool, error) {
	// 2秒間の遅延。フロントの画面の遷移確認用
	// time.Sleep(1 * time.Second)

	// トークンの検証とEmpIDの抽出
	empID, valid, err := utils.ValidateToken(tokenString)
	if err != nil || !valid {
		fmt.Println("Token validation failed or invalid token")
		return nil, false, err
	}
	fmt.Println("Token validated successfully, EmpID:", empID)

	// todaysCount = 0 // テスト時の制限解除。当日分の回数がリセットされ制限がかからなくなる

	necessaryQuestions := s.dailyQuestionCount - todaysCount // 日に必要な問題数
	fmt.Println("Necessary questions for today:", necessaryQuestions)

	// 日のノルマが達成されたかどうかを確認
	if int(necessaryQuestions) <= 0 {
		fmt.Println("Daily quota already achieved")
		return nil, true, nil
	}

	// SQLクエリを読み込む
	query, err := utils.LoadSQLFile("src/services/queries/select_questions_excluding_streak_3.sql")
	if err != nil {
		log.Fatalf("Failed to load SQL file: %v", err)
	}
	fmt.Println("SQL query loaded successfully")

	// 1. answersテーブルからデータを取得
	questionLimit := s.dailyQuestionCount * s.pastDaysRange // 検索レコード数 = 一日の問題数 × 検索範囲日数
	selectedQuestions, err := s.repository.GetTopQuestionsByEmpID(query, empID, questionLimit, necessaryQuestions)
	if err != nil {
		fmt.Println("Failed to get top questions by EmpID:", err)
		return nil, false, err
	}
	fmt.Println("Top questions retrieved, count:", len(selectedQuestions))

	// 2. 不足分のquestion_idを補完
	if len(selectedQuestions) < int(necessaryQuestions) {
		fmt.Println("Starting to supplement missing question IDs")

		currentQID, err := s.repository.GetCurrentQIDByEmpID(empID)
		if err != nil {
			fmt.Println("Failed to get current QID by EmpID:", err)
			return nil, false, err
		}
		fmt.Println("Current QID:", currentQID)

		maxTries := 1000 // 確認する最大ID数（必要に応じて調整）これがないと無限ループになる
		tries := 0

		for len(selectedQuestions) < int(necessaryQuestions) && tries < maxTries {
			currentQID++
			tries++
			fmt.Println("Checking if question ID exists:", currentQID)

			// currentQIDが存在するか確認し、存在する場合はselectedQuestionsに追加
			exists, err := s.repository.ExistsById(currentQID)
			if err != nil {
				fmt.Println("Failed to check existence by ID:", err)
				return nil, false, err
			}
			if exists {
				selectedQuestions = append(selectedQuestions, currentQID)
				fmt.Println("Added question ID:", currentQID)
			} else {
				fmt.Println("Question ID does not exist:", currentQID)
			}
		}

		// 以下のif文に引っ掛かる時は、通常、すべての問題をクリアしたことになる
		// ただし、あるタイミングで問題を消したときにもおこりうる
		if len(selectedQuestions) < int(necessaryQuestions) {
			fmt.Println("Unable to supplement enough question IDs after", tries, "tries")
			return nil, false, fmt.Errorf("unable to supplement enough question IDs")
		}

		fmt.Println("Supplemented question IDs, total count:", len(selectedQuestions))
	}

	// 4. selectedQuestionsから詳細データを取得
	quizDetails, err := s.repository.GetQuestionDetails(selectedQuestions)
	if err != nil {
		fmt.Println("Failed to get question details:", err)
		return nil, false, err
	}
	fmt.Println("Question details retrieved, count:", len(quizDetails))

	// DTOに変換
	quizData := make([]dto.QuizData, len(quizDetails))
	for i, question := range quizDetails {
		quizData[i] = dto.QuizData{
			ID:         question.ID,
			Question:   question.Question,
			Options:    question.Options,
			Supplement: question.Supplement,
		}
		fmt.Println("Question ID:", question.ID, "has been added to the quiz data")
	}

	return &quizData, false, nil
}
