package models

import (
	"gorm.io/gorm"
)

// Answers構造体は、クイズの回答を表現する
type Answers struct {
	gorm.Model             // IDフィールドを含む構造体
	EmpID          string  `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:EmpID;references:EmpID"`   // 外部キーとしてUsersのEmpIDを設定
	QuestionID     uint    `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;foreignKey:QuestionID;references:ID"` // 外部キーとしてQuestionsのIDを設定
	UserQuestionID *string // ユーザー管理用の質問ID、NULLを許容
	AnswerID       uint    `gorm:"not null"`
	StreakCount    uint    `gorm:"not null;default:0"` // 連続正解数を保持するフィールド、出題時に使用する
	// IsCorrect  bool      `gorm:"not null"` // Answer == 0 ならば正解
	// Timestamp time.Time `gorm:"not null"` // GORMのCreatedAt time.Timeを使用する
}
