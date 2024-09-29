package services

import (
	"backend/dto"
	"backend/models"
	"backend/repositories"
	"backend/utils"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type IAdminsService interface {
	FindAllQuestions() (*[]dto.AdmQuizData, error) // 修正
	FindQuestionsById(QuestionsId uint) (*models.Questions, error)
	// CreateQuestions(createQuestionsInput dto.CreateQuestionsInput) (*models.Questions, error)
	UpdateQuestions(QuestionsId uint, updateQuestionsInput dto.UpdateQuestionsInput) (*dto.UpdateQuestionsOutput, error)
	DeleteQuestions(QuestionsId uint) error
	ProcessCSVData(filepath string) error // 追加
	GetUsersInfomation() ([]*dto.AdmUserData, error)
	UpdateUsers(dbId uint, updateUsers dto.AdmUserData) (*dto.AdmUserData, error)
}

type AdminsService struct {
	repository repositories.IAdminsRepository
}

func NewAdminsService(repository repositories.IAdminsRepository) IAdminsService {
	return &AdminsService{repository: repository}
}

func (s *AdminsService) FindAllQuestions() (*[]dto.AdmQuizData, error) { // 修正
	questions, err := s.repository.FindAllQuestions()
	if err != nil {
		return nil, err
	}

	var quizData []dto.AdmQuizData
	for _, question := range *questions {
		quizData = append(quizData, dto.AdmQuizData{
			ID:             question.ID,
			UserQuestionID: question.UserQuestionID,
			Question:       question.Question,
			Options:        question.Options,
			Supplement:     question.Supplement,
			Difficulty:     question.Difficulty,
			CreatedAt:      question.CreatedAt.Format(time.RFC3339),
			UpdatedAt:      question.UpdatedAt.Format(time.RFC3339),
		})
	}

	return &quizData, nil
}

func (s *AdminsService) FindQuestionsById(QuestionsId uint) (*models.Questions, error) {
	return s.repository.FindQuestionsById(QuestionsId)
}

// func (s *AdminsService) CreateQuestions(createQuestionsInput dto.CreateQuestionsInput) (*models.Questions, error) {
// 	newQuestions := models.Questions{
// 		Question:   createQuestionsInput.Question,
// 		Options:    createQuestionsInput.Options,
// 		Supplement: createQuestionsInput.Supplement,
// 		Difficulty: createQuestionsInput.Difficulty,
// 	}
// 	return s.repository.Create(newQuestions)
// }

func (s *AdminsService) UpdateQuestions(QuestionsId uint, updateQuestionsInput dto.UpdateQuestionsInput) (*dto.UpdateQuestionsOutput, error) {
	targetQuestions, err := s.FindQuestionsById(QuestionsId)
	if err != nil {
		return nil, err
	}
	if updateQuestionsInput.Question != nil {
		targetQuestions.UserQuestionID = updateQuestionsInput.UserQuestionID // ポインタ型
	}
	if updateQuestionsInput.Question != nil {
		targetQuestions.Question = *updateQuestionsInput.Question
	}
	if updateQuestionsInput.Options != nil {
		targetQuestions.Options = pq.StringArray(*updateQuestionsInput.Options) // 一応キャスト　pq.StringArrayは[]stringのエイリアスで、PostgreSQLのtext[]型と直接互換性がある
	}
	if updateQuestionsInput.Supplement != nil {
		targetQuestions.Supplement = *updateQuestionsInput.Supplement
	}
	if updateQuestionsInput.Difficulty != nil {
		targetQuestions.Difficulty = *updateQuestionsInput.Difficulty
	}
	updatedQuestions, err := s.repository.UpdateQuestions(targetQuestions)
	if err != nil {
		return nil, err
	}

	// モデル構造体をDTO構造体に変換
	updateQuestionsOutput := &dto.UpdateQuestionsOutput{
		ID:             updatedQuestions.ID,
		UserQuestionID: updatedQuestions.UserQuestionID,
		Question:       updatedQuestions.Question,
		Options:        updatedQuestions.Options,
		Supplement:     updatedQuestions.Supplement,
		Difficulty:     updatedQuestions.Difficulty,
		CreatedAt:      updatedQuestions.CreatedAt.Format(time.RFC3339),
		UpdatedAt:      updatedQuestions.UpdatedAt.Format(time.RFC3339),
	}

	return updateQuestionsOutput, nil
}

func (s *AdminsService) DeleteQuestions(QuestionsId uint) error {
	return s.repository.DeleteQuestions(QuestionsId)
}

func (s *AdminsService) ProcessCSVData(filePath string) error {
	data, err := utils.ParseCSV(filePath)
	if err != nil {
		return err
	}

	// データをリポジトリに直接渡す
	if err := s.repository.CreateQuestionsBatch(data); err != nil {
		return err
	}

	return nil
}

func (s *AdminsService) GetUsersInfomation() ([]*dto.AdmUserData, error) {

	// SQLクエリを読み込む
	query, err := utils.LoadSQLFile("services/queries/select_users_with_roles.sql")
	if err != nil {
		log.Fatalf("Failed to load SQL file: %v", err)
	}

	userList, err := s.repository.GetUsersInfomation(query)

	return userList, err
}

func (s *AdminsService) UpdateUsers(dbId uint, updateUsers dto.AdmUserData) (*dto.AdmUserData, error) {
	// データベースから既存のユーザーを取得、ユーザーIDではなくGORMのidで検索する
	user, err := s.repository.GetUserByDBID(dbId)
	if err != nil {
		log.Printf("Failed to get user by DB ID: %v", err)
		return nil, fmt.Errorf("failed to get user by DB ID: %w", err)
	}

	// パスワードの更新が必要か確認
	if updateUsers.Password_1 != "" && updateUsers.Password_2 != "" {
		// 現在のパスワードを比較
		err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(updateUsers.Password_1))
		if err != nil {
			log.Printf("Old password does not match: %v", err)
			return nil, errors.New("old password does not match")
		}

		// 新しいパスワードをハッシュ化
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(updateUsers.Password_2), bcrypt.DefaultCost)
		if err != nil {
			log.Printf("Failed to hash password: %v", err)
			return nil, fmt.Errorf("failed to hash password: %w", err)
		}
		user.Password = string(hashedPassword)
	}

	// その他のフィールドを更新
	user.EmpID = updateUsers.EmpID
	user.Username = *updateUsers.Username
	user.Email = updateUsers.Email

	// 更新されたユーザー情報をリポジトリを通じて保存
	updatedUser, err := s.repository.UpdateUsers(user)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// RoleIDが0の場合、既存のRoleIDを取得して設定
	if updateUsers.RoleID == 0 {
		updateUsers.RoleID, err = s.repository.GetRoleIDByEmpID(user.EmpID)
		if err != nil {
			log.Printf("Failed to get existing role ID: %v", err)
			return nil, fmt.Errorf("failed to get existing role ID: %w", err)
		}
	} else {
		// RoleIDが0でない場合のみ、users_rolesテーブルにレコードを挿入
		log.Printf("Inserting user role: EmpID=%s, RoleID=%d", user.EmpID, updateUsers.RoleID)
		err = s.repository.InsertUserRole(user.EmpID, updateUsers.RoleID)
		if err != nil {
			log.Printf("Failed to insert user role: %v", err)
			return nil, fmt.Errorf("failed to insert user role: %w", err)
		}
	}

	// rolesテーブルからRoleNameを取得
	log.Printf("Getting role name for RoleID=%d", updateUsers.RoleID)
	roleName, err := s.repository.GetRoleNameByID(updateUsers.RoleID)
	if err != nil {
		log.Printf("Failed to get role name: %v", err)
		return nil, fmt.Errorf("failed to get role name: %w", err)
	}

	// DTOに変換して返却
	updatedUserData := &dto.AdmUserData{
		ID:        updatedUser.ID,
		EmpID:     updatedUser.EmpID,
		Username:  &updatedUser.Username,
		Email:     updatedUser.Email,
		RoleName:  roleName,
		CreatedAt: updatedUser.CreatedAt.Format(time.RFC3339),
		UpdatedAt: updatedUser.UpdatedAt.Format(time.RFC3339),
	}

	return updatedUserData, nil
}
func (s *AdminsService) AddUser(newUsers dto.AdmUserData) (*dto.AdmUserData, error) {
	user := &models.Users{}

	// パスワードと確認用パスワードを比較
	if newUsers.Password_1 == "" && newUsers.Password_2 != "" {

		log.Printf("password does not match")
		return nil, errors.New("password does not match")
	}

	// パスワードをハッシュ化
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUsers.Password_1), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = string(hashedPassword)

	// その他のフィールドを更新
	user.EmpID = newUsers.EmpID
	user.Username = *newUsers.Username
	user.Email = newUsers.Email

	// 更新されたユーザー情報をリポジトリを通じて保存
	addedUsers, err := s.repository.AddUser(user)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		return nil, fmt.Errorf("failed to update user: %w", err)
	}

	// 1. users_rolesテーブルにレコードを挿入
	log.Printf("Inserting user role: EmpID=%s, RoleID=%d", user.EmpID, newUsers.RoleID)
	err = s.repository.InsertUserRole(user.EmpID, newUsers.RoleID)
	if err != nil {
		log.Printf("Failed to insert user role: %v", err)
		return nil, fmt.Errorf("failed to insert user role: %w", err)
	}

	// 2. rolesテーブルからRoleNameを取得
	log.Printf("Getting role name for RoleID=%d", newUsers.RoleID)
	roleName, err := s.repository.GetRoleNameByID(newUsers.RoleID)
	if err != nil {
		log.Printf("Failed to get role name: %v", err)
		return nil, fmt.Errorf("failed to get role name: %w", err)
	}

	// DTOに変換して返却
	addedUsersData := &dto.AdmUserData{
		ID:        addedUsers.ID,
		EmpID:     addedUsers.EmpID,
		Username:  &addedUsers.Username,
		Email:     addedUsers.Email,
		RoleName:  roleName,
		CreatedAt: addedUsers.CreatedAt.Format(time.RFC3339),
		UpdatedAt: addedUsers.UpdatedAt.Format(time.RFC3339),
	}

	return addedUsersData, nil
}
