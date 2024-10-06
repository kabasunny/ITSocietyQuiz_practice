package repositories

import (
	"backend/src/dto"
	"backend/src/models"
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type IAdminsRepository interface {
	FindAllQuestions() (*[]models.Questions, error)
	FindQuestionsById(QuestionsId uint) (*models.Questions, error) // 単一の質問IDに基づいて、questionsテーブルからデータを取得する場合
	// CreateQuestions(newQuestions models.Questions) (*models.Questions, error)
	UpdateQuestions(updateQuestions *models.Questions) (*models.Questions, error)
	DeleteQuestions(QuestionsId uint) error
	CountQuestions() (int64, error)                              // 格納されたクイズのレコード数を取得するメソッドを追加
	CreateQuestionsBatch([]*models.Questions) error              // 追加
	GetUsersInfomation(query string) ([]*dto.AdmUserData, error) // ユーザーの一覧を取得する
	UpdateUsers(updateUsers *models.Users) (*models.Users, error)
	GetUserByDBID(dbId uint) (*models.Users, error) // 主キーの取得
	InsertUserRole(empID string, roleID uint) error
	GetRoleIDByEmpID(empID string) (uint, error)
	GetRoleNameByID(roleID uint) (string, error)
	AddUsers(newUser *models.Users) (*models.Users, error)
	DeleteUsers(dbId uint) error
	GetRanking(query string) ([]*dto.RankingData, error)                               // ランキングを取得する
	GetOverallPerformanceData(query string) (*models.Users, *models.Answers, error)    // 全体の傾向のデータを取得する
	GetIndividualPerformanceData(empID string) (*models.Users, *models.Answers, error) // 個人の成績のデータを取得する

}

type AdminsRepository struct {
	db *gorm.DB
}

func NewAdminsRepository(db *gorm.DB) IAdminsRepository {
	return &AdminsRepository{db: db}
}

// func (r *AdminsRepository) CreateQuestions(newQuestions models.Questions) (*models.Questions, error) {
// 	result := r.db.Create(&newQuestions)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &newQuestions, nil
// }

func (r *AdminsRepository) DeleteQuestions(QuestionsId uint) error {
	deleteQuestions, err := r.FindQuestionsById(QuestionsId)
	if err != nil {
		return err
	}
	// result := r.db.Delete(&deleteQuestions) //論理削除
	result := r.db.Unscoped().Delete(&deleteQuestions) // ユーザーの削除に合わせて、物理削除に変更
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AdminsRepository) FindAllQuestions() (*[]models.Questions, error) {
	var Questions []models.Questions
	result := r.db.Find(&Questions)
	if result.Error != nil {
		return nil, result.Error
	}
	return &Questions, nil
}

func (r *AdminsRepository) FindQuestionsById(QuestionsId uint) (*models.Questions, error) {
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

func (r *AdminsRepository) UpdateQuestions(updateQuestions *models.Questions) (*models.Questions, error) {
	result := r.db.Save(&updateQuestions)
	if result.Error != nil {
		return nil, result.Error
	}
	return updateQuestions, nil
}

// クイズデータのレコード総数をカウント
func (r *AdminsRepository) CountQuestions() (int64, error) {
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

func (r *AdminsRepository) GetUsersInfomation(query string) ([]*dto.AdmUserData, error) {
	var users []*dto.AdmUserData
	if err := r.db.Raw(query).Scan(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *AdminsRepository) UpdateUsers(updateUsers *models.Users) (*models.Users, error) {
	result := r.db.Save(&updateUsers)
	if result.Error != nil {
		return nil, result.Error
	}
	return updateUsers, nil
}

func (r *AdminsRepository) GetUserByDBID(dbId uint) (*models.Users, error) {
	var user models.Users
	result := r.db.First(&user, dbId)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}
	return &user, nil
}

func (r *AdminsRepository) InsertUserRole(empID string, roleID uint) error {
	userRole := models.UsersRoles{
		EmpID:  empID,
		RoleID: roleID,
	}
	result := r.db.Create(&userRole)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AdminsRepository) GetRoleIDByEmpID(empID string) (uint, error) {
	var userRole models.UsersRoles
	result := r.db.First(&userRole, "emp_id = ?", empID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, errors.New("role_id not found")
		}
		return 0, result.Error
	}
	return userRole.RoleID, nil
}

func (r *AdminsRepository) GetRoleNameByID(roleID uint) (string, error) {
	var role models.Roles
	result := r.db.First(&role, roleID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "", errors.New("role not found")
		}
		return "", result.Error
	}
	return role.RoleName, nil
}

func (r *AdminsRepository) AddUsers(newUser *models.Users) (*models.Users, error) {
	result := r.db.Create(newUser)
	if result.Error != nil {
		return nil, result.Error
	}
	return newUser, nil
}

func (r *AdminsRepository) DeleteUsers(dbId uint) error {
	deleteUsers, err := r.GetUserByDBID(dbId) // フロントからdbIdが送信できれば流用可能
	if err != nil {
		return err
	}
	// result := r.db.Delete(&deleteUsers) //論理削除
	result := r.db.Unscoped().Delete(&deleteUsers) // 物理削除。論理削除では、削除後に同じユーザーIDを登録できない
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *AdminsRepository) GetRanking(query string) ([]*dto.RankingData, error) {
	var ranking []*dto.RankingData
	if err := r.db.Raw(query).Scan(&ranking).Error; err != nil {
		return nil, err
	}
	for _, data := range ranking {
		fmt.Printf("EmpID: %s, Username: %s, CurrentQID: %d, CorrectAnswerRate: %f, PerformanceIndicator: %f, Rank: %d\n",
			data.EmpID, *data.Username, data.CurrentQID, data.C, data.P, data.Rank) // correctAnswerRateやperformanceIndicatorでは、何故がうまくいかない
	}
	return ranking, nil
}

func (r *AdminsRepository) GetOverallPerformanceData(query string) (*models.Users, *models.Answers, error) {
	return nil, nil, nil
}

func (r *AdminsRepository) GetIndividualPerformanceData(empID string) (*models.Users, *models.Answers, error) {
	return nil, nil, nil

}
