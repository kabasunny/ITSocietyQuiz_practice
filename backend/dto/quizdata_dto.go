package dto

type CreateQuizDataInput struct {
	Question   string   `json:"question" binding:"required"`
	Options    []string `json:"options" binding:"required"`
	Correct    string   `json:"correct" binding:"required"`
	Supplement string   `json:"supplement" binding:"required"`
}

type UpdateQuizDataInput struct {
	Question   *string   `json:"question" binding:"omitempty"`
	Options    *[]string `json:"options" binding:"omitempty"`
	Correct    *string   `json:"correct" binding:"omitempty"`
	Supplement *string   `json:"supplement" binding:"omitempty"`
}
