package data // テスト用データ

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Usersのデータモデルを表現する構造体
type User struct {
	EmpID          string
	Username       string
	Email          string
	Password       string
	TotalQuestions int
	CorrectAnswers int
}

// Usersのデータを格納する変数
// ユーザーのデータ登録時は、パスワードハッシュ化するので、下方に定義した関数を呼び出す
var UsersList = []User{
	{
		EmpID:          "EMP1234",
		Username:       "ITSocietyQuiz",
		Email:          "quize@example.com",
		Password:       "password",
		TotalQuestions: 0,
		CorrectAnswers: 0,
	},
	{
		EmpID:          "EMP2345",
		Username:       "ITSocietyQuiz_2",
		Email:          "quize_2@example.com",
		Password:       "password_2",
		TotalQuestions: 0,
		CorrectAnswers: 0,
	},
}

// ハッシュ化されたパスワードを持つ新しいUsersListを返却する関数
func GetHashedUsersList() []User {
	var hashedUsersList []User

	for _, user := range UsersList {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		//コストパラメータは、デフォルトだと10(範囲4～31)
		if err != nil {
			fmt.Printf("パスワードのハッシュ化に失敗しました: %v", err)
			return nil
		}

		hashedUser := User{
			EmpID:          user.EmpID,
			Username:       user.Username,
			Email:          user.Email,
			Password:       string(hashedPassword),
			TotalQuestions: user.TotalQuestions,
			CorrectAnswers: user.CorrectAnswers,
		}

		hashedUsersList = append(hashedUsersList, hashedUser)
	}

	return hashedUsersList
}
