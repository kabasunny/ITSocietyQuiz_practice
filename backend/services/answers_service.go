package services

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"backend/utils" // ValidateToken(tokenString string) (string, bool, error)
	"errors"

	"gorm.io/gorm"
)

type IAnswersService interface {
	SaveAnswers(inputs []dto.AnswersInput, tokenString string) error
	// ValidateToken(tokenString string) (string, bool, error) // トークンの検証メソッド　utilsにて共通化処理とする
}

type AnswersService struct {
	repository repositories.IAnswersRepository
}

func NewAnswersService(repository repositories.IAnswersRepository) IAnswersService {
	return &AnswersService{repository: repository}
}

func (s *AnswersService) SaveAnswers(inputs []dto.AnswersInput, tokenString string) error {
	empID, valid, err := utils.ValidateToken(tokenString)
	if err != nil || !valid {
		return err
	}

	var answersBatch []models.Answers

	for _, input := range inputs {
		latestAnswer, err := s.repository.GetLatestAnswer(empID, input.QuestionID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}

		streakCount := uint(0)
		if input.AnswerID == 0 {
			streakCount = latestAnswer.StreakCount + 1
		} else {
			streakCount = 0
		}

		answers := models.Answers{
			EmpID:       empID,
			QuestionID:  input.QuestionID,
			AnswerID:    input.AnswerID,
			StreakCount: streakCount,
		}

		answersBatch = append(answersBatch, answers)
	}

	err = s.repository.CreateAnswersBatch(answersBatch)
	if err != nil {
		return err
	}

	return nil
}
