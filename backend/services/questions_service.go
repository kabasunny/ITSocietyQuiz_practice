package services

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"backend/utils" // ValidateToken(tokenString string) (string, bool, error)
	"log"
)

type IQuestionsService interface {
	FindAll() (*[]models.Questions, error)
	FindById(QuestionsId uint) (*models.Questions, error)
	Create(createQuestionsInput dto.CreateQuestionsInput) (*models.Questions, error)
	Update(QuestionsId uint, updateQuestionsInput dto.UpdateQuestionsInput) (*models.Questions, error)
	Delete(QuestionsId uint) error
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

func (s *QuestionsService) Create(createQuestionsInput dto.CreateQuestionsInput) (*models.Questions, error) {
	newQuestions := models.Questions{
		Question:   createQuestionsInput.Question,
		Options:    createQuestionsInput.Options,
		Supplement: createQuestionsInput.Supplement,
		Difficulty: createQuestionsInput.Difficulty,
	}
	return s.repository.Create(newQuestions)
}

func (s *QuestionsService) Update(QuestionsId uint, updateQuestionsInput dto.UpdateQuestionsInput) (*models.Questions, error) {
	targetQuestions, err := s.FindById(QuestionsId)
	if err != nil {
		return nil, err
	}
	if updateQuestionsInput.Question != nil {
		targetQuestions.Question = *updateQuestionsInput.Question
	}
	if updateQuestionsInput.Options != nil {
		targetQuestions.Options = *updateQuestionsInput.Options
	}
	if updateQuestionsInput.Supplement != nil {
		targetQuestions.Supplement = *updateQuestionsInput.Supplement
	}
	if updateQuestionsInput.Difficulty != nil {
		targetQuestions.Difficulty = *updateQuestionsInput.Difficulty
	}
	return s.repository.Update(*targetQuestions)
}

func (s *QuestionsService) Delete(QuestionsId uint) error {
	return s.repository.Delete(QuestionsId)
}

func (s *QuestionsService) GetOneDaysQuiz(tokenString string, todaysCount uint) (*[]dto.QuizData, bool, error) {
	// 2秒間の遅延。フロントの画面の遷移確認用
	// time.Sleep(1 * time.Second)

	// トークンの検証とEmpIDの抽出
	empID, valid, err := utils.ValidateToken(tokenString)
	if err != nil || !valid {
		return nil, false, err
	}

	todaysCount = 0 // テスト時の制限解除。当日分の回数がリセットされ制限がかからなくなる

	necessaryQuestions := s.dailyQuestionCount - todaysCount // 日に必要な問題数

	// 日のノルマが達成されたかどうかを確認
	if int(necessaryQuestions) <= 0 {
		return nil, true, nil
	}

	// SQLクエリを読み込む
	query, err := utils.LoadSQLFile("services/queries/select_questions_excluding_streak_3.sql")
	if err != nil {
		log.Fatalf("Failed to load SQL file: %v", err)
	}

	// 1. answersテーブルからデータを取得
	questionLimit := s.dailyQuestionCount * s.pastDaysRange // 検索レコード数 = 一日の問題数 × 検索範囲日数
	selectedQuestions, err := s.repository.GetTopQuestionsByEmpID(query, empID, questionLimit, necessaryQuestions)
	if err != nil {
		return nil, false, err
	}

	// 2. 不足分のquestion_idを補完
	if len(selectedQuestions) < int(necessaryQuestions) {
		currentQID, err := s.repository.GetCurrentQIDByEmpID(empID)
		if err != nil {
			return nil, false, err
		}

		for len(selectedQuestions) < int(necessaryQuestions) {
			currentQID++
			selectedQuestions = append(selectedQuestions, currentQID)
		}
	}

	// 4. selectedQuestionsから詳細データを取得
	quizDetails, err := s.repository.GetQuestionDetails(selectedQuestions)
	if err != nil {
		return nil, false, err
	}

	// DTOに変換
	quizData := make([]dto.QuizData, len(quizDetails))
	for i, question := range quizDetails {
		quizData[i] = dto.QuizData{
			ID:         question.ID,
			Question:   question.Question,
			Options:    question.Options,
			Supplement: question.Supplement,
		}
	}

	return &quizData, false, nil
}
