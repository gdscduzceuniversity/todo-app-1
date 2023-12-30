package routes

import (
	"github.com/gdscduzceuniversity/todo-app-1/handlers"
	"github.com/gin-gonic/gin"
)

func StartServer() error {
	router := SetupRouter()
	err := router.Run(":3000")
	return err
}

/*
Select All: /tasks GET
Select: /tasks/:id GET
Create: /tasks PUT
Edit: /tasks/:id POST
Delete: /tasks/:id DELETE
Select All: /tasks?limit=5&page=1
*/

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Simple group: api
	api := router.Group("/api")
	{
		api.GET("/ping", handlers.Ping)
		api.POST("/tasks", handlers.CreateTask)
	}
	return router
}
