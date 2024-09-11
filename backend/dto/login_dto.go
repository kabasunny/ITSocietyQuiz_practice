package dto

// フロントからログイン情報の受取用
type LoginInput struct {
	EmpID    string `json:"empid" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=8"`
}

// フロントにトークンと管理者フラグの返却用
type LoginResponse struct {
	Token string `json:"token"`
	Admin bool   `json:"admin"`
}
