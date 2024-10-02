package data // テスト用データ
import "backend/src/models"

var UsersRolesList = []models.UsersRoles{
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
