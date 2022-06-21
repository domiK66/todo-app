package controllers

import (
    http "net/http"
    gin "github.com/gin-gonic/gin"
)

func NoRouteEndpointController(c *gin.Context) {
	c.JSON(
        http.StatusNotFound,
        gin.H {"todo-backend-api": "no route here"},
    )
}