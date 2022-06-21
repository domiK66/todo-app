package server

import (
    gin "github.com/gin-gonic/gin"
    cors "github.com/rs/cors/wrapper/gin"
	controllers "todo-backend/controllers"
)

func setRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.Default())
	router.NoRoute(controllers.NoRouteEndpointController)
	api := router.Group("/api")
	{
		api.GET("/", controllers.DefaultEndpointController)

		// endpoints here:
		api.GET("/todo-lists", controllers.GetTodolistsController)




	}
	return router
}