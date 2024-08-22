package db_op

import (
	"backend/infra"
	"log"
)

// テーブルを削除する関数
func DropTable() {
	infra.Initialize() //.env ファイルから環境変数を読み込み、アプリケーションにロードするための初期化処理を行う。

	db := infra.SetupDB() //データベース接続を設定し、*gorm.DB オブジェクトを返す。このオブジェクトは、データベース操作を行うためのインターフェースを提供。

	// テーブルが存在するかチェックしてから削除
	if db.Migrator().HasTable("quiz_data") {
		if err := db.Migrator().DropTable("quiz_data"); err != nil {
			log.Fatalf("Failed to drop table: %v", err)
		}
		log.Println("Table dropped successfully!")
	} else {
		log.Println("Table does not exist, no need to drop.")
	}
}
