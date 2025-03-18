package tasks

import "time"

// Task represents a todo task
type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
