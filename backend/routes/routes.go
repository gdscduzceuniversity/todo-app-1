package routes

import (
	docs "github.com/gdscduzceuniversity/todo-app-1/docs"
	"github.com/gdscduzceuniversity/todo-app-1/handlers"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

	docs.SwaggerInfo.BasePath = "/api"
	// Simple group: api
	api := router.Group("/api")
	{
		// Auth Routes
		api.POST("/register", handlers.Register)
		api.POST("/login", handlers.Login)
		api.POST("/logout", handlers.Logout)
		//api.GET("/user", handlers.User)

		// Task Routes
		api.POST("/tasks", handlers.CreateTask)
		api.GET("/tasks/:id", handlers.ReadTask)
		api.GET("/tasks", handlers.ReadTasks)
		api.PUT("/tasks/:id", handlers.UpdateTask)
		api.DELETE("/tasks/:id", handlers.DeleteTask)

	}
	// Swagger Route for connect to swagger ui go to http://127.0.0.1:3000/swagger/index.html
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	return router
}
