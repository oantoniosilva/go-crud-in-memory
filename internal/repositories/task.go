package repositories

import "github.com/oantoniosilva/go-crud-in-memory/internal/models"

type Task interface {
	Get(taskUUID string) (models.Task, error)
	Register(task models.Task) models.Task
	Update(task models.Task) error
	Delete(taskUUID string)
}
