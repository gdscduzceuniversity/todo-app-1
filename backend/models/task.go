package models

import (
	"context"
	"fmt"
	"github.com/gdscduzceuniversity/todo-app-1/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Task struct {
	ID          string    `bson:"_id,omitempty" json:"id"`
	UserID      string    `bson:"userID,omitempty" json:"userID"`
	Title       string    `bson:"title,omitempty" json:"title"`
	Description string    `bson:"description,omitempty" json:"description"`
	Completed   int       `bson:"completed,omitempty" json:"completed"` // 0: to-do, 1: in progress, 2: completed
	CreateDate  time.Time `bson:"createDate,omitempty" json:"createDate"`
	DueDate     time.Time `bson:"dueDate,omitempty" json:"dueDate"`
}

const taskCollection = "tasks"
const databaseName = "todo-database"

//•  Task
//    - Create
//    - Read - Select
//    - Update - Edit
//    - Delete
//    - Change Status (1,2,3)

// InsertTask creates a task in the database
func InsertTask(title, description, UserID string, completed int, dueDate time.Time) error {

	// Get the collection
	coll := db.Client.Database(databaseName).Collection(taskCollection)

	// Create a new task
	doc := Task{Title: title, Description: description, Completed: completed, CreateDate: time.Now(), DueDate: dueDate}

	// Insert the task to the database
	result, err := coll.InsertOne(context.Background(), doc)
	if err != nil {
		return fmt.Errorf("error while inserting task: %w", err)
	}
	// Print the inserted task's ID
	fmt.Println("Inserted a task, id:", result.InsertedID)
	return nil
}

// ReadTask reads a task from the database
func ReadTask(id string) (Task, error) {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(taskCollection)

	// Create a filter
	filter := bson.D{{"_id", id}}

	// Find one document
	var result Task
	err := coll.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return Task{}, fmt.Errorf("error while getting task: %w", err)
	}

	return result, nil
}

// ReadTasks reads tasks by pagination from the database
func ReadTasks(page, limit int64) ([]Task, int64, error) {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(taskCollection)

	// Create a filter
	filter := bson.D{{}}

	// Find documents
	var results []Task
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip((page - 1) * limit)
	cur, err := coll.Find(context.Background(), filter, findOptions)
	if err != nil {
		return []Task{}, 0, fmt.Errorf("error while getting tasks: %w", err)
	}
	for cur.Next(context.Background()) {
		var result Task
		err = cur.Decode(&result)
		if err != nil {
			return []Task{}, 0, fmt.Errorf("error while getting tasks: %w", err)
		}
		results = append(results, result)
	}

	// return the number of all tasks
	count, err := coll.CountDocuments(context.Background(), filter)
	if err != nil {
		return []Task{}, 0, fmt.Errorf("error while getting tasks: %w", err)
	}
	fmt.Println("Number of tasks:", count)

	return results, count, nil
}

// UpdateTaskStatus updates a task's completed status in the database
func UpdateTaskStatus(id string, completed int) error {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(taskCollection)

	// Create a filter
	filter := bson.D{{"_id", id}}

	// Create an update
	update := bson.D{{"$set", bson.D{{"completed", completed}}}}

	// Update the task in the database
	result, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error while updating task: %w", err)
	}
	// Print the number of updated documents
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return nil
}

// UpdateTaskTitle updates a task's title in the database
func UpdateTaskTitle(id, title string) error {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(taskCollection)

	// Create a filter
	filter := bson.D{{"_id", id}}

	// Create an update
	update := bson.D{{"$set", bson.D{{"title", title}}}}

	// Update the task in the database
	result, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error while updating task: %w", err)
	}
	// Print the number of updated documents
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return nil
}

// UpdateTaskDescription updates a task's description in the database
func UpdateTaskDescription(id, description string) error {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(taskCollection)

	// Create a filter
	filter := bson.D{{"_id", id}}

	// Create an update
	update := bson.D{{"$set", bson.D{{"description", description}}}}

	// Update the task in the database
	result, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error while updating task: %w", err)
	}
	// Print the number of updated documents
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return nil
}

// UpdateTaskDueDate updates a task's due date in the database
func UpdateTaskDueDate(id string, dueDate time.Time) error {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(taskCollection)

	// Create a filter
	filter := bson.D{{"_id", id}}

	// Create an update
	update := bson.D{{"$set", bson.D{{"dueDate", dueDate}}}}

	// Update the task in the database
	result, err := coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("error while updating task: %w", err)
	}
	// Print the number of updated documents
	fmt.Printf("Updated %v Documents!\n", result.ModifiedCount)
	return nil
}

// DeleteTask deletes a task from the database
func DeleteTask(id string) error {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(taskCollection)

	// Create a filter
	filter := bson.D{{"_id", id}}

	// Delete the task from the database
	result, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error while deleting task: %w", err)
	}
	// Print the number of deleted documents
	fmt.Printf("Deleted %v Documents!\n", result.DeletedCount)
	return nil
}
