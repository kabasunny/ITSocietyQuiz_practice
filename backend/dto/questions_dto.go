package dto

type CreateQuestionsInput struct {
	Question   string   `json:"question" binding:"required"`
	Options    []string `json:"options" binding:"required"`
	Supplement string   `json:"supplement" binding:"required"`
	Difficulty int      `json:"difficulty" binding:"required"` // 追加
}

type UpdateQuestionsInput struct {
	Question   *string   `json:"question" binding:"omitempty"`
	Options    *[]string `json:"options" binding:"omitempty"`
	Supplement *string   `json:"supplement" binding:"omitempty"`
	Difficulty *int      `json:"difficulty" binding:"omitempty"` // 追加
}

// とりあえず作成、テスト前に要確認
type QuizData struct {
	Question   string   `json:"question"`
	Options    []string `json:"options"`
	Supplement string   `json:"supplement"`
	Difficulty int      `json:"difficulty"`
}

// とりあえず作成、テスト前に要確認
type QuizResult struct {
	QuestionID int    `json:"question_id"`
	UserAnswer int    `json:"user_answer"`
	Correct    bool   `json:"correct"`
	Supplement string `json:"supplement"`
}
