package models

import (
	"context"
	"fmt"
	"github.com/gdscduzceuniversity/todo-app-1/db"
	"go.mongodb.org/mongo-driver/bson/primitive"

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

// todo bug fix: unable to save completed field to database

// InsertTask creates a task in the database
func InsertTask(title, description, UserID string, completed int, dueDate time.Time) error {

	// Get the collection
	coll := db.Client.Database(databaseName).Collection(taskCollection)

	// Create a new task
	doc := Task{Title: title, Description: description, UserID: UserID, Completed: completed, CreateDate: time.Now(), DueDate: dueDate}

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

	// Convert string ID to ObjectId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return Task{}, fmt.Errorf("error converting ID to ObjectId: %w", err)
	}

	// Create a filter
	filter := bson.D{{"_id", objID}}

	// Find one document
	var result Task
	err = coll.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return Task{}, fmt.Errorf("error while getting task: %w", err)
	}

	return result, nil
}

// ReadTasks reads tasks by pagination from the database for a specific userId
func ReadTasks(userId string, page, limit int64) ([]Task, int64, error) {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(taskCollection)

	// Create a filter to match the userId
	filter := bson.D{{"userID", userId}}

	// Find documents with pagination
	var results []Task
	findOptions := options.Find()
	findOptions.SetLimit(limit)
	findOptions.SetSkip((page - 1) * limit)
	cur, err := coll.Find(context.Background(), filter, findOptions)
	if err != nil {
		return []Task{}, 0, fmt.Errorf("error while getting tasks: %w", err)
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		var result Task
		err = cur.Decode(&result)
		if err != nil {
			return []Task{}, 0, fmt.Errorf("error while getting tasks: %w", err)
		}
		results = append(results, result)
	}

	// Return the number of tasks for the specific userId
	count, err := coll.CountDocuments(context.Background(), filter)
	if err != nil {
		return []Task{}, 0, fmt.Errorf("error while getting tasks: %w", err)
	}

	return results, count, nil
}

// UpdateTask updates various fields of a task in the database.
// It takes task ID and the fields to be updated as parameters.
func UpdateTask(id string, title string, description string, completed int, dueDate time.Time) error {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(taskCollection)

	// Convert string ID to ObjectId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting ID to ObjectId: %w", err)
	}

	// Create a filter to match the task ID
	filter := bson.D{{"_id", objID}}

	// Create an update
	update := bson.D{
		{"$set", bson.D{
			{"title", title},
			{"description", description},
			{"completed", completed},
			{"dueDate", dueDate},
		}},
	}

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

	// Convert string ID to ObjectId
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return fmt.Errorf("error converting ID to ObjectId: %w", err)
	}

	// Create a filter with ObjectId
	filter := bson.D{{"_id", objID}}

	// Delete the task from the database
	result, err := coll.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("error while deleting task: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("error while deleting task: %w", err)
	}

	// Print the number of deleted documents
	fmt.Printf("Deleted %v Documents!\n", result.DeletedCount)
	return nil
}
