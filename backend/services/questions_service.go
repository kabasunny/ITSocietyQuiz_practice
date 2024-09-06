package services

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"backend/utils" // ValidateToken(tokenString string) (string, bool, error)
	"fmt"
	"math/rand"
	"time"
)

type IQuestionsService interface {
	FindAll() (*[]models.Questions, error)
	FindById(QuestionsId uint) (*models.Questions, error)
	Create(createQuestionsInput dto.CreateQuestionsInput) (*models.Questions, error)
	Update(QuestionsId uint, updateQuestionsInput dto.UpdateQuestionsInput) (*models.Questions, error)
	Delete(QuestionsId uint) error
	GetOneDaysQuiz(tokenString string) (*[]dto.QuizData, error) // 1日分のクイズを取得する
}

type QuestionsService struct {
	repository repositories.IQuestionsRepository
}

func NewQuestionsService(repository repositories.IQuestionsRepository) IQuestionsService {
	return &QuestionsService{repository: repository}
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
		Difficulty: createQuestionsInput.Difficulty, // 追加
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
		targetQuestions.Difficulty = *updateQuestionsInput.Difficulty // 追加
	}
	return s.repository.Update(*targetQuestions)
}

func (s *QuestionsService) Delete(QuestionsId uint) error {
	return s.repository.Delete(QuestionsId)
}

func (s *QuestionsService) GetOneDaysQuiz(tokenString string) (*[]dto.QuizData, error) {
	// トークンの検証とEmpIDの抽出
	empID, valid, err := utils.ValidateToken(tokenString)
	if err != nil || !valid { // nilでないか、trueでない場合
		return nil, err
	}
	fmt.Print(empID) //ダミー、すぐ消す

	NumberOfQuestions := uint(5) // 1日あたりの問題数を入力

	// 問題データ数を取得
	totalQuestions, err := s.repository.Count()
	if err != nil {
		return nil, err
	}

	// 仮実装として、ランダムにIDを生成してクイズデータを取得
	selectedQuestions := make([]models.Questions, 0, NumberOfQuestions) // スライスの初期化
	usedIDs := make(map[uint]bool)                                      // 使用済みの質問IDを追跡するためのマップ、デフォルトはfalse

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 現在の時刻から、ランダム数生成器を初期化

	for uint(len(selectedQuestions)) < NumberOfQuestions { // selectedQuestionsの長さがNumberOfQuestionsに達するまでループ
		randomID := uint(r.Intn(int(totalQuestions)) + 1) // 0からtotalQuestions - 1までのランダムな整数(int)を生成し、1を足す
		if !usedIDs[randomID] {                           // 質問IDが未使用であれば、以下を行う
			question, err := s.repository.FindById(randomID) // questionは、ポインタ型
			if err == nil {
				selectedQuestions = append(selectedQuestions, *question) // *questionは、値を取り出して渡す
				usedIDs[randomID] = true
			}
		}
	}

	// DTOに変換
	quizData := make([]dto.QuizData, len(selectedQuestions))
	for i, question := range selectedQuestions {
		quizData[i] = dto.QuizData{
			ID:         question.ID,
			Question:   question.Question,
			Options:    question.Options,
			Supplement: question.Supplement,
		}
	}

	return &quizData, nil // quizDataの参照アドレス値(ポインタ)を返す
}

// // トークンの検証メソッド utilsにて共通化処理とする
// func (s *QuestionsService) ValidateToken(tokenString string) (string, bool, error) {
// 	// Bearer トークンの形式を確認
// 	if !strings.HasPrefix(tokenString, "Bearer ") {
// 		return "", false, fmt.Errorf("invalid token format")
// 	}

// 	// Bearer プレフィックスを取り除く
// 	tokenString = strings.TrimPrefix(tokenString, "Bearer ")

// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok { // jwt.SigningMethodHS256 で生成されたトークンを、HMAC 系列で検証
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}
// 		return []byte(os.Getenv("SECRET_KEY")), nil
// 	})
// 	if err != nil {
// 		return "", false, fmt.Errorf("failed to parse token: %v", err)
// 	}

// 	// トークンのクレームを検証
// 	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
// 		// クレームの内容を確認
// 		sub, subOk := claims["sub"].(string)
// 		exp, expOk := claims["exp"].(float64)

// 		if !subOk || !expOk {
// 			return "", false, fmt.Errorf("invalid token claims")
// 		}

// 		// 有効期限の確認
// 		if time.Unix(int64(exp), 0).Before(time.Now()) {
// 			return "", false, fmt.Errorf("token has expired")
// 		}
// 		return sub, true, nil
// 	}
// 	return "", false, fmt.Errorf("invalid token")
// }
