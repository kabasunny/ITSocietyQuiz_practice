package data // テスト用データ
import "backend/models"

var RolesList = []models.Roles{
	{
		RoleID:   models.RoleAdmin, // 1
		RoleName: "クイズ管理者",
	},
	{
		RoleID:   models.RoleUser, // 2
		RoleName: "一般",
	},
}

// ロールIDの定数を別途定義している。以下は参考
// const (
// 	RoleAdmin = 1 // 管理者
// 	RoleUser  = 2 // 一般
// )
