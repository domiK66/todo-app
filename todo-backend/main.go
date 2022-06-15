package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
    cors "github.com/rs/cors/wrapper/gin"
)

func main() {
    router := gin.Default()
    router.Use(cors.Default())

    api := router.Group("/api")
    {
        api.GET("/", func(context *gin.Context) {
            context.JSON(http.StatusOK, gin.H{"api": "is working"})
        })

        api.GET("/todo-lists", getTodoLists)
    }
    
    router.Run("localhost:3001")
}

type Priority struct {
    Id string `json:"id"`
    Name string `json:"name"`
}

type Task struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    DueDate string `json:"due_date"`
    Priority Priority `json:"priority"`
}

type Todolist struct {
    Id string `json:"id"`
    Title string `json:"title"`
    Username string `json:"username"`
    Tasks []Task `json:"tasks"`
}

var priorities = []Priority {
    {Id: "1", Name: "High" },
    {Id: "2", Name: "Mid" },
    {Id: "3", Name: "Low" },
}

var tasks = []Task {
    {Id:"1", Title: "Creating an Application", Description: "smth", DueDate: "15.06.2022", Priority: priorities[0]},
}

var todolists = []Todolist {
    {Id: "1", Title: "Work", Username: "kainzdom", Tasks: tasks },
}


func getTodoLists (c *gin.Context) {
    c.Header("Content-Type", "application/json")
    c.IndentedJSON(http.StatusOK, todolists)
}


//TODO
type Item struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Quantity string `json:"quantity"`
}