package models

import (
	"time"

	"gorm.io/gorm"
)

// アプリケーションのデータモデルを表現し、データベースのテーブルを表現するための構造体
type Answers struct {
	gorm.Model           // ID unit を含む構造体となっている
	UserID     uint      `gorm:"not null;constraint:OnDelete:CASCADE"`
	QuestionID uint      `gorm:"not null;constraint:OnDelete:CASCADE"`
	Answer     int       `gorm:"not null"`
	IsCorrect  bool      `gorm:"not null"`
	Timestamp  time.Time `gorm:"not null"`
}
