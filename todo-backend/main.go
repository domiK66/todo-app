package main

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/todo-lists", getTodoLists)

    router.Run("localhost:8080")
}

type todolist struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Username string  `json:"username"`
}

var todolists = []todolist {
    {ID: "1", Name: "Blue Train", Username: "John Coltrane"},
}

func getTodoLists (c *gin.Context) {
    c.IndentedJSON(http.StatusOK, todolists)
}