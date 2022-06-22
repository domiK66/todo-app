package controllers

import (
	http "net/http"
	models "todo-backend/models"
	gin "github.com/gin-gonic/gin"
)

type inputJson struct {
    ID  uint `json:"id"`
    Title string `json:"title"`
    Description string `json:"description"`
    DueDate string `json:"due_date"`
    PriorityID uint `json:"priority_id"`
}

func CreateTaskController(c *gin.Context) {
    var input inputJson
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    var prio models.Priority
    models.DB.First(&prio,input.PriorityID)

    task := models.Task {
        Title: input.Title,
        Description: input.Description,
        DueDate: input.DueDate,
        Priority: prio,
    }
    models.DB.Create(&task)
    c.JSON(http.StatusOK, task)
}