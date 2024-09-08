package data // テスト用データ

import (
	"backend/models"
	"time"

	"gorm.io/gorm"
)

var AnswersList = []models.Answers{
	// EMP1234のデータ
	// {EmpID: "EMP1234", QuestionID: 1, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
	// {EmpID: "EMP1234", QuestionID: 2, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
	// {EmpID: "EMP1234", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
	// {EmpID: "EMP1234", QuestionID: 4, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
	// {EmpID: "EMP1234", QuestionID: 5, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
	// {EmpID: "EMP1234", QuestionID: 6, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}},
	// {EmpID: "EMP1234", QuestionID: 7, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}},
	// {EmpID: "EMP1234", QuestionID: 8, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
	// {EmpID: "EMP1234", QuestionID: 9, AnswerID: 0, StreakCount: 2, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}},   // ヒットしない
	// {EmpID: "EMP1234", QuestionID: 10, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}}, // 105
	// {EmpID: "EMP1234", QuestionID: 1, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
	// {EmpID: "EMP1234", QuestionID: 2, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
	// {EmpID: "EMP1234", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
	// {EmpID: "EMP1234", QuestionID: 4, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
	// {EmpID: "EMP1234", QuestionID: 5, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
	// {EmpID: "EMP1234", QuestionID: 6, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}},
	// {EmpID: "EMP1234", QuestionID: 7, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}},
	// {EmpID: "EMP1234", QuestionID: 8, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
	// {EmpID: "EMP1234", QuestionID: 9, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -13)}},
	// {EmpID: "EMP1234", QuestionID: 10, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}}, // 95
	// {EmpID: "EMP1234", QuestionID: 1, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
	// {EmpID: "EMP1234", QuestionID: 2, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
	// {EmpID: "EMP1234", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
	// {EmpID: "EMP1234", QuestionID: 4, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
	// {EmpID: "EMP1234", QuestionID: 5, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
	// {EmpID: "EMP1234", QuestionID: 6, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}},
	// {EmpID: "EMP1234", QuestionID: 7, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}},
	// {EmpID: "EMP1234", QuestionID: 8, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
	// {EmpID: "EMP1234", QuestionID: 9, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -13)}},
	// {EmpID: "EMP1234", QuestionID: 10, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}}, // 85
	// {EmpID: "EMP1234", QuestionID: 1, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
	// {EmpID: "EMP1234", QuestionID: 2, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
	// {EmpID: "EMP1234", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
	// {EmpID: "EMP1234", QuestionID: 4, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
	// {EmpID: "EMP1234", QuestionID: 5, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
	// {EmpID: "EMP1234", QuestionID: 6, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}},
	// {EmpID: "EMP1234", QuestionID: 7, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}},
	// {EmpID: "EMP1234", QuestionID: 8, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
	// {EmpID: "EMP1234", QuestionID: 9, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -13)}},
	// {EmpID: "EMP1234", QuestionID: 10, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}}, // 75
	// {EmpID: "EMP1234", QuestionID: 1, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
	// {EmpID: "EMP1234", QuestionID: 2, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
	// {EmpID: "EMP1234", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
	// {EmpID: "EMP1234", QuestionID: 4, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
	// {EmpID: "EMP1234", QuestionID: 5, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
	// {EmpID: "EMP1234", QuestionID: 6, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}},
	// {EmpID: "EMP1234", QuestionID: 7, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}},
	// {EmpID: "EMP1234", QuestionID: 8, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
	// {EmpID: "EMP1234", QuestionID: 9, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -13)}},
	// {EmpID: "EMP1234", QuestionID: 10, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}}, // 65
	// {EmpID: "EMP1234", QuestionID: 1, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
	// {EmpID: "EMP1234", QuestionID: 2, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
	// {EmpID: "EMP1234", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
	// {EmpID: "EMP1234", QuestionID: 4, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
	// {EmpID: "EMP1234", QuestionID: 5, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
	// {EmpID: "EMP1234", QuestionID: 6, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}},
	// {EmpID: "EMP1234", QuestionID: 7, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}},
	// {EmpID: "EMP1234", QuestionID: 8, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
	// {EmpID: "EMP1234", QuestionID: 9, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -13)}},
	// {EmpID: "EMP1234", QuestionID: 10, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}}, // 55
	// {EmpID: "EMP1234", QuestionID: 1, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
	// {EmpID: "EMP1234", QuestionID: 2, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
	// {EmpID: "EMP1234", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
	// {EmpID: "EMP1234", QuestionID: 4, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
	// {EmpID: "EMP1234", QuestionID: 5, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
	// {EmpID: "EMP1234", QuestionID: 6, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}},
	// {EmpID: "EMP1234", QuestionID: 7, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}},
	// {EmpID: "EMP1234", QuestionID: 8, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
	// {EmpID: "EMP1234", QuestionID: 9, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -13)}},
	// {EmpID: "EMP1234", QuestionID: 10, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}}, // 45
	// {EmpID: "EMP1234", QuestionID: 1, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
	// {EmpID: "EMP1234", QuestionID: 2, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
	// {EmpID: "EMP1234", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
	// {EmpID: "EMP1234", QuestionID: 4, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
	// {EmpID: "EMP1234", QuestionID: 5, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
	// {EmpID: "EMP1234", QuestionID: 6, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}},
	// {EmpID: "EMP1234", QuestionID: 7, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}},
	// {EmpID: "EMP1234", QuestionID: 8, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
	// {EmpID: "EMP1234", QuestionID: 9, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -13)}},
	// {EmpID: "EMP1234", QuestionID: 10, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}}, // 35
	// {EmpID: "EMP1234", QuestionID: 1, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
	// {EmpID: "EMP1234", QuestionID: 2, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
	// {EmpID: "EMP1234", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
	// {EmpID: "EMP1234", QuestionID: 4, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
	// {EmpID: "EMP1234", QuestionID: 5, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
	// {EmpID: "EMP1234", QuestionID: 6, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}},
	// {EmpID: "EMP1234", QuestionID: 7, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}},
	// {EmpID: "EMP1234", QuestionID: 8, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
	// {EmpID: "EMP1234", QuestionID: 9, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -13)}},
	// {EmpID: "EMP1234", QuestionID: 10, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}}, // 25
	{EmpID: "EMP1234", QuestionID: 1, AnswerID: 0, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
	{EmpID: "EMP1234", QuestionID: 2, AnswerID: 0, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
	{EmpID: "EMP1234", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
	{EmpID: "EMP1234", QuestionID: 4, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
	{EmpID: "EMP1234", QuestionID: 5, AnswerID: 2, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
	{EmpID: "EMP1234", QuestionID: 6, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}}, // ヒット
	{EmpID: "EMP1234", QuestionID: 7, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}}, // ヒット
	{EmpID: "EMP1234", QuestionID: 8, AnswerID: 3, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
	{EmpID: "EMP1234", QuestionID: 9, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -13)}},
	{EmpID: "EMP1234", QuestionID: 10, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}}, // 15
	{EmpID: "EMP1234", QuestionID: 11, AnswerID: 0, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -11)}},
	{EmpID: "EMP1234", QuestionID: 12, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -10)}},
	{EmpID: "EMP1234", QuestionID: 13, AnswerID: 0, StreakCount: 2, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -9)}}, // ヒット
	{EmpID: "EMP1234", QuestionID: 14, AnswerID: 3, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -8)}},
	{EmpID: "EMP1234", QuestionID: 15, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -7)}},
	{EmpID: "EMP1234", QuestionID: 16, AnswerID: 0, StreakCount: 3, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -6)}},
	{EmpID: "EMP1234", QuestionID: 17, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -5)}},
	{EmpID: "EMP1234", QuestionID: 18, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -4)}},
	{EmpID: "EMP1234", QuestionID: 19, AnswerID: 0, StreakCount: 2, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -3)}}, // ヒット
	{EmpID: "EMP1234", QuestionID: 20, AnswerID: 0, StreakCount: 3, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -2)}}, // 5
	{EmpID: "EMP1234", QuestionID: 1, AnswerID: 0, StreakCount: 3, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}},
	{EmpID: "EMP1234", QuestionID: 2, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}}, // ヒット
	{EmpID: "EMP1234", QuestionID: 3, AnswerID: 0, StreakCount: 3, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}},
	{EmpID: "EMP1234", QuestionID: 4, AnswerID: 3, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}},
	{EmpID: "EMP1234", QuestionID: 5, AnswerID: 0, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}},

	// EMP2345のデータ
	// {EmpID: "EMP2345", QuestionID: 1, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -21)}},
	// {EmpID: "EMP2345", QuestionID: 2, AnswerID: 0, StreakCount: 2, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -20)}},
	// {EmpID: "EMP2345", QuestionID: 3, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -19)}},
	// {EmpID: "EMP2345", QuestionID: 4, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -18)}},
	// {EmpID: "EMP2345", QuestionID: 5, AnswerID: 2, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -17)}},
	// {EmpID: "EMP2345", QuestionID: 6, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -16)}},
	// {EmpID: "EMP2345", QuestionID: 7, AnswerID: 0, StreakCount: 2, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -15)}},
	// {EmpID: "EMP2345", QuestionID: 8, AnswerID: 3, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -14)}},
	// {EmpID: "EMP2345", QuestionID: 9, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -13)}},
	// {EmpID: "EMP2345", QuestionID: 10, AnswerID: 0, StreakCount: 2, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -12)}},
	// {EmpID: "EMP2345", QuestionID: 1, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -11)}},
	// {EmpID: "EMP2345", QuestionID: 2, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -10)}},
	// {EmpID: "EMP2345", QuestionID: 3, AnswerID: 0, StreakCount: 2, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -9)}},
	// {EmpID: "EMP2345", QuestionID: 4, AnswerID: 3, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -8)}},
	// {EmpID: "EMP2345", QuestionID: 5, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -7)}},
	// {EmpID: "EMP2345", QuestionID: 6, AnswerID: 0, StreakCount: 2, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -6)}},
	// {EmpID: "EMP2345", QuestionID: 7, AnswerID: 1, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -5)}},
	// {EmpID: "EMP2345", QuestionID: 8, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -4)}},
	// {EmpID: "EMP2345", QuestionID: 9, AnswerID: 2, StreakCount: 0, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -3)}},
	// {EmpID: "EMP2345", QuestionID: 10, AnswerID: 0, StreakCount: 1, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -2)}},
	// {EmpID: "EMP2345", QuestionID: 1, AnswerID: 0, StreakCount: 2, Model: gorm.Model{CreatedAt: time.Now().AddDate(0, 0, -1)}},
}
