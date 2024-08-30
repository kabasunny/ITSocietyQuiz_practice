package models

import "gorm.io/gorm"

// アプリケーションのデータモデルを表現し、データベースのテーブルを表現するための構造体
type Users_roles struct {
	gorm.Model
	EmpID  string `gorm:"not null;constraint:OnDelete:CASCADE"`
	RoleID uint   `gorm:"not null;constraint:OnDelete:CASCADE"`
}
