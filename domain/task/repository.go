package task

type Repository interface {
	Create(t *Task) error
	FindAll() ([]Task, error)
}
