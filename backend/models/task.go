package models

import "time"

type Task struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Title       string    `bson:"title,omitempty" json:"title"`
	Description string    `bson:"description,omitempty" json:"description"`
	Completed   int       `bson:"completed,omitempty" json:"completed"` // 0: not completed, 1: in progress, 2: completed
	CreateDate  time.Time `bson:"createDate,omitempty" json:"createDate"`
	DueDate     time.Time `bson:"dueDate,omitempty" json:"dueDate"`
}
