package store

import "todo-app/tasks"

type Store interface {
	InitDB() (bool, error)
	ListTasks() ([]tasks.Task, error)
	CreateTask(title, description string) (int64, error)
	UpdateTask(id int, title, description string) (int64, error)
	CompleteTask(id int) (int64, error)
	UndoTask(id int) (int64, error)
	DeleteTask(id int) (int64, error)
	DeleteCompletedTasks() (int64, error)
}
