package models

import "gorm.io/gorm"

// アプリケーションのデータモデルを表現し、データベースのテーブルを表現するための構造体
type Users struct {
	gorm.Model                  // ID unit を含む構造体となっている
	EmpID          string       `gorm:"type:text;not null;unique"` // employeeId を追加
	Username       string       `gorm:"type:text;not null"`
	Email          string       `gorm:"type:text;not null"`
	Password       string       `gorm:"type:text;not null"`
	CurrentQID     uint         `gorm:"not null;default:0"` // 現在の最も進捗した問題の番号
	TotalQuestions int          `gorm:"type:integer;not null"`
	CorrectAnswers int          `gorm:"type:integer;not null"`
	UsersRoles     []UsersRoles `gorm:"foreignKey:EmpID;references:EmpID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // リレーションにOnUpdate:CASCADEとOnDelete:CASCADEを追加
	Answers        []Answers    `gorm:"foreignKey:EmpID;references:EmpID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"` // Answersのリレーションを追加
}
