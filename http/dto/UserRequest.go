package dto

type CreateUserRequest struct {
	Id       string `json:"id"`
	Name     string `json:"nome"`
	Email    string `json:"email"`
	Password string `json:"senha"`
}
