package handlers

import (
	"github.com/gdscduzceuniversity/todo-app-1/middlewares"
	"github.com/gdscduzceuniversity/todo-app-1/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

//•  Task
//    - Create
//    - Read - Select
//    - Update - Edit
//    - Delete
//    - Change Status (1,2,3)

// time data format RFC3339

type taskRequest struct {
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Completed   *int      `json:"completed" binding:"required"`
	DueDate     time.Time `json:"dueDate" binding:"required"`
}

// CreateTask endpoint calls InsertTask function from repository/task.go
// CreateTask godoc
// @Summary Create new task
// @Description Add a new task to the database.
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body taskRequest true "Create Task"
// @Success 200 {object} map[string]interface{} "Task created successfully"
// @Failure 400 "Failed to bind JSON"
// @Failure 401 "User not logged in"
// @Failure 500 "Internal Server Error"
// @Router /tasks [post]
func CreateTask(c *gin.Context) {

	// Check if the user is authenticated
	if auth := models.ValidateUser(c); !auth.IsAuthenticated {
		c.JSON(401, gin.H{"message": "User not logged in"})
		return
	}

	// Get the authenticated user's ID
	authUserId := models.ValidateUser(c).Id

	// Create a new request object
	var req taskRequest
	// Bind the request body to req
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"Failed to bind JSON": err.Error()})
		return
	}

	// Call CreateTask function from repository/task.go
	if err := models.InsertTask(req.Title, req.Description, authUserId, *req.Completed, req.DueDate); err != nil {
		c.JSON(500, gin.H{"Failed to create task": err.Error()})
		return
	}

	// Return a success message
	c.JSON(200, gin.H{"message": "Task created successfully"})
}

// ReadTask endpoint gets a task from the database
// ReadTask godoc
// @Summary Get task
// @Description Get details of a specific task.
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} map[string]interface{} "Task details"
// @Failure 401 "User not logged in"
// @Failure 500 "Internal Server Error"
// @Router /tasks/{id} [get]
func ReadTask(c *gin.Context) {

	// Check if the user is authenticated
	if auth := models.ValidateUser(c); !auth.IsAuthenticated {
		c.JSON(401, gin.H{"message": "User not logged in"})
		return
	}

	// Get the authenticated user's ID
	authUserId := models.ValidateUser(c).Id

	// Get the task ID from the URL
	taskId := c.Param("id")

	// Call ReadTask function from repository/task.go
	task, err := models.ReadTask(taskId)
	if err != nil {
		c.JSON(500, gin.H{"Failed to read task": err.Error()})
		return
	}

	// Check if the task belongs to the authenticated user
	if task.UserID != authUserId {
		c.JSON(401, gin.H{"message": "Unauthorized"})
		return
	}

	// Get the task's activities
	activities, err := models.ReadActivities(taskId)
	if err != nil {
		c.JSON(500, gin.H{"Failed to read activities": err.Error()})
		return
	}

	// Return the task and activities
	c.JSON(200, gin.H{
		"task":       task,
		"activities": activities})
}

// ReadTasks endpoint gets tasks by page and limit from the database
// ReadTasks godoc
// @Summary List tasks
// @Description Get a list of tasks with pagination.
// @Tags tasks
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Param limit query int false "Limit"
// @Success 200 {object} map[string]interface{} "List of tasks"
// @Failure 401 "User not logged in"
// @Failure 500 "Internal Server Error"
// @Router /tasks [get]
func ReadTasks(c *gin.Context) {

	// Check if the user is authenticated
	if auth := models.ValidateUser(c); !auth.IsAuthenticated {
		c.JSON(401, gin.H{"message": "User not logged in"})
		return
	}

	// Get the authenticated user's ID
	authUserId := models.ValidateUser(c).Id

	// Get the page and limit from the URL
	page := c.Query("page")
	limit := c.Query("limit")

	// Convert page and limit to int64
	pageInt, err := strconv.ParseInt(page, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"Failed to convert page to int64": err.Error()})
		return
	}
	limitInt, err := strconv.ParseInt(limit, 10, 64)
	if err != nil {
		c.JSON(400, gin.H{"Failed to convert limit to int64": err.Error()})
		return
	}

	// Call ReadTasks function from repository/task.go
	tasks, count, err := models.ReadTasks(authUserId, pageInt, limitInt)
	if err != nil {
		c.JSON(500, gin.H{"Failed to read tasks": err.Error()})
		return
	}

	// Return the tasks
	c.JSON(200, gin.H{"tasks": tasks, "count": count})
}

// UpdateTask endpoint calls UpdateTask function from repository/task.go
// UpdateTask godoc
// @Summary Update task
// @Description Update details of a specific task.
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body taskRequest true "Update Task"
// @Success 200 {object} map[string]interface{} "Task updated successfully"
// @Failure 400 "Failed to bind JSON"
// @Failure 401 "User not logged in"
// @Failure 500 "Internal Server Error"
// @Router /tasks/{id} [put]
func UpdateTask(c *gin.Context) {

	// Check if the user is authenticated
	if auth := models.ValidateUser(c); !auth.IsAuthenticated {
		c.JSON(401, gin.H{"message": "User not logged in"})
		return
	}

	// Get the task ID from the URL
	taskId := c.Param("id")

	// Create a new request object
	var req taskRequest
	// Bind the request body to req
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"Failed to bind JSON": err.Error()})
		return
	}

	// Call ReadTask function from repository/task.go
	oldTask, err := models.ReadTask(taskId)
	if err != nil {
		c.JSON(500, gin.H{"Failed to read task": err.Error()})
		return
	}

	// Check if the task has been changed by comparing the old and new values
	isChanged := false

	// Compare the tasks by old and new completed values
	flag, activityType := middlewares.CompareTaskByStatus(oldTask.Completed, *req.Completed)
	if flag {
		isChanged = true
		if err = models.InsertActivity(oldTask.ID, activityType, strconv.Itoa(oldTask.Completed), strconv.Itoa(*req.Completed)); err != nil {
			c.JSON(500, gin.H{"Failed to insert activity": err.Error()})
			return
		}
	}

	// Compare the tasks by old and new titles
	flag, activityType = middlewares.CompareTaskByTitleAndDescription(oldTask.Title, oldTask.Description, req.Title, req.Description)
	if flag {
		isChanged = true
		if err = models.InsertActivity(oldTask.ID, activityType, "title: "+oldTask.Title+" description: "+oldTask.Description, "title: "+req.Title+" description: "+req.Description); err != nil {
			c.JSON(500, gin.H{"Failed to insert activity": err.Error()})
			return
		}
	}

	// Compare the tasks by old and new due dates
	flag, activityType = middlewares.CompareTaskByDueDate(oldTask.DueDate, req.DueDate)
	if flag {
		isChanged = true
		if err = models.InsertActivity(oldTask.ID, activityType, oldTask.DueDate.String(), req.DueDate.String()); err != nil {
			c.JSON(500, gin.H{"Failed to insert activity": err.Error()})
			return
		}
	}

	// if the task has been changed, call UpdateTask function from repository/task.go
	if isChanged {
		if err = models.UpdateTask(taskId, req.Title, req.Description, *req.Completed, req.DueDate); err != nil {
			c.JSON(500, gin.H{"Failed to update task": err.Error()})
			return
		}
		// Return a success message
		c.JSON(200, gin.H{"message": "Task updated successfully"})
	} else {
		// Return a success message
		c.JSON(200, gin.H{"message": "Task not changed"})
	}
}

// todo: delete related activities when a task is deleted

// DeleteTask endpoint calls DeleteTask function from repository/task.go
// DeleteTask godoc
// @Summary Delete task
// @Description Delete a specific task from the database.
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {object} map[string]interface{} "Task deleted successfully"
// @Failure 401 "User not logged in"
// @Failure 500 "Internal Server Error"
// @Router /tasks/{id} [delete]
func DeleteTask(c *gin.Context) {

	// Check if the user is authenticated
	if auth := models.ValidateUser(c); !auth.IsAuthenticated {
		c.JSON(401, gin.H{"message": "User not logged in"})
		return
	}

	// Get the task ID from the URL
	taskId := c.Param("id")

	// Call DeleteTask function from repository/task.go
	if err := models.DeleteTask(taskId); err != nil {
		c.JSON(500, gin.H{"Failed to delete task": err.Error()})
		return
	}

	// Return a success message
	c.JSON(200, gin.H{"message": "Task deleted successfully"})
}
