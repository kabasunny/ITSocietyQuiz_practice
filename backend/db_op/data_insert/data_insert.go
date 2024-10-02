package main

import (
	"backend/db_op/data"
	"backend/src/infra"
	"fmt"
	"log"
	"reflect"
)

func main() {
	infra.Initialize() //.env ファイルから環境変数を読み込み、アプリケーションにロードするための初期化処理を行う。

	db := infra.SetupDB() //データベース接続を設定し、*gorm.DB オブジェクトを返す。このオブジェクトは、データベース操作を行うためのインターフェースを提供。

	// insertData := data.AnswersList // data.～を切り替えて、データ挿入。
	// insertData := data.QuestionsList // data.～を切り替えて、データ挿入。
	// insertData := data.GetHashedUsersList() // data.～を切り替えて、データ挿入。ユーザーの登録時はパスワードをハッシュ化するので、関数を呼ぶ
	// insertData := data.RolesList // data.～を切り替えて、データ挿入。
	// insertData := data.UsersRolesList // data.～を切り替えて、データ挿入。

	// for _, data := range insertData {
	// 	if err := db.Create(&data).Error; err != nil {
	// 		log.Printf("Failed to insert quiz: %v", err)
	// 	}
	// }

	// 以下は一括で行うとき

	// データの挿入順序を調整
	dataLists := [][]interface{}{
		// toInterfaceSlice関数を使用して、各データリストを[]interface{}型に変換
		toInterfaceSlice(data.QuestionsList),
		toInterfaceSlice(data.GetHashedUsersList()),
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

	// 依存関係のあるUsersRolesListを最後に挿入 UsersやRolesにcascadeしている
	for _, data := range data.UsersRolesList {
		if err := db.Create(&data).Error; err != nil { //引数はアドレスで
			log.Printf("Failed to insert UsersRoles: %v", err)
		}
	}

	// 依存関係のあるAnswersListを最後に挿入
	for _, data := range data.AnswersList {
		if err := db.Create(&data).Error; err != nil { //引数はアドレスで
			log.Printf("Failed to insert Answers: %v", err)
		}
	}

	fmt.Println("Data inserted successfully!")
}

// 各データリストを[]interface{}に変換するヘルパー関数
func toInterfaceSlice(slice interface{}) []interface{} {
	v := reflect.ValueOf(slice)
	result := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		// ポインタを設定
		result[i] = v.Index(i).Addr().Interface()
	}
	return result
}
