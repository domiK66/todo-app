package controllers

import (
    http "net/http"
    gin "github.com/gin-gonic/gin"
	models "todo-backend/models"
)

func GetTodolistsController(c *gin.Context) {
    c.Header("Content-Type", "application/json")
    c.IndentedJSON(http.StatusOK, todolists)
}


//TODO: database - just for testing
var priorities = []models.Priority {
    {Id: "1", Name: "High" },
    {Id: "2", Name: "Mid" },
    {Id: "3", Name: "Low" },
}

var tasks = []models.Task {
    {Id:"1", Title: "Creating an Application", Description: "smth", DueDate: "15.06.2022", Priority: priorities[0]},
    {Id:"2", Title: "Creating an Application1", Description: "smth", DueDate: "15.06.2022", Priority: priorities[0]},
    {Id:"3", Title: "Creating an Application2", Description: "smth", DueDate: "15.06.2022", Priority: priorities[0]},

}

var todolists = []models.Todolist {
    {Id: "1", Title: "Work styria", Username: "kainzdom", Tasks: tasks },
	{Id: "2", Title: "Work direktbad", Username: "kainzdom", Tasks: tasks },
    {Id: "3", Title: "Private", Username: "kainzdom", Tasks: tasks },
	{Id: "4", Title: "FH", Username: "kainzdom", Tasks: tasks },
}