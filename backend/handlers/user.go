package handlers

import "github.com/gin-gonic/gin"

// http://127.0.0.1:3000/api/ping
// {"message":"pong"}
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
