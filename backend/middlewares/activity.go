package middlewares

import "time"

// This middleware will be used for inserting an activity to the database
// when the completion of the task changes, when the title and description
// of the task changes, and when the due date of the task changes.

// 0: status change, 1: title, description change, 2: due date change

// CompareTaskByStatus compares two tasks by their completion status.
func CompareTaskByStatus(oldCompleted, newCompleted int) (bool, int) {
	if oldCompleted != newCompleted {
		return true, 0
	}
	return false, -1
}

// CompareTaskByTitleAndDescription compares two tasks by their title and description.
func CompareTaskByTitleAndDescription(oldTitle, oldDescription, newTitle, newDescription string) (bool, int) {
	if oldTitle != newTitle || oldDescription != newDescription {
		return true, 1
	}
	return false, -1
}

func CompareTaskByDueDate(oldDueDate, newDueDate time.Time) (bool, int) {
	// The conversion converts both dates to UTC
	oldDueDateUTC := oldDueDate.UTC()
	newDueDateUTC := newDueDate.UTC()

	// Comparison in years, months and days
	if oldDueDateUTC.Year() != newDueDateUTC.Year() ||
		oldDueDateUTC.Month() != newDueDateUTC.Month() ||
		oldDueDateUTC.Day() != newDueDateUTC.Day() {
		return true, 2
	}
	return false, -1
}
