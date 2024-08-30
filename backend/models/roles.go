package models

import "gorm.io/gorm"

// アプリケーションのデータモデルを表現し、データベースのテーブルを表現するための構造体
type Roles struct {
	gorm.Model               // ID unit を含む構造体となっている
	RoleID     uint          `gorm:"not null;unique"`
	RoleName   string        `gorm:"not null"`
	UsersRoles []Users_roles `gorm:"foreignKey:RoleID;references:RoleID"` // UsersRolesへの値は指定しない。GORMがリレーションを自動的に処理
}
