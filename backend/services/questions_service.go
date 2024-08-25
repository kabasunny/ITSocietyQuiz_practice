package services

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"math/rand"
	"time"
)

type IQuestionsService interface {
	FindAll() (*[]models.Questions, error)
	FindById(QuestionsId uint) (*models.Questions, error)
	Create(createQuestionsInput dto.CreateQuestionsInput) (*models.Questions, error)
	Update(QuestionsId uint, updateQuestionsInput dto.UpdateQuestionsInput) (*models.Questions, error)
	Delete(QuestionsId uint) error
	GetOneDaysQuiz(NunmberOfQuestions uint) (*[]models.Questions, error) // 1日分のクイズを取得する
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

func (s *QuestionsService) GetOneDaysQuiz(NumberOfQuestions uint) (*[]models.Questions, error) {
	// 問題データ数を取得
	totalQuestions, err := s.repository.Count()
	if err != nil {
		return nil, err
	}

	// ランダムにIDを生成してクイズデータを取得
	selectedQuestions := make([]models.Questions, 0, NumberOfQuestions)
	usedIDs := make(map[uint]bool) // 使用済みの質問IDを追跡するためのマップ、デフォルトはfalse
	// 将来、忘却アルゴリズムを組み合わせる際、ランダムする前に質問IDを格納

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

	return &selectedQuestions, nil // selectedQuestionsの参照アドレス値(ポインタ)を返す
}
