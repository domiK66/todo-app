package controllers

import (
    http "net/http"
    gin "github.com/gin-gonic/gin"
	models "todo-backend/models"
)

func GetPrioritiesController(c *gin.Context) {
    var priorities []models.Priority
    models.DB.Find(&priorities)
    c.JSON(http.StatusOK, priorities)
}

func PopulateDatabase(c *gin.Context) {
	var pri1 = models.Priority{Name: "High"}
    models.DB.Create(&pri1)
	var pri2 = models.Priority{Name: "Low"}
    models.DB.Create(&pri2)
	var pri3 = models.Priority{Name: "Middle"}
    models.DB.Create(&pri3)
}