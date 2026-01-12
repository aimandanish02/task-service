// declare vars and data structure of the api
package models

import "time"

type Task struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Status    TaskStatus    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

type TaskStatus string

const (
	StatusPending   TaskStatus = "pending"
	StatusProcessing TaskStatus = "processing"
	StatusCompleted  TaskStatus = "completed"
)
