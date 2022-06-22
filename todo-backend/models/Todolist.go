package models

import (
	time "time"
)

type Todolist struct {
    ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
    Title string `json:"title"`
    Username string `json:"username"`
    Tasks []Task `json:"tasks"`
}