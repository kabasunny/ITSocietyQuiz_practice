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

	// ランダムシードを設定
	rand.Seed(time.Now().UnixNano())

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
	for i := 0; i < 200; i++ {
		empID := fmt.Sprintf("EMP%d", 100+i)
		username := fmt.Sprintf("ITSocietyQuiz%d", i+1)
		email := fmt.Sprintf("quize%d@example.com", i+1)
		password := fmt.Sprintf("password%d", i+1)
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("パスワードのハッシュ化に失敗しました: %v", err)
			continue
		}
		currentQID := uint(rand.Intn(101))
		totalQuestions := rand.Intn(201)
		correctAnswers := rand.Intn(totalQuestions + 1)

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

// // Usersのデータを格納する変数
// // ユーザーのデータ登録時は、パスワードハッシュ化するので、下方に定義した関数を呼び出す
// var UsersList = []models.Users{
// 	{
// 		EmpID:          "ADM1234",
// 		Username:       "ITSocietyQuiz_adm",
// 		Email:          "quize_adm@example.com",
// 		Password:       "password_a",
// 		TotalQuestions: 0,
// 		CorrectAnswers: 0,
// 	},
// 	{
// 		EmpID:          "ADM2345",
// 		Username:       "ITSocietyQuiz_adm",
// 		Email:          "quize_adm@example.com",
// 		Password:       "password_b",
// 		TotalQuestions: 0,
// 		CorrectAnswers: 0,
// 	},
// 	{
// 		EmpID:          "EMP1234",
// 		Username:       "ITSocietyQuiz",
// 		Email:          "quize@example.com",
// 		Password:       "password",
// 		CurrentQID:     60,
// 		TotalQuestions: 100,
// 		CorrectAnswers: 90,
// 	},
// 	{
// 		EmpID:          "EMP2234",
// 		Username:       "ITSocietyQuiz2",
// 		Email:          "quize2@example.com",
// 		Password:       "password2",
// 		CurrentQID:     58,
// 		TotalQuestions: 100,
// 		CorrectAnswers: 80,
// 	},
// 	{
// 		EmpID:          "EMP3234",
// 		Username:       "ITSocietyQuiz3",
// 		Email:          "quize3@example.com",
// 		Password:       "password3",
// 		CurrentQID:     56,
// 		TotalQuestions: 100,
// 		CorrectAnswers: 70,
// 	},
// 	{
// 		EmpID:          "EMP4234",
// 		Username:       "ITSocietyQuiz4",
// 		Email:          "quize4@example.com",
// 		Password:       "password4",
// 		CurrentQID:     54,
// 		TotalQuestions: 100,
// 		CorrectAnswers: 60,
// 	},
// 	{
// 		EmpID:          "EMP5234",
// 		Username:       "ITSocietyQuiz5",
// 		Email:          "quize5@example.com",
// 		Password:       "password5",
// 		CurrentQID:     52,
// 		TotalQuestions: 100,
// 		CorrectAnswers: 50,
// 	},
// 	{
// 		EmpID:          "EMP6234",
// 		Username:       "ITSocietyQuiz6",
// 		Email:          "quize6@example.com",
// 		Password:       "password6",
// 		CurrentQID:     50,
// 		TotalQuestions: 200,
// 		CorrectAnswers: 90,
// 	},
// 	{
// 		EmpID:          "EMP7234",
// 		Username:       "ITSocietyQuiz7",
// 		Email:          "quize7@example.com",
// 		Password:       "password7",
// 		CurrentQID:     100,
// 		TotalQuestions: 200,
// 		CorrectAnswers: 180,
// 	},
// 	{
// 		EmpID:          "EMP8234",
// 		Username:       "ITSocietyQuiz8",
// 		Email:          "quize8@example.com",
// 		Password:       "password8",
// 		CurrentQID:     100,
// 		TotalQuestions: 200,
// 		CorrectAnswers: 170,
// 	},
// 	{
// 		EmpID:          "EMP9234",
// 		Username:       "ITSocietyQuiz9",
// 		Email:          "quize9@example.com",
// 		Password:       "password9",
// 		CurrentQID:     100,
// 		TotalQuestions: 200,
// 		CorrectAnswers: 160,
// 	},
// 	{
// 		EmpID:          "EMP1334",
// 		Username:       "ITSocietyQuiz10",
// 		Email:          "quize10@example.com",
// 		Password:       "password10",
// 		CurrentQID:     100,
// 		TotalQuestions: 200,
// 		CorrectAnswers: 150,
// 	},
// }

// // ハッシュ化されたパスワードを持つ新しいUsersListを返却する関数
// func GetHashedUsersList() []models.Users {
// 	var hashedUsersList []models.Users

// 	for _, user := range UsersList {
// 		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
// 		//コストパラメータは、デフォルトだと10(範囲4～31)
// 		if err != nil {
// 			fmt.Printf("パスワードのハッシュ化に失敗しました: %v", err)
// 			return nil
// 		}

// 		hashedUser := models.Users{
// 			EmpID:          user.EmpID,
// 			Username:       user.Username,
// 			Email:          user.Email,
// 			Password:       string(hashedPassword),
// 			TotalQuestions: user.TotalQuestions,
// 			CurrentQID:     user.CurrentQID,
// 			CorrectAnswers: user.CorrectAnswers,
// 		}

// 		hashedUsersList = append(hashedUsersList, hashedUser)
// 	}

// 	return hashedUsersList
// }
