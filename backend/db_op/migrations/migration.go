package main

import (
	"backend/db_op"
	"backend/infra"
	"backend/models"
	"log"
)

// マイグレーションは、独立して行うため、mainで定義している
func main() {
	infra.Initialize() //.env ファイルから環境変数を読み込み、アプリケーションにロードするための初期化処理を行う。

	db := infra.SetupDB() //データベース接続を設定し、*gorm.DB オブジェクトを返す。このオブジェクトは、データベース操作を行うためのインターフェースを提供。

	// テーブルを削除
	db_op.DropTable()

	// AutoMigrate:構造体を引数として渡し、構造体に定義されているフィールドに基づいて、データベースにテーブルを作成、更新
	if err := db.AutoMigrate(&models.QuizData{}); err != nil {
		panic("Failed to migrate database")
	}

	log.Println("Database migrated successfully!")
}
