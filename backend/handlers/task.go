package handlers

import (
	"github.com/gdscduzceuniversity/todo-app-1/repository"
	"github.com/gin-gonic/gin"
	"time"
)

//•  Task
//    - Create
//    - Read - Select
//    - Update - Edit
//    - Delete
//    - Change Status (1,2,3)

type createTaskRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Completed   *int      `json:"completed" binding:"required"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
}

// CreateTask endpoint calls CreateTask function from repository/task.go
func CreateTask(c *gin.Context) {

	// todo: add user session validation

	// Create a new request object
	var req createTaskRequest
	// Bind the request body to req
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"Failed to bind JSON": err.Error()})
		return
	}

	// Call CreateTask function from repository/task.go
	if err := repository.InsertTask(req.Title, req.Description, *req.Completed, req.DueDate); err != nil {
		c.JSON(500, gin.H{"Failed to create task": err.Error()})
		return
	}

	// Return a success message
	c.JSON(200, gin.H{"message": "Task created successfully"})
}
