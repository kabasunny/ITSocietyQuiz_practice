package data // テスト用データ
import "backend/models"

var UsersRolesList = []models.Users_roles{
	{
		EmpID:  "EMP1234",
		RoleID: models.RoleUser, // 2
	},
	{
		EmpID:  "EMP2345",
		RoleID: models.RoleUser, // 2
	},
	{
		EmpID:  "ADM1234",
		RoleID: models.RoleAdmin, // 1
	},
}