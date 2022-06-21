package controllers

import (
    http "net/http"
    gin "github.com/gin-gonic/gin"
)

func DefaultEndpointController(c *gin.Context) {
	c.JSON(
        http.StatusOK,
        gin.H{"todo-backend-api": "default endpoint"},
    )
}