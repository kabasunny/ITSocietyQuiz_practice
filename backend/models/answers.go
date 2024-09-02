package models

import (
	"gorm.io/gorm"
)

// Answers構造体は、クイズの回答を表現する
type Answers struct {
	gorm.Model        // IDフィールドを含む構造体
	EmpID      string `gorm:"not null;constraint:OnDelete:CASCADE"`
	QuestionID uint   `gorm:"not null;constraint:OnDelete:CASCADE"` // 外部キーとしてQuestionIDを設定
	Answer     uint   `gorm:"not null"`
	// IsCorrect  bool      `gorm:"not null"` // Answer == 0 ならば正解
	// Timestamp time.Time `gorm:"not null"` // GROMのCreatedAt time.Timeを使用する
}
