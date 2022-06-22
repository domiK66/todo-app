package controllers

import (
    http "net/http"
    gin "github.com/gin-gonic/gin"
	models "todo-backend/models"
)

func GetTasksController(c *gin.Context) {
    var tasks []models.Task
    models.DB.Find(&tasks)
    c.JSON(http.StatusOK, tasks)
}