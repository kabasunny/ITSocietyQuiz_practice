package dto

type LoginInput struct {
	EmpID    string `json:"empid" binding:"required,min=6"`
	Password string `json:"password" binding:"required,min=8"`
}
