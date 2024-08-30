package dto

// AnswersInput構造体は、クイズの回答を表現する
type AnswersInput struct {
	EmpID      string `json:"empid" binding:"required"`
	QuestionID uint   `json:"questionid" binding:"required"`
	Answer     int    `json:"answer" binding:"required"`
	// Timestamp  time.Time `json:"timestamp" binding:"required"`
}
