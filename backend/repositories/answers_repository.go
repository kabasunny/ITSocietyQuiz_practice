package repositories

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type IAnswersRepository interface {
	CreateAnswers(answers *models.Answers) error                            // 新しい回答をデータベースに保存
	FindByEmpID(empID string) (*models.Answers, error)                      // 指定されたemp_idに基づいて回答を検索
	FindByQuestionID(QuestionID int) (*models.Answers, error)               // 指定されたquestion_idに基づいて回答を検索
	UpdateStreakCount(answer *models.Answers) error                         // 連続正解数を更新
	GetLatestAnswer(empID string, questionID uint) (*models.Answers, error) // 指定されたemp_idとquestion_idに基づいて最新の回答を取得
}

type AnswersRepository struct {
	db *gorm.DB
}

func NewAnswersRepository(db *gorm.DB) IAnswersRepository {
	return &AnswersRepository{db: db}
}

func (r *AnswersRepository) CreateAnswers(answers *models.Answers) error {
	result := r.db.Create(&answers)
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

func (r *AnswersRepository) UpdateStreakCount(answer *models.Answers) error {
	return r.db.Save(answer).Error
}

func (r *AnswersRepository) GetLatestAnswer(empID string, questionID uint) (*models.Answers, error) {
	var answer models.Answers
	err := r.db.Where("emp_id = ? AND question_id = ?", empID, questionID).Order("created_at desc").First(&answer).Error
	return &answer, err
}
