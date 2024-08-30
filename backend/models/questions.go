package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Questions構造体は、クイズの質問を表現する
type Questions struct {
	gorm.Model                // IDフィールドを含む構造体
	QuestionID string         `gorm:"unique"`
	Question   string         `gorm:"type:text;not null"`
	Options    pq.StringArray `gorm:"type:text[];not null"` // Options[0]を正解とする
	Supplement string         `gorm:"type:text;not null"`
	Difficulty int            `gorm:"type:integer;not null"`
}
