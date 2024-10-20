package main

import (
	"backend/db_op/data"
	"backend/src/infra"
	"backend/src/models"
	"fmt"
	"log"
	"math/rand"
	"reflect"
	"time"

	"gorm.io/gorm"
)

func main() {
	infra.Initialize() //.env ファイルから環境変数を読み込み、アプリケーションにロードするための初期化処理を行う。

	db := infra.SetupDB() //データベース接続を設定し、*gorm.DB オブジェクトを返す。このオブジェクトは、データベース操作を行うためのインターフェースを提供。

	// データの挿入順序を調整
	dataLists := [][]interface{}{
		toInterfaceSlice(data.QuestionsList),
		toInterfaceSlice(data.UsersList),
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
	for _, data := range data.GenerateAnswersList() {
		if err := db.Create(&data).Error; err != nil {
			log.Printf("Failed to insert Answers: %v", err)
		}
	}

	// Answers_dimensionテーブルにダミーデータの挿入
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

// Answers_dimensionテーブルにダミーデータを挿入する関数
func insertDummyData(db *gorm.DB) {
	numEmployees := 200
	startDate := time.Date(2024, 4, 1, 0, 0, 0, 0, time.Local)
	endDate := time.Date(2024, 9, 30, 0, 0, 0, 0, time.Local)
	day := 0

	for date := startDate; date.Before(endDate) || date.Equal(endDate); date = date.AddDate(0, 0, 1) {
		// 休日を除外
		if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
			continue
		}

		day++
		empIDs := make([]string, numEmployees)
		empIDs[0] = "AVG100" // iが0の時はempIDを"AVG100"に設定
		for i := 1; i < numEmployees; i++ {
			empIDs[i] = fmt.Sprintf("EMP%d", 100+i) // EmpIDを100から始め、"EMP"を先頭に追加
		}

		for i := 0; i < numEmployees; i++ {
			empID := empIDs[i]

			// dayが0の場合を考慮して、最低値を1に設定
			correctAnswers := rand.Intn(max(1, day/2+1)) + int(day/2) + 3
			performanceIndex := float64(rand.Intn(max(1, day/2+1))+int(day/2)+3) * (float64(correctAnswers) / float64(day*4-(rand.Intn(5))+rand.Intn(max(1, day/10))+1))

			dimension := models.AnswersDimension{
				EmpID:            empID,
				CorrectAnswers:   correctAnswers,
				PerformanceIndex: performanceIndex,
			}

			// CreatedAtフィールドに日付を設定
			dimension.CreatedAt = date

			db.Create(&dimension)
		}
	}

}

// max関数を追加して、引数が0にならないようにする
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
