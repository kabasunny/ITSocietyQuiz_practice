package models

import "gorm.io/gorm"

// アプリケーションのデータモデルを表現し、データベースのテーブルを表現するための構造体
type Users struct {
	gorm.Model            // ID unit を含む構造体となっている
	Username       string `gorm:"type:text;not null"`
	Email          string `gorm:"type:text;not null"`
	TotalQuestions int    `gorm:"type:integer;not null"`
	CorrectAnswers int    `gorm:"type:integer;not null"`
}
