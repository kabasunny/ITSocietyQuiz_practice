package repositories

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type IQuestionsRepository interface {
	FindAll() (*[]models.Questions, error)
	FindById(QuestionsId uint) (*models.Questions, error) // 単一の質問IDに基づいて、questionsテーブルからデータを取得する場合
	Create(newQuestions models.Questions) (*models.Questions, error)
	Update(updateQuestions models.Questions) (*models.Questions, error)
	Delete(QuestionsId uint) error
	Count() (int64, error)                                                                                  // 格納されたクイズのレコード数を取得するメソッドを追加
	GetTopQuestionsByEmpID(query string, empID string, limit uint, dailyQuestionCount uint) ([]uint, error) // answersテーブルからlimit件の質問IDを取得し、優先度順に並べる
	GetCurrentQIDByEmpID(empID string) (uint, error)                                                        // usersテーブルからcurrentq_idを取得する
	UpdateCurrentQID(empID string, newQID uint) error                                                       // usersテーブルのcurrentq_idを更新する
	GetQuestionDetails(questionIDs []uint) ([]models.Questions, error)                                      // 複数の質問IDに基づいて、questionsテーブルからデータを取得する場合
}

type QuestionsMemoryRepository struct {
	Questionss []models.Questions
}

type QuestionsRepository struct {
	db *gorm.DB
}

func NewQuestionsRepository(db *gorm.DB) IQuestionsRepository {
	return &QuestionsRepository{db: db}
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
			return nil, errors.New("questions not found")
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
func (r *QuestionsRepository) GetTopQuestionsByEmpID(query string, empID string, limit uint, dailyQuestionCount uint) ([]uint, error) {
	var questionIDs []uint
	// query2 := `
	// WITH LatestAnswers AS (
	//     SELECT
	//         question_id,
	//         streak_count,
	//         ROW_NUMBER() OVER (PARTITION BY question_id ORDER BY created_at DESC) AS rn
	//     FROM
	//         answers
	//     WHERE
	//         emp_id = ?
	//     LIMIT ?
	// )
	// SELECT
	//     question_id
	// FROM
	//     LatestAnswers
	// WHERE
	//     rn = 1 AND streak_count IN (0, 1, 2)
	// ORDER BY
	//     streak_count DESC,
	//     question_id ASC
	// LIMIT ?;
	//     `
	result := r.db.Raw(query, empID, limit, dailyQuestionCount).Scan(&questionIDs)
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

// usersテーブルのcurrentq_idを更新する
func (r *QuestionsRepository) UpdateCurrentQID(empID string, newQID uint) error {
	result := r.db.Model(&models.Users{}).Where("emp_id = ?", empID).Update("current_q_id", newQID)
	if result.Error != nil {
		return result.Error
	}
	return nil
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
