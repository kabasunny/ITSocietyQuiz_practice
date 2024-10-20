package data // テスト用データ
import (
	"backend/src/models"
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// 200人分のユーザーデータを生成し、パスワードをハッシュ化する関数
func GenerateHashedUsersList() []models.Users {
	var usersList []models.Users

	// ローカルなランダムジェネレータを作成
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// 管理者ユーザーを追加
	adminUsers := []models.Users{
		{
			EmpID:    "ADM1234",
			Username: "ITSocietyQuiz_adm1",
			Email:    "quize_adm1@example.com",
			Password: "password_a",
		},
		{
			EmpID:    "ADM2345",
			Username: "ITSocietyQuiz_adm2",
			Email:    "quize_adm2@example.com",
			Password: "password_b",
		},
	}

	for _, admin := range adminUsers {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(admin.Password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("パスワードのハッシュ化に失敗しました: %v", err)
			continue
		}
		admin.Password = string(hashedPassword)
		usersList = append(usersList, admin)
	}

	// 一般ユーザーを追加
	for i := 1; i < 200; i++ { // 社員199人分
		empID := fmt.Sprintf("EMP%d", 100+i)
		username := fmt.Sprintf("ITSocietyQuiz%d", i)
		email := fmt.Sprintf("quize%d@example.com", i)
		password := fmt.Sprintf("password%d", i)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("パスワードのハッシュ化に失敗しました: %v", err)
			continue
		}
		currentQID := uint(r.Intn(150) + 50)
		totalQuestions := r.Intn(400) + 200
		correctAnswers := r.Intn(350) + 150
		if correctAnswers > totalQuestions {
			correctAnswers = totalQuestions - int(correctAnswers/totalQuestions)
		}

		user := models.Users{
			EmpID:          empID,
			Username:       username,
			Email:          email,
			Password:       string(hashedPassword),
			CurrentQID:     currentQID,
			TotalQuestions: totalQuestions,
			CorrectAnswers: correctAnswers,
		}

		usersList = append(usersList, user)
	}

	return usersList
}

var UsersList = GenerateHashedUsersList()
