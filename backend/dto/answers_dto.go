package dto

// AnswersInput構造体は、クイズの回答を表現する
type AnswersInput struct {
	// EmpID      string `json:"empid" binding:"required"` // トークンからIDを抽出する
	QuestionID uint `json:"question_id"`
	AnswerID   uint `json:"answer_id"`
	// Timestamp  time.Time `json:"timestamp" binding:"required"`

	// 以下のbinding:"required"では 0 を受け付けなくなるの
	// QuestionID uint json:"question_id" binding:"required"
	// AnswerID   uint json:"answer_id" binding:"required"
}
