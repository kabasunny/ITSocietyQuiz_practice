package models

import "gorm.io/gorm"

// アプリケーションのデータモデルを表現し、データベースのテーブルを表現するための構造体
type Roles struct {
	gorm.Model        // ID unit を含む構造体となっている
	RoleID     uint   `gorm:"not null;constraint:OnDelete:CASCADE"`
	RoleName   string `gorm:"not null"`
}
