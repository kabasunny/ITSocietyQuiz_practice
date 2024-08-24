package repositories

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

// 初期テスト用
type IQuestionsRepository interface {
	FindAll() (*[]models.Questions, error)
	FindById(QuestionsId uint) (*models.Questions, error)
	Create(newQuestions models.Questions) (*models.Questions, error)
	Update(updateQuestions models.Questions) (*models.Questions, error)
	Delete(QuestionsId uint) error
}

type QuestionsMemoryRepository struct {
	Questionss []models.Questions
}

type QuestionsRepository struct {
	db *gorm.DB
}

func (r *QuestionsRepository) Create(newQuestions models.Questions) (*models.Questions, error) {
	result := r.db.Create(&newQuestions)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newQuestions, nil
}

func (r *QuestionsRepository) Delete(QuestionsId uint) error {
	deleteQuestions, err := r.FindById(QuestionsId)
	if err != nil {
		return err
	}
	result := r.db.Delete(&deleteQuestions) //論理削除
	// result := r.db.Unscoped().Delete(&deleteQuestions) // 物理削除
	if result.Error != nil {
		return result.Error
	}
	return nil
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
			return nil, errors.New("Questions not found")
		}
		return nil, result.Error
	}
	return &Questions, nil
}

func (r *QuestionsRepository) Update(updateQuestions models.Questions) (*models.Questions, error) {
	result := r.db.Save(&updateQuestions)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateQuestions, nil
}

func NewQuestionsRepository(db *gorm.DB) IQuestionsRepository {
	return &QuestionsRepository{db: db}
}
