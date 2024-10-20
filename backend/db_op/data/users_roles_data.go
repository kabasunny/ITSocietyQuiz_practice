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

	for i := 1; i < 200; i++ { // 社員199人分
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
