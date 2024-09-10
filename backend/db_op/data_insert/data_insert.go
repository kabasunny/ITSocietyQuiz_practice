package main

import (
	"backend/db_op/data"
	"backend/infra"
	"fmt"
	"log"
)

func main() {
	infra.Initialize() //.env ファイルから環境変数を読み込み、アプリケーションにロードするための初期化処理を行う。

	db := infra.SetupDB() //データベース接続を設定し、*gorm.DB オブジェクトを返す。このオブジェクトは、データベース操作を行うためのインターフェースを提供。

	// insertData := data.AnswersList // data.～を切り替えて、データ挿入。
	// insertData := data.QuestionsList // data.～を切り替えて、データ挿入。
	insertData := data.GetHashedUsersList() // data.～を切り替えて、データ挿入。ユーザーの登録時はパスワードをハッシュ化するので、関数を呼ぶ

	for _, data := range insertData {
		if err := db.Create(&data).Error; err != nil {
			log.Printf("Failed to insert quiz: %v", err)
		}
	}

	fmt.Println("Data inserted successfully!")
}
