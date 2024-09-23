package dto

// フロントへのユーザー用クイズデータ返却用
type QuizData struct {
	ID         uint     `json:"id"` // GROMのIDを使用する
	Question   string   `json:"question"`
	Options    []string `json:"options"`
	Supplement string   `json:"supplement"`
	Difficulty uint     `json:"difficulty"`
}

// フロントへのクイズデータ返却用
type AdmQuizData struct {
	ID             uint     `json:"id"` // GORMのIDを使用する
	UserQuestionID *string  `json:"userQuestionID"`
	Question       string   `json:"question"`
	Options        []string `json:"options"`
	Supplement     string   `json:"supplement"`
	Difficulty     uint     `json:"difficulty"`
	CreatedAt      string   `json:"createdAt"` // 作成日を追加
}

type UpdateQuestionsInput struct {
	ID             uint      `json:"id"` // GORMのIDを使用する
	UserQuestionID *string   `json:"userQuestionID"`
	Question       *string   `json:"question" binding:"omitempty"`
	Options        *[]string `json:"options" binding:"omitempty"`
	Supplement     *string   `json:"supplement" binding:"omitempty"`
	Difficulty     *uint     `json:"difficulty" binding:"omitempty"` // 追加
}

type CreateQuestionsInput struct {
	ID             uint     `json:"id"` // GORMのIDを使用する
	UserQuestionID *string  `json:"userQuestionID"`
	Question       string   `json:"question" binding:"required"`
	Options        []string `json:"options" binding:"required"`
	Supplement     string   `json:"supplement" binding:"required"`
	Difficulty     uint     `json:"difficulty" binding:"required"` // 追加
}
