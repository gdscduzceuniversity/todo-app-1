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
func SetupRouter() *gin.Engine {
	router := gin.New()
	router.SetTrustedProxies(nil)
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// Simple group: api
	api := router.Group("/api")
	{
		api.GET("/ping", handlers.Ping)
		//api.POST("/login", loginEndpoint)

		// Auth Routes

		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		api.POST("/logout", handlers.Logout)

	}
	return router
}
