// 初期テスト用

package services

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
)

type IQuizDataService interface {
	FindAll() (*[]models.QuizData, error)
	FindById(QuizDataId uint) (*models.QuizData, error)
	Create(createQuizDataInuput dto.CreateQuizDataInput) (*models.QuizData, error)
	Update(QuizDataId uint, updateQuizDataInput dto.UpdateQuizDataInput) (*models.QuizData, error)
	Delete(QuizDataId uint) error
}

type QuizDataService struct {
	repository repositories.IQuizDataRepository
}

func NewQuizDataService(repository repositories.IQuizDataRepository) IQuizDataService {
	return &QuizDataService{repository: repository}
}

func (s *QuizDataService) FindAll() (*[]models.QuizData, error) {
	return s.repository.FindAll()
}

func (s *QuizDataService) FindById(QuizDataId uint) (*models.QuizData, error) {
	return s.repository.FindById(QuizDataId)
}

func (s *QuizDataService) Create(createQuizDataInuput dto.CreateQuizDataInput) (*models.QuizData, error) {
	newQuizData := models.QuizData{
		Question:   createQuizDataInuput.Question,
		Options:    createQuizDataInuput.Options,
		Correct:    createQuizDataInuput.Correct,
		Supplement: createQuizDataInuput.Supplement,
	}
	return s.repository.Create(newQuizData)
}

func (s *QuizDataService) Update(QuizDataId uint, updateQuizDataInput dto.UpdateQuizDataInput) (*models.QuizData, error) {
	targetQuizData, err := s.FindById(QuizDataId)
	if err != nil {
		return nil, err
	}
	if updateQuizDataInput.Question != nil {
		targetQuizData.Question = *updateQuizDataInput.Question
	}
	if updateQuizDataInput.Options != nil {
		targetQuizData.Options = *updateQuizDataInput.Options
	}
	if updateQuizDataInput.Correct != nil {
		targetQuizData.Correct = *updateQuizDataInput.Correct
	}
	if updateQuizDataInput.Supplement != nil {
		targetQuizData.Supplement = *updateQuizDataInput.Supplement
	}
	return s.repository.Update(*targetQuizData)
}

func (s *QuizDataService) Delete(QuizDataId uint) error {
	return s.repository.Delete(QuizDataId)
}
