package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

// Questions構造体は、クイズの質問を表現する
type Questions struct {
	gorm.Model                    // IDフィールドを含む構造体
	UserQuestionID *string        `gorm:"type:text"` // ユーザー管理用の質問ID, NULLを許容 Goの仕様上stringはnilを許容しないので*とする
	Question       string         `gorm:"type:text;not null"`
	Options        pq.StringArray `gorm:"type:text[];not null"` // Options[0]を正解とする
	Supplement     string         `gorm:"type:text;not null"`
	Difficulty     uint           `gorm:"type:integer;not null"`
}
