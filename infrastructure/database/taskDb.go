package database

import (
	"crud/domain/task"
	"database/sql"
)

type TaskDb struct {
	db *sql.DB
}

func NewTaskDb(db *sql.DB) *TaskDb {
	return &TaskDb{db: db}
}

func (r *TaskDb) Create(t *task.Task) error {
	query := `INSERT INTO tasks (title, description, is_completed) VALUES ($1, $2, $3)`
	_, err := r.db.Exec(query, t.Title, t.Description, t.IsCompleted)
	return err
}

func (r *TaskDb) FindAll() ([]task.Task, error) {
	rows, err := r.db.Query("SELECT id, title, description, is_completed FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []task.Task

	for rows.Next() {
		var t task.Task
		err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.IsCompleted)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}

	return tasks, nil
}
