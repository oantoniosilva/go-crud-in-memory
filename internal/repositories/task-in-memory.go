package repositories

import (
	"sync"

	"github.com/oantoniosilva/go-crud-in-memory/internal/models"
)

type TaskInMemory struct {
	tasks map[string]models.Task
	mu    sync.RWMutex
}

func NewTaskInMemory() Task {
	return &TaskInMemory{
		tasks: make(map[string]models.Task),
	}
}

func (t *TaskInMemory) Get(taskUUID string) (models.Task, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()

	task, found := t.tasks[taskUUID]
	if !found {
		return models.Task{}, models.ErrNotFound
	}

	return task, nil
}

func (t *TaskInMemory) Register(task models.Task) models.Task {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.tasks[task.UUID] = task

	return task
}

func (t *TaskInMemory) Update(task models.Task) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	_, found := t.tasks[task.UUID]
	if !found {
		return models.ErrNotFound
	}

	t.tasks[task.UUID] = task

	return nil
}

func (t *TaskInMemory) Delete(taskUUID string) {
	t.mu.Lock()
	defer t.mu.Unlock()
	delete(t.tasks, taskUUID)
}
