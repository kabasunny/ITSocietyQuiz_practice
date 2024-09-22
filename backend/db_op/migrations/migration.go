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

	// db_op.DropTable("answers") // 削除したいテーブルを引数に渡して削除できる "questions" "users" "answers" "users_roles"
	// db_op.DropTable("users")   // 削除したいテーブルを引数に渡して削除できる "questions" "users" "answers" "users_roles"
	db_op.DropTable("questions") // 削除したいテーブルを引数に渡して削除できる "questions" "users" "answers" "users_roles"
	// db_op.DropTable("users_roles") // 削除したいテーブルを引数に渡して削除できる "questions" "users" "answers" "users_roles"

	// AutoMigrate:構造体を引数として渡し、構造体に定義されているフィールドに基づいて、データベースにテーブルを作成、更新
	// if err := db.AutoMigrate(&models.Questions{}, &models.Users{}, &models.Answers{}, &models.Users_roles{}); err != nil {
	// if err := db.AutoMigrate(&models.Answers{}); err != nil {
	// if err := db.AutoMigrate(&models.Users_roles{}); err != nil {
	// if err := db.AutoMigrate(&models.Users{}, &models.Answers{}); err != nil {
	if err := db.AutoMigrate(&models.Questions{}); err != nil {
		panic("Failed to migrate database")
	}

	log.Println("Database migrated successfully!")

	// 部分インデックスの作成、questionsテーブルのマイグレーション時に使用
	// tx := db.Exec(`
	//     DO $$
	//     BEGIN
	//         IF NOT EXISTS (SELECT 1 FROM pg_class WHERE relname = 'unique_user_question_id') THEN
	//             CREATE UNIQUE INDEX unique_user_question_id
	//             ON questions(user_question_id)
	//             WHERE user_question_id IS NOT NULL;
	//         END IF;
	//     END
	//     $$;
	// `)
	// if tx.Error != nil {
	// 	log.Fatal("Failed to create partial index: ", tx.Error)
	// }
}
