package models

import "gorm.io/gorm"

// アプリケーションのデータモデルを表現し、データベースのテーブルを表現するための構造体
type UsersRoles struct {
	gorm.Model
	EmpID  string `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	RoleID uint   `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

// ロールIDの定数を定義
const (
	RoleAdmin = 1 // 管理者
	RoleUser  = 2 // 一般
)
