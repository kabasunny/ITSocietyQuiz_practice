package repositories

import (
	"backend/models"
	"errors"

	"gorm.io/gorm"
)

type ILoginRepository interface {
	CreateUsers(Users models.Users) error
	FindUsers(empID string) (*models.Users, error)
	FindUsersRole(empID string) ([]models.Users_roles, error) // ユーザーの役割（権限）を取得、とりあえずスライスを返す
}

type LoginRepository struct {
	db *gorm.DB
}

func NewLoginRepository(db *gorm.DB) ILoginRepository {
	return &LoginRepository{db: db}
}

func (r *LoginRepository) CreateUsers(Users models.Users) error {
	result := r.db.Create(&Users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *LoginRepository) FindUsers(empID string) (*models.Users, error) {
	var Users models.Users
	result := r.db.First(&Users, "emp_id = ?", empID) //第2引数はSQLのWhere句に相当し、プレースホルダに第3引数が格納される
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, errors.New("users not found")
		}
		return nil, result.Error
	}
	return &Users, nil
}

func (r *LoginRepository) FindUsersRole(empID string) ([]models.Users_roles, error) {
	var roles []models.Users_roles
	result := r.db.Where("emp_id = ?", empID).Find(&roles)
	if result.Error != nil {
		return nil, result.Error
	}
	return roles, nil
}
