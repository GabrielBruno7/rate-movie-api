package handlers

import (
	"crud/domain/task"
	"crud/infrastructure/database"
	"database/sql"
	"encoding/json"
	"net/http"
)

type TaskHandler struct {
	repo task.Repository
}

func NewTaskHandler(databaseConnection *sql.DB) *TaskHandler {
	return &TaskHandler{
		repo: database.NewTaskDb(databaseConnection),
	}
}

func (h *TaskHandler) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := h.repo.FindAll()
	if err != nil {
		http.Error(w, "Erro ao listar tasks", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(tasks)
}

func (h *TaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Title string `json:"title"`
	}

	json.NewDecoder(r.Body).Decode(&data)

	t := task.NewTask(data.Title)

	err := h.repo.Create(t)
	if err != nil {
		http.Error(w, "Erro ao criar task", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(201)
	w.Write([]byte("Task criada"))
}
