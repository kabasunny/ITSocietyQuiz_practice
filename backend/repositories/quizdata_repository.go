package repositories

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type IQuizDataRepository interface {
	FindAll() (*[]models.QuizData, error)
	FindById(QuizDataId uint) (*models.QuizData, error)
	Create(newQuizData models.QuizData) (*models.QuizData, error)
	Update(updateQuizData models.QuizData) (*models.QuizData, error)
	Delete(QuizDataId uint) error
}

type QuizDataMemoryRepository struct {
	QuizDatas []models.QuizData
}

type QuizDataRepository struct {
	db *gorm.DB
}

func (r *QuizDataRepository) Create(newQuizData models.QuizData) (*models.QuizData, error) {
	result := r.db.Create(&newQuizData)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newQuizData, nil
}

func (r *QuizDataRepository) Delete(QuizDataId uint) error {
	deleteQuizData, err := r.FindById(QuizDataId)
	if err != nil {
		return err
	}
	result := r.db.Delete(&deleteQuizData) //論理削除
	// result := r.db.Unscoped().Delete(&deleteQuizData) // 物理削除
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *QuizDataRepository) FindAll() (*[]models.QuizData, error) {
	var QuizDatas []models.QuizData
	result := r.db.Find(&QuizDatas)
	if result.Error != nil {
		return nil, result.Error
	}
	return &QuizDatas, nil
}

func (r *QuizDataRepository) FindById(QuizDataId uint) (*models.QuizData, error) {
	var QuizData models.QuizData
	result := r.db.First(&QuizData, QuizDataId)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("QuizData not found")
		}
		return nil, result.Error
	}
	return &QuizData, nil
}

func (r *QuizDataRepository) Update(updateQuizData models.QuizData) (*models.QuizData, error) {
	result := r.db.Save(&updateQuizData)
	if result.Error != nil {
		return nil, result.Error
	}
	return &updateQuizData, nil
}

func NewQuizDataRepository(db *gorm.DB) IQuizDataRepository {
	return &QuizDataRepository{db: db}
}
