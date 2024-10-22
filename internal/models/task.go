package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	UUID      string    `json:"uuid"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (t *Task) Validate() error {
	var err error

	if t.Title == "" {
		err = errors.Join(err, ErrTitle)
	}

	if uErr := uuid.Validate(t.UUID); uErr != nil {
		err = errors.Join(err, ErrInvalidUUID)
	}

	return err
}
