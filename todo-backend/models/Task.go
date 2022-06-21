package models

type Task struct {
    Id string `json:"id" gorm:"primary_key"`
    Title string `json:"title"`
    Description string `json:"description"`
    DueDate string `json:"due_date"`
    Priority Priority `json:"priority"`
}