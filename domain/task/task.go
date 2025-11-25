package task

type Task struct {
	Id          string  `json:"id"`
	Title       string  `json:"title"`
	Description *string `json:"description"`
	IsCompleted bool    `json:"is_completed"`
}

func NewTask(id string, title string, description *string, isCompleted bool) *Task {
	return &Task{
		Id:          id,
		Title:       title,
		Description: description,
		IsCompleted: isCompleted,
	}
}
