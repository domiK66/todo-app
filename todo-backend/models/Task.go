package models

type Task struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    DueDate string `json:"due_date"`
    Priority Priority `json:"priority"`
}