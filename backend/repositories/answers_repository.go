package repositories

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type IAnswersRepository interface {
	CreateAnswersBatch(answers []models.Answers) error                      // 新しい回答をデータベースに保存
	FindByEmpID(empID string) (*models.Answers, error)                      // 指定されたemp_idに基づいて回答を検索
	FindByQuestionID(QuestionID int) (*models.Answers, error)               // 指定されたquestion_idに基づいて回答を検索
	UpdateStreakCount(answer *models.Answers) error                         // 連続正解数を更新
	GetLatestAnswer(empID string, questionID uint) (*models.Answers, error) // 指定されたemp_idとquestion_idに基づいて最新の回答を取得
	UpdateCurrentQID(empID string, newQID uint) error                       // usersテーブルのcurrentq_idを更新する
	GetCurrentQIDByEmpID(empID string) (uint, error)                        // usersテーブルからcurrentq_idを取得する
}

type AnswersRepository struct {
	db *gorm.DB
}

func NewAnswersRepository(db *gorm.DB) IAnswersRepository {
	return &AnswersRepository{db: db}
}

func (r *AnswersRepository) CreateAnswersBatch(answers []models.Answers) error {
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	result := tx.Create(&answers) // コントローラ、サービスで繰り返し処理をしていたので、ここで一括して行う様変更
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	if err := tx.Commit().Error; err != nil {
		return err
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
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // record not found はエラーにしない
	}
	return &answer, err
}

// usersテーブルのcurrentq_idを更新する
func (r *AnswersRepository) UpdateCurrentQID(empID string, newQID uint) error {
	result := r.db.Exec("UPDATE users SET current_q_id = ? WHERE emp_id = ? AND current_q_id < ?", newQID, empID, newQID)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// usersテーブルからcurrentq_idを取得する
func (r *AnswersRepository) GetCurrentQIDByEmpID(empID string) (uint, error) {
	var currentQID uint
	result := r.db.Model(&models.Users{}).Where("emp_id = ?", empID).Select("current_q_id").Scan(&currentQID)
	if result.Error != nil {
		return 0, result.Error
	}
	return currentQID, nil
}
