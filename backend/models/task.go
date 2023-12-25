package models

import "time"

type Task struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	Title       string    `bson:"title,omitempty" json:"title"`
	Description string    `bson:"description,omitempty" json:"description"`
	Completed   bool      `bson:"completed,omitempty" json:"completed"`
	CreateDate  time.Time `bson:"createDate,omitempty" json:"createDate"`
	DueDate     time.Time `bson:"dueDate,omitempty" json:"dueDate"`
}
