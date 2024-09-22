package services

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"time"
)

type IAdminsService interface {
	FindAll() (*[]dto.AdmQuizData, error) // 修正
	FindById(QuestionsId uint) (*models.Questions, error)
	Create(createQuestionsInput dto.CreateQuestionsInput) (*models.Questions, error)
	Update(QuestionsId uint, updateQuestionsInput dto.UpdateQuestionsInput) (*models.Questions, error)
	Delete(QuestionsId uint) error
	ProcessCSVData(filepath string) error // 追加
}

type AdminsService struct {
	repository repositories.IAdminsRepository
}

func NewAdminsService(repository repositories.IAdminsRepository) IAdminsService {
	return &AdminsService{repository: repository}
}

func (s *AdminsService) FindAll() (*[]dto.AdmQuizData, error) { // 修正
	questions, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	var quizData []dto.AdmQuizData
	for _, question := range *questions {
		quizData = append(quizData, dto.AdmQuizData{
			ID:             question.ID,
			UserQuestionID: question.UserQuestionID,
			Question:       question.Question,
			Options:        question.Options,
			Supplement:     question.Supplement,
			Difficulty:     question.Difficulty,
			CreatedAt:      question.CreatedAt.Format(time.RFC3339),
		})
	}

	return &quizData, nil
}

func (s *AdminsService) FindById(QuestionsId uint) (*models.Questions, error) {
	return s.repository.FindById(QuestionsId)
}

func (s *AdminsService) Create(createQuestionsInput dto.CreateQuestionsInput) (*models.Questions, error) {
	newQuestions := models.Questions{
		Question:   createQuestionsInput.Question,
		Options:    createQuestionsInput.Options,
		Supplement: createQuestionsInput.Supplement,
		Difficulty: createQuestionsInput.Difficulty,
	}
	return s.repository.Create(newQuestions)
}

func (s *AdminsService) Update(QuestionsId uint, updateQuestionsInput dto.UpdateQuestionsInput) (*models.Questions, error) {
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

func (s *AdminsService) Delete(QuestionsId uint) error {
	return s.repository.Delete(QuestionsId)
}

func (s *AdminsService) ProcessCSVData(filePath string) error {
	data, err := utils.ParseCSV(filePath)
	if err != nil {
		return err
	}

	// データをリポジトリに直接渡す
	if err := s.repository.CreateQuestionsBatch(data); err != nil {
		return err
	}

	return nil
}
