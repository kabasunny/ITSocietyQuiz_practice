package dto

// AnswersInput構造体は、クイズの回答を表現する
type AnswersInput struct {
	// EmpID      string `json:"empid" binding:"required"` // トークンからIDを抽出する
	QuestionID uint `json:"questionid" binding:"required"`
	Answer     uint `json:"answer" binding:"required"`
	// Timestamp  time.Time `json:"timestamp" binding:"required"`
}
