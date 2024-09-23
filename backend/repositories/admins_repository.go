package repositories

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type IAdminsRepository interface {
	FindAll() (*[]models.Questions, error)
	FindById(QuestionsId uint) (*models.Questions, error) // 単一の質問IDに基づいて、questionsテーブルからデータを取得する場合
	Create(newQuestions models.Questions) (*models.Questions, error)
	Update(updateQuestions *models.Questions) (*models.Questions, error)
	Delete(QuestionsId uint) error
	Count() (int64, error)                          // 格納されたクイズのレコード数を取得するメソッドを追加
	CreateQuestionsBatch([]*models.Questions) error // 追加
}

type AdminsRepository struct {
	db *gorm.DB
}

func NewAdminsRepository(db *gorm.DB) IAdminsRepository {
	return &AdminsRepository{db: db}
}

func (r *AdminsRepository) Create(newQuestions models.Questions) (*models.Questions, error) {
	result := r.db.Create(&newQuestions)
	if result.Error != nil {
		return nil, result.Error
	}
	return &newQuestions, nil
}

func (r *AdminsRepository) Delete(QuestionsId uint) error {
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

func (r *AdminsRepository) FindAll() (*[]models.Questions, error) {
	var Questions []models.Questions
	result := r.db.Find(&Questions)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Questions, nil
}

func (r *AdminsRepository) FindById(QuestionsId uint) (*models.Questions, error) {
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

func (r *AdminsRepository) Update(updateQuestions *models.Questions) (*models.Questions, error) {
	result := r.db.Save(&updateQuestions)
	if result.Error != nil {
		return nil, result.Error
	}
	return updateQuestions, nil
}

// クイズデータのレコード総数をカウント
func (r *AdminsRepository) Count() (int64, error) {
	var count int64
	result := r.db.Model(&models.Questions{}).Count(&count)
	if result.Error != nil {
		return 0, result.Error
	}
	return count, nil
}

// クイズデータをバッチで作成
func (r *AdminsRepository) CreateQuestionsBatch(data []*models.Questions) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Create(&data)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
