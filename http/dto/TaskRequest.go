package dto

type CreateTaskRequest struct {
	Id          string  `json:"id"`
	Title       string  `json:"titulo"`
	Description *string `json:"descricao"`
}
