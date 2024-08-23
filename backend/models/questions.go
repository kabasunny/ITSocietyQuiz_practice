package models

import "gorm.io/gorm"

// アプリケーションのデータモデルを表現し、データベースのテーブルを表現するための構造体
type Questions struct {
	gorm.Model          // ID unit を含む構造体となっている
	Question   string   `gorm:"type:text;not null"`
	Options    []string `gorm:"type:text[];not null"` //Options[0]を正解とする
	Supplement string   `gorm:"type:text;not null"`
	Difficulty int      `gorm:"type:integer;not null"`
}
