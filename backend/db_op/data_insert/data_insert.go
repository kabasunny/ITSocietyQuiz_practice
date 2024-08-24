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

	insertData := data.QuestionsList // data.～を切り替えて、データ挿入

	for _, quiz := range insertData {
		if err := db.Create(&quiz).Error; err != nil {
			log.Printf("Failed to insert quiz: %v", err)
		}
	}

	fmt.Println("Data inserted successfully!")
}
