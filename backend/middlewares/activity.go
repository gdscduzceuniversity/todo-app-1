package middlewares

import "github.com/gdscduzceuniversity/todo-app-1/models"

// This middleware will be used for inserting an activity to the database
// when the completion of the task changes, when the title and description
// of the task changes, and when the due date of the task changes.

// 0: status change, 1: title, description change, 2: due date change

// CompareTaskByStatus compares two tasks by their completion status.
func CompareTaskByStatus(oldTask, newTask models.Task) (bool, int) {
	if oldTask.Completed != newTask.Completed {
		return true, 0
	}
	return false, -1
}

// CompareTaskByTitleAndDescription compares two tasks by their title and description.
func CompareTaskByTitleAndDescription(oldTask, newTask models.Task) (bool, int) {
	if oldTask.Title != newTask.Title || oldTask.Description != newTask.Description {
		return true, 1
	}
	return false, -1
}

// CompareTaskByDueDate compares two tasks by their due date.
func CompareTaskByDueDate(oldTask, newTask models.Task) (bool, int) {
	if oldTask.DueDate != newTask.DueDate {
		return true, 2
	}
	return false, -1
}
