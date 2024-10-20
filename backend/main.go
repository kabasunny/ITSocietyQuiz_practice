package main

import (
	"backend/src/infra"
	"backend/src/router"
)

func main() {
	infra.Initialize() // .envファイルから環境変数を読み込む
	// log.Println("POSTGRES_USER=", os.Getenv("POSTGRES_USER")) //環境変数POSTGRES_USERの値を取得し、ログに表示
	db := infra.SetupDB()
	r := router.SetupRouter(db)
	r.Run("localhost:8082")
}
