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
