package models

type Todolist struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Username string `json:"username"`
    Tasks []Task `json:"tasks"`
}