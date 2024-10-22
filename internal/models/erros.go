package models

import "errors"

var (
	ErrTitle       = errors.New("title is required")
	ErrInvalidUUID = errors.New("invalid uuid")
	ErrNotFound    = errors.New("not found")
)
