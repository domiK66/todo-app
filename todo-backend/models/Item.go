package models

import (
	time "time"
)
type Item struct {
    ID          uint       `json:"id"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
    Name string `json:"name"`
    Quantity string `json:"quantity"`
}