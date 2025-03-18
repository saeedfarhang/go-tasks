package store

import (
	"database/sql"
	"log"
	"todo-app/tasks"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteStore struct {
	db *sql.DB
}

func NewSQLiteStore(dbPath string) (*SQLiteStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	store := &SQLiteStore{db: db}
	_, err = store.InitDB()
	if err != nil {
		return nil, err
	}

	return store, nil
}

func (s *SQLiteStore) InitDB() (bool, error) {
	query := `CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		description TEXT,
		completed BOOLEAN,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
		updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := s.db.Exec(query)
	return err == nil, err
}

func (s *SQLiteStore) ListTasks() ([]tasks.Task, error) {
	rows, err := s.db.Query("SELECT id, title, description, completed, created_at, updated_at FROM tasks")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasksList []tasks.Task
	for rows.Next() {
		var t tasks.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Description, &t.Completed, &t.CreatedAt, &t.UpdatedAt); err != nil {
			log.Fatal(err)
			return nil, err
		}
		tasksList = append(tasksList, t)
	}
	return tasksList, nil
}

func (s *SQLiteStore) CreateTask(title, description string) (int64, error) {
	result, err := s.db.Exec("INSERT INTO tasks (title, description, completed) VALUES (?, ?, 0)", title, description)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (s *SQLiteStore) UpdateTask(id int, title, description string) (int64, error) {
	result, err := s.db.Exec("UPDATE tasks SET title = ?, description = ? WHERE id = ?", title, description, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *SQLiteStore) CompleteTask(id int) (int64, error) {
	result, err := s.db.Exec("UPDATE tasks SET completed = 1, updated_at = CURRENT_TIMESTAMP WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *SQLiteStore) UndoTask(id int) (int64, error) {
	result, err := s.db.Exec("UPDATE tasks SET completed = 0, updated_at = CURRENT_TIMESTAMP WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *SQLiteStore) DeleteTask(id int) (int64, error) {
	result, err := s.db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

func (s *SQLiteStore) DeleteCompletedTasks() (int64, error) {
	result, err := s.db.Exec("DELETE FROM tasks WHERE completed = 1")
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
