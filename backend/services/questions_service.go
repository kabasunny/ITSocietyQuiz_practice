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
	GetOneDaysQuiz(NunmberOfQuestions uint) (*[]models.Questions, error)
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
	// 総問題数を取得
	totalQuestions, err := s.repository.Count()
	if err != nil {
		return nil, err
	}

	// ランダムにIDを生成してクイズデータを取得
	selectedQuestions := make([]models.Questions, 0, NumberOfQuestions)
	usedIDs := make(map[uint]bool)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for uint(len(selectedQuestions)) < NumberOfQuestions {
		randomID := uint(r.Intn(int(totalQuestions)) + 1)
		if !usedIDs[randomID] {
			question, err := s.repository.FindById(randomID)
			if err == nil {
				selectedQuestions = append(selectedQuestions, *question)
				usedIDs[randomID] = true
			}
		}
	}

	return &selectedQuestions, nil
}
