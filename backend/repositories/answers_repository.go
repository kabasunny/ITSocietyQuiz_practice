package repositories

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type IAnswersRepository interface {
	CreateAnswers(Answers models.Answers) error
	FindByEmpID(empID string) (*models.Answers, error)
	FindByQuestionID(QuestionID int) (*models.Answers, error)
}

type AnswersRepository struct {
	db *gorm.DB
}

func NewAnswersRepository(db *gorm.DB) IAnswersRepository {
	return &AnswersRepository{db: db}
}

func (r *AnswersRepository) CreateAnswers(Answers models.Answers) error {
	result := r.db.Create(&Answers)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AnswersRepository) FindByEmpID(empID string) (*models.Answers, error) {
	var Answers models.Answers
	result := r.db.First(&Answers, "emp_id = ?", empID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("answers not found")
		}
		return nil, result.Error
	}
	return &Answers, nil
}

func (r *AnswersRepository) FindByQuestionID(QuestionID int) (*models.Answers, error) {
	var Answers models.Answers
	result := r.db.First(&Answers, "question_id = ?", QuestionID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("answers not found")
		}
		return nil, result.Error
	}
	return &Answers, nil
}
