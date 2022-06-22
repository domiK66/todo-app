package models

import (
	time "time"
)

type Task struct {
    ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
    Title string `json:"title"`
    Description string `json:"description"`
    DueDate string `json:"due_date"`
    Priority Priority `json:"priority"`
}