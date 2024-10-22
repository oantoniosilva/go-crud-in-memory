package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/oantoniosilva/go-crud-in-memory/internal/models"
	"github.com/oantoniosilva/go-crud-in-memory/internal/repositories"
)

type Task struct {
	taskRepository repositories.Task
}

func NewTask(taskRepository repositories.Task) *Task {
	return &Task{
		taskRepository: taskRepository,
	}
}

func (handler *Task) Show(w http.ResponseWriter, r *http.Request) {
	taskUUID := r.PathValue("uuid")

	task, err := handler.taskRepository.Get(taskUUID)
	if err != nil {
		log.Println(err)

		if errors.Is(err, models.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func (handler *Task) Register(w http.ResponseWriter, r *http.Request) {
	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task.UUID = uuid.NewString()
	if err := task.Validate(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	handler.taskRepository.Register(task)

	res, err := json.Marshal(task)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

func (handler *Task) Update(w http.ResponseWriter, r *http.Request) {
	taskUUID := r.PathValue("uuid")

	var task models.Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	task.UUID = taskUUID
	if err := handler.taskRepository.Update(task); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	res, err := json.Marshal(task)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(res)
}

func (handler *Task) Destroy(w http.ResponseWriter, r *http.Request) {
	taskUUID := r.PathValue("uuid")
	handler.taskRepository.Delete(taskUUID)
	w.WriteHeader(http.StatusNoContent)
}
