package data // テスト用データ

import (
	"backend/src/models"
	"fmt"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

// ローカルなランダムジェネレータを作成
var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// AnswersListのデータを生成する関数
func GenerateAnswersList() []models.Answers {
	var answersList []models.Answers

	for i := 0; i < 200; i++ { // 社員200人分
		empID := fmt.Sprintf("EMP%d", 100+i)
		for j := 0; j < 21; j++ { // 過去21日分
			questionID := uint(r.Intn(3) + 1 + int(j/5))
			answerID := uint(r.Intn(5))    // 0から4のランダムな値
			streakCount := uint(r.Intn(3)) // 0から2のランダムな値

			answer := models.Answers{
				EmpID:       empID,
				QuestionID:  questionID,
				AnswerID:    answerID,
				StreakCount: streakCount,
				Model:       gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -j)}, // 過去21日間の日付
			}
			answersList = append(answersList, answer)
		}
	}

	return answersList
}

var AnswersList = GenerateAnswersList()

// 出題アルゴリズムの確認用
// var AnswersList = []models.Answers{
// 	// EMP1234のデータ
// 	{EmpID: "EMP1234", QuestionID: 1, AnswerID: 0, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
// 	{EmpID: "EMP1234", QuestionID: 2, AnswerID: 0, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
// 	{EmpID: "EMP1234", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
// 	{EmpID: "EMP1234", QuestionID: 4, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
// 	{EmpID: "EMP1234", QuestionID: 5, AnswerID: 2, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
// 	{EmpID: "EMP1234", QuestionID: 6, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}}, // ヒット
// 	{EmpID: "EMP1234", QuestionID: 7, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}}, // ヒット
// 	{EmpID: "EMP1234", QuestionID: 8, AnswerID: 3, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
// 	{EmpID: "EMP1234", QuestionID: 9, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -13)}},
// 	{EmpID: "EMP1234", QuestionID: 10, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}}, // 15
// 	{EmpID: "EMP1234", QuestionID: 11, AnswerID: 0, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -11)}},
// 	{EmpID: "EMP1234", QuestionID: 12, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -10)}},
// 	{EmpID: "EMP1234", QuestionID: 13, AnswerID: 0, StreakCount: 2, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -9)}}, // ヒット
// 	{EmpID: "EMP1234", QuestionID: 14, AnswerID: 3, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -8)}},
// 	{EmpID: "EMP1234", QuestionID: 15, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -7)}},
// 	{EmpID: "EMP1234", QuestionID: 16, AnswerID: 0, StreakCount: 3, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -6)}},
// 	{EmpID: "EMP1234", QuestionID: 17, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -5)}},
// 	{EmpID: "EMP1234", QuestionID: 18, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -4)}},
// 	{EmpID: "EMP1234", QuestionID: 19, AnswerID: 0, StreakCount: 2, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -3)}}, // ヒット
// 	{EmpID: "EMP1234", QuestionID: 20, AnswerID: 0, StreakCount: 3, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -2)}}, // 5
// 	{EmpID: "EMP1234", QuestionID: 1, AnswerID: 0, StreakCount: 3, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}},
// 	{EmpID: "EMP1234", QuestionID: 2, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}}, // ヒット
// 	{EmpID: "EMP1234", QuestionID: 3, AnswerID: 0, StreakCount: 3, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}},
// 	{EmpID: "EMP1234", QuestionID: 4, AnswerID: 3, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}},
// 	{EmpID: "EMP1234", QuestionID: 5, AnswerID: 0, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}},
