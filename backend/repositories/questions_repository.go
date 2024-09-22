package repositories

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type IQuestionsRepository interface {
	FindAll() (*[]models.Questions, error)
	FindById(QuestionsId uint) (*models.Questions, error)                                                   // 単一の質問IDに基づいて、questionsテーブルからデータを取得する場合
	Count() (int64, error)                                                                                  // 格納されたクイズのレコード数を取得するメソッドを追加
	GetTopQuestionsByEmpID(query string, empID string, limit uint, necessaryQuestions uint) ([]uint, error) // answersテーブルからlimit件の質問IDを取得し、優先度順に並べる
	GetCurrentQIDByEmpID(empID string) (uint, error)                                                        // usersテーブルからcurrentq_idを取得する
	GetQuestionDetails(questionIDs []uint) ([]models.Questions, error)                                      // 複数の質問IDに基づいて、questionsテーブルからデータを取得する場合
}

type QuestionsRepository struct {
	db *gorm.DB
}

func NewQuestionsRepository(db *gorm.DB) IQuestionsRepository {
	return &QuestionsRepository{db: db}
}

func (r *QuestionsRepository) FindAll() (*[]models.Questions, error) {
	var Questions []models.Questions
	result := r.db.Find(&Questions)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Questions, nil
}

func (r *QuestionsRepository) FindById(QuestionsId uint) (*models.Questions, error) {
	var Questions models.Questions
	result := r.db.First(&Questions, QuestionsId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("questions not found")
		}
		return nil, result.Error
	}
	return &Questions, nil
}

// クイズデータのレコード総数をカウント
func (r *QuestionsRepository) Count() (int64, error) {
	var count int64
	result := r.db.Model(&models.Questions{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// answersテーブルからlimit件の質問IDを取得し、優先度順に並べる
func (r *QuestionsRepository) GetTopQuestionsByEmpID(query string, empID string, limit uint, necessaryQuestions uint) ([]uint, error) {
	var questionIDs []uint
	//     `
	result := r.db.Raw(query, empID, limit, necessaryQuestions).Scan(&questionIDs)
	if result.Error != nil {
		return nil, result.Error
	}
	return questionIDs, nil
}

// usersテーブルからcurrentq_idを取得する
func (r *QuestionsRepository) GetCurrentQIDByEmpID(empID string) (uint, error) {
	var currentQID uint
	result := r.db.Model(&models.Users{}).Where("emp_id = ?", empID).Select("current_q_id").Scan(&currentQID)
	if result.Error != nil {
		return 0, result.Error
	}
	return currentQID, nil
}

// questionsテーブルから、詳細データを取得する
func (r *QuestionsRepository) GetQuestionDetails(questionIDs []uint) ([]models.Questions, error) {
	var questions []models.Questions
	result := r.db.Where("id IN (?)", questionIDs).Find(&questions)
	if result.Error != nil {
		return nil, result.Error
	}
	return questions, nil
}
