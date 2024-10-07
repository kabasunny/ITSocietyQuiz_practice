package main

import (
	"backend/db_op/data"
	"backend/src/infra"
	"backend/src/models"
	"fmt"
	"log"
	"math/rand"
	"reflect"

	"gorm.io/gorm"
)

func main() {
	infra.Initialize() //.env ファイルから環境変数を読み込み、アプリケーションにロードするための初期化処理を行う。

	db := infra.SetupDB() //データベース接続を設定し、*gorm.DB オブジェクトを返す。このオブジェクトは、データベース操作を行うためのインターフェースを提供。

	// データの挿入順序を調整
	dataLists := [][]interface{}{
		toInterfaceSlice(data.QuestionsList),
		toInterfaceSlice(data.GenerateHashedUsersList()),
		toInterfaceSlice(data.RolesList),
	}

	// 依存関係のないデータを先に挿入
	for _, dataList := range dataLists {
		for _, data := range dataList {
			if err := db.Create(data).Error; err != nil {
				log.Printf("Failed to insert data: %v", err)
			}
		}
	}

	// 依存関係のあるUsersRolesListを最後に挿入
	for _, data := range data.GenerateUsersRolesList() {
		if err := db.Create(&data).Error; err != nil {
			log.Printf("Failed to insert UsersRoles: %v", err)
		}
	}

	// 依存関係のあるAnswersListを最後に挿入
	for _, data := range data.AnswersList {
		if err := db.Create(&data).Error; err != nil {
			log.Printf("Failed to insert Answers: %v", err)
		}
	}

	// ダミーデータの挿入
	insertDummyData(db)

	fmt.Println("Data inserted successfully!")
}

// 各データリストを[]interface{}に変換するヘルパー関数
func toInterfaceSlice(slice interface{}) []interface{} {
	v := reflect.ValueOf(slice)
	result := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		result[i] = v.Index(i).Addr().Interface()
	}
	return result
}

// ダミーデータを挿入する関数
func insertDummyData(db *gorm.DB) {
	numEmployees := 200

	for i := 0; i < numEmployees; i++ {
		empID := fmt.Sprintf("EMP%d", 100+i) // EmpIDを100から始め、"EMP"を先頭に追加
		correctAnswers := rand.Intn(601)     // 0から600の範囲でランダムな正答数を生成
		performanceIndex := float64(rand.Intn(186)+15) * (float64(correctAnswers) / float64(rand.Intn(601)))

		dimension := models.AnswersDimension{
			EmpID:            empID,
			CorrectAnswers:   correctAnswers,
			PerformanceIndex: performanceIndex,
		}

		db.Create(&dimension)
	}
}
