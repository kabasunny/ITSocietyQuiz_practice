package models

import (
	"gorm.io/gorm"
)

type AnswersDimension struct {
	gorm.Model
	EmpID            string  `gorm:"not null"`
	CorrectAnswers   int     `gorm:"not null"`
	PerformanceIndex float64 `gorm:"not null"`
}
