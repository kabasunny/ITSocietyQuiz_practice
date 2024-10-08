package data // テスト用データ
import (
	"backend/src/models"
	"fmt"
)

func GenerateUsersRolesList() []models.UsersRoles {
	var usersRolesList = []models.UsersRoles{
		// adminを二つ入れておく
		{
			EmpID:  "ADM1234",
			RoleID: models.RoleAdmin, // 1
		},
		{
			EmpID:  "ADM2345",
			RoleID: models.RoleAdmin, // 1
		},
	}

	for i := 0; i < 200; i++ {
		empID := fmt.Sprintf("EMP%d", 100+i)

		userRole := models.UsersRoles{
			EmpID:  empID,
			RoleID: models.RoleUser, // 2
		}

		usersRolesList = append(usersRolesList, userRole)
	}

	return usersRolesList
}

var UsersRolesList = GenerateUsersRolesList()

// var UsersRolesList = []models.UsersRoles{
// 	{
// 		EmpID:  "ADM1234",
// 		RoleID: models.RoleAdmin, // 1
// 	},
// 	{
// 		EmpID:  "ADM2345",
// 		RoleID: models.RoleAdmin, // 1
// 	},
// 	{
// 		EmpID:  "EMP1234",
// 		RoleID: models.RoleUser, // 2
// 	},
// 	{
// 		EmpID:  "EMP2234",
// 		RoleID: models.RoleUser, // 2
// 	},
// 	{
// 		EmpID:  "EMP3234",
// 		RoleID: models.RoleUser, // 2
// 	},
// 	{
// 		EmpID:  "EMP4234",
// 		RoleID: models.RoleUser, // 2
// 	},
// 	{
// 		EmpID:  "EMP5234",
// 		RoleID: models.RoleUser, // 2
// 	},
// 	{
// 		EmpID:  "EMP6234",
// 		RoleID: models.RoleUser, // 2
// 	},
// 	{
// 		EmpID:  "EMP7234",
// 		RoleID: models.RoleUser, // 2
// 	},
// 	{
// 		EmpID:  "EMP8234",
// 		RoleID: models.RoleUser, // 2
// 	},
// 	{
// 		EmpID:  "EMP9234",
// 		RoleID: models.RoleUser, // 2
// 	},
// 	{
// 		EmpID:  "EMP1334",
// 		RoleID: models.RoleUser, // 2
// 	},
// }
