package dto

type CreateUserRequest struct {
	Id       string `json:"id"`
	Name     string `json:"nome" binding:"required,max=100"`
	Email    string `json:"email" binding:"required,email,max=100"`
	Password string `json:"senha" binding:"required,min=6,max=64"`
}
