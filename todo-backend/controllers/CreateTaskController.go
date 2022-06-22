package controllers

import (
    http "net/http"
    gin "github.com/gin-gonic/gin"
	models "todo-backend/models"
)

func CreateTaskController(c *gin.Context) {
    var input models.Task
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    task := models.Task {
        Title: input.Title,
        Description: input.Description,
        DueDate: input.DueDate,
        Priority: input.Priority,
    }
    models.DB.Create(&task)
    c.JSON(http.StatusOK, task)
}