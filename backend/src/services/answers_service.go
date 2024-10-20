package services

import (
	"backend/src/dto"
	"backend/src/models"
	"backend/src/repositories"
	"backend/src/utils" // ValidateToken(tokenString string) (string, bool, error)
	"errors"
	"fmt"

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

	// ブラウザ更新時に重複したデータを格納しないための実装
	// 	// 「日付が本日」かつ「QuestionIDが一致」するレコードが存在する場合は、[]dto.AnswersInputの中から当該QuestionIDを持つインスタンスを削除
	// 	「日付が本日」かつ「QuestionIDが一致」するレコードが存在する場合は、[]dto.AnswersInputの中から当該QuestionIDを持つインスタンスを除外する処理をAPIに追加しようと思います

	// リポジトリに新規のメソッドを追加し
	// サービス側から、[]dto.AnswersInputからQuestionIDを取り出し、リポジトリにQuestionIDを引数で渡す。

	empID, valid, err := utils.ValidateToken(tokenString)
	if err != nil || !valid {
		fmt.Println("Error in ValidateToken:", err)
		return err
	}

	fmt.Println("empID:", empID)
	fmt.Println("valid:", valid)

	var answersBatch []models.Answers

	currentQID, err := s.repository.GetCurrentQIDByEmpID(empID)
	if err != nil {
		fmt.Println("Error in GetCurrentQIDByEmpID:", err)
		return err
	}

	fmt.Println("currentQID:", currentQID)

	for _, input := range inputs {
		latestAnswer, err := s.repository.GetLatestAnswer(empID, input.QuestionID)
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			fmt.Println("Error in GetLatestAnswer:", err)
			return err
		}

		streakCount := uint(0)
		if latestAnswer != nil {
			if input.AnswerID == 0 {
				streakCount = latestAnswer.StreakCount + 1
			} else {
				streakCount = 0
			}
		} else {
			if input.AnswerID == 0 {
				streakCount = 1
			} else {
				streakCount = 0
			}
		}

		answers := models.Answers{
			EmpID:       empID,
			QuestionID:  input.QuestionID,
			AnswerID:    input.AnswerID,
			StreakCount: streakCount,
		}

		fmt.Println("Answer being added:", answers)

		if answers.QuestionID > currentQID {
			currentQID = answers.QuestionID
		}

		answersBatch = append(answersBatch, answers)
	}

	err = s.repository.CreateAnswersBatch(answersBatch)
	if err != nil {
		fmt.Println("Error in CreateAnswersBatch:", err)
		return err
	}

	err = s.repository.UpdateCurrentQID(empID, currentQID)
	if err != nil {
		fmt.Println("Error in UpdateCurrentQID:", err)
		return err
	}

	return nil
}
