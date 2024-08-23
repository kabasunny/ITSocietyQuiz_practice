// テスト用データ

package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// アプリケーションのデータモデルを表現し、データベースのテーブルを表現するための構造体
type QuizData struct {
	gorm.Model                // ID unit を含む構造体となっている
	Question   string         `gorm:"type:text;not null"`
	Options    pq.StringArray `gorm:"type:text[];not null"`
	Correct    string         `gorm:"type:text;not null"`
	Supplement string         `gorm:"type:text;not null"`
}
