package services

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"time"

	"github.com/lib/pq"
)

type IAdminsService interface {
	FindAllQuestions() (*[]dto.AdmQuizData, error) // 修正
	FindQuestionsById(QuestionsId uint) (*models.Questions, error)
	// CreateQuestions(createQuestionsInput dto.CreateQuestionsInput) (*models.Questions, error)
	UpdateQuestions(QuestionsId uint, updateQuestionsInput dto.UpdateQuestionsInput) (*dto.UpdateQuestionsOutput, error)
	DeleteQuestions(QuestionsId uint) error
	ProcessCSVData(filepath string) error // 追加
}

type AdminsService struct {
	repository repositories.IAdminsRepository
}

func NewAdminsService(repository repositories.IAdminsRepository) IAdminsService {
	return &AdminsService{repository: repository}
}

func (s *AdminsService) FindAllQuestions() (*[]dto.AdmQuizData, error) { // 修正
	questions, err := s.repository.FindAllQuestions()
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
			UpdatedAt:      question.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &quizData, nil
}

func (s *AdminsService) FindQuestionsById(QuestionsId uint) (*models.Questions, error) {
	return s.repository.FindQuestionsById(QuestionsId)
}

// func (s *AdminsService) CreateQuestions(createQuestionsInput dto.CreateQuestionsInput) (*models.Questions, error) {
// 	newQuestions := models.Questions{
// 		Question:   createQuestionsInput.Question,
// 		Options:    createQuestionsInput.Options,
// 		Supplement: createQuestionsInput.Supplement,
// 		Difficulty: createQuestionsInput.Difficulty,
// 	}
// 	return s.repository.Create(newQuestions)
// }

func (s *AdminsService) UpdateQuestions(QuestionsId uint, updateQuestionsInput dto.UpdateQuestionsInput) (*dto.UpdateQuestionsOutput, error) {
	targetQuestions, err := s.FindQuestionsById(QuestionsId)
	if err != nil {
		return nil, err
	}
	if updateQuestionsInput.Question != nil {
		targetQuestions.UserQuestionID = updateQuestionsInput.UserQuestionID // ポインタ型
	}
	if updateQuestionsInput.Question != nil {
		targetQuestions.Question = *updateQuestionsInput.Question
	}
	if updateQuestionsInput.Options != nil {
		targetQuestions.Options = pq.StringArray(*updateQuestionsInput.Options) // 一応キャスト　pq.StringArrayは[]stringのエイリアスで、PostgreSQLのtext[]型と直接互換性がある
	}
	if updateQuestionsInput.Supplement != nil {
		targetQuestions.Supplement = *updateQuestionsInput.Supplement
	}
	if updateQuestionsInput.Difficulty != nil {
		targetQuestions.Difficulty = *updateQuestionsInput.Difficulty
	}
	updatedQuestions, err := s.repository.UpdateQuestions(targetQuestions)
	if err != nil {
		return nil, err
	}

	// モデル構造体をDTO構造体に変換
	updateQuestionsOutput := &dto.UpdateQuestionsOutput{
		ID:             updatedQuestions.ID,
		UserQuestionID: updatedQuestions.UserQuestionID,
		Question:       updatedQuestions.Question,
		Options:        updatedQuestions.Options,
		Supplement:     updatedQuestions.Supplement,
		Difficulty:     updatedQuestions.Difficulty,
		CreatedAt:      updatedQuestions.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      updatedQuestions.UpdatedAt.Format(time.RFC3339),
	}

	return updateQuestionsOutput, nil
}

func (s *AdminsService) DeleteQuestions(QuestionsId uint) error {
	return s.repository.DeleteQuestions(QuestionsId)
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
