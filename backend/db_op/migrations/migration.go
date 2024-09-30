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

	db := infra.SetupDB() // データベース接続を設定し、*gorm.DB オブジェクトを返す。このオブジェクトは、データベース操作を行うためのインターフェースを提供。

	// 外部キー依存関係を考慮してテーブルを削除
	db_op.DropTable("answers")     // 削除したいテーブルを引数に渡して削除できる "questions" "users" "answers" "users_roles"
	db_op.DropTable("users_roles") // 削除したいテーブルを引数に渡して削除できる "questions" "users" "answers" "users_roles"
	db_op.DropTable("questions")   // 削除したいテーブルを引数に渡して削除できる "questions" "users" "answers" "users_roles"
	db_op.DropTable("users")       // 削除したいテーブルを引数に渡して削除できる "questions" "users" "answers" "users_roles"
	db_op.DropTable("roles")       // 削除したいテーブルを引数に渡して削除できる "questions" "users" "answers" "users_roles" "roles"

	// AutoMigrate:構造体を引数として渡し、構造体に定義されているフィールドに基づいて、データベースにテーブルを作成、更新
	// 必要な順序でテーブルを作成する
	if err := db.AutoMigrate(&models.Users{}); err != nil {
		panic("Failed to migrate Users table")
	}
	if err := db.AutoMigrate(&models.Roles{}); err != nil {
		panic("Failed to migrate Roles table")
	}
	if err := db.AutoMigrate(&models.UsersRoles{}); err != nil {
		panic("Failed to migrate UsersRoles table")
	}
	if err := db.AutoMigrate(&models.Questions{}); err != nil {
		panic("Failed to migrate Questions table")
	}
	if err := db.AutoMigrate(&models.Answers{}); err != nil {
		panic("Failed to migrate Answers table")
	}

	// 以下はコメントアウトされた個別のAutoMigrate文
	// if err := db.AutoMigrate(&models.Answers{}); err != nil {
	// if err := db.AutoMigrate(&models.Users_roles{}); err != nil {
	// if err := db.AutoMigrate(&models.Users{}); err != nil {
	// if err := db.AutoMigrate(&models.Questions{}); err != nil {
	// if err := db.AutoMigrate(&models.Roles{}); err != nil {

	log.Println("Database migrated successfully!")

	// 部分インデックスの作成、questionsテーブルのマイグレーション時に使用 (GORMで記述できないため)
	// user_question_idがNULLでない場合にユニークであることを保証
	// PostgreSQL特有の文法で、PostgreSQLのPL/pgSQL（Procedural Language/PostgreSQL）を使用して、条件付きでインデックスを作成するためのブロック
	tx := db.Exec(`
	    DO $$
	    BEGIN
	        IF NOT EXISTS (SELECT 1 FROM pg_class WHERE relname = 'unique_user_question_id') THEN
	            CREATE UNIQUE INDEX unique_user_question_id
	            ON questions(user_question_id)
	            WHERE user_question_id IS NOT NULL;
	        END IF;
	    END
	    $$;
	`)
	if tx.Error != nil {
		log.Fatal("Failed to create partial index: ", tx.Error)
	}

	// // 外部キー制約の追加（ON UPDATE CASCADE）以下はGORMで設定できたため不要

	// // UsersRolesテーブルのEmpIDに対してカスケードを追加
	// tx = db.Exec(`
	//     ALTER TABLE users_roles
	//     ADD CONSTRAINT fk_users_empid
	//     FOREIGN KEY (emp_id) REFERENCES users(emp_id)
	//     ON DELETE CASCADE
	//     ON UPDATE CASCADE;
	// `)
	// if tx.Error != nil {
	// 	log.Fatal("Failed to add foreign key with cascade for users_roles.emp_id: ", tx.Error)
	// }

	// // AnswersテーブルのEmpIDに対してカスケードを追加
	// tx = db.Exec(`
	//     ALTER TABLE answers
	//     ADD CONSTRAINT fk_answers_empid
	//     FOREIGN KEY (emp_id) REFERENCES users(emp_id)
	//     ON DELETE CASCADE
	//     ON UPDATE CASCADE;
	// `)
	// if tx.Error != nil {
	// 	log.Fatal("Failed to add foreign key with ON UPDATE CASCADE for answers.emp_id: ", tx.Error)
	// }

	// // AnswersテーブルのUserQuestionIDに対してカスケードを追加
	// tx = db.Exec(`
	//     ALTER TABLE answers
	//     ADD CONSTRAINT fk_answers_userquestionid
	//     FOREIGN KEY (user_question_id) REFERENCES questions(user_question_id)
	//     ON DELETE CASCADE
	//     ON UPDATE CASCADE;
	// `)
	// if tx.Error != nil {
	// 	log.Fatal("Failed to add foreign key with ON UPDATE CASCADE for answers.user_question_id: ", tx.Error)
	// }

	log.Println("Indexes and foreign key constraints added successfully!")
}
