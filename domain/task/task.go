package task

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool   `json:"is_completed"`
}

func NewTask(title string) *Task {
	return &Task{
		Title:       title,
		Description: "",
		IsCompleted: false,
	}
}
