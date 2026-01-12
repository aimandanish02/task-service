//go's standard db abstraction
package repository

import (
	"database/sql"
	"errors"

	_ "github.com/mattn/go-sqlite3"
	"task-service/pkg/models"
)

type SQLiteTaskRepository struct {
	db *sql.DB
}

func NewSQLiteRepository(dbPath string) (*SQLiteTaskRepository, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS tasks (
		id TEXT PRIMARY KEY,
		title TEXT,
		status TEXT,
		created_at DATETIME
	);
	`

	if _, err := db.Exec(query); err != nil {
		return nil, err
	}

	return &SQLiteTaskRepository{db: db}, nil
}

func (r *SQLiteTaskRepository) Create(task models.Task) error {
	_, err := r.db.Exec(
		"INSERT INTO tasks (id, title, status, created_at) VALUES (?, ?, ?, ?)",
		task.ID,
		task.Title,
		task.Status,
		task.CreatedAt,
	)
	return err
}

func (r *SQLiteTaskRepository) GetByID(id string) (models.Task, error) {
	row := r.db.QueryRow(
		"SELECT id, title, status, created_at FROM tasks WHERE id = ?",
		id,
	)

	var task models.Task
	if err := row.Scan(&task.ID, &task.Title, &task.Status, &task.CreatedAt); err != nil {
		if err == sql.ErrNoRows {
			return task, errors.New("task not found")
		}
		return task, err
	}

	return task, nil
}

func (r *SQLiteTaskRepository) Update(task models.Task) error {
	_, err := r.db.Exec(
		"UPDATE tasks SET status = ? WHERE id = ?",
		task.Status,
		task.ID,
	)
	return err
}
