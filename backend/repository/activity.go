package repository

import (
	"context"
	"fmt"
	"github.com/gdscduzceuniversity/todo-app-1/db"
	"github.com/gdscduzceuniversity/todo-app-1/models"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

//•  Activity
//    - Save to activity collection when the completion of the task changes  (0: to-do, 1: in progress, 2: completed)
//    - Save to activity collection when the title and description of the task changes
//    - Save to activity collection when the due date of the task changes

const activityCollection = "activities"

// InsertActivity inserts an activity to the database
func InsertActivity(taskID string, activityType int, old, new string) {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(activityCollection)

	// Create a new activity
	newActivity := models.Activity{Time: time.Now(), TaskID: taskID, Type: activityType, Old: old, New: new}

	// Insert the activity to the database
	result, err := coll.InsertOne(context.Background(), newActivity)
	if err != nil {
		panic(err)
	}
	// Print the inserted activity's ID
	fmt.Println("Inserted an activity, id:", result.InsertedID)
}

// ReadActivities reads activities by taskID from the database
func ReadActivities(taskID string) ([]models.Activity, error) {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(activityCollection)

	// Create a filter
	filter := bson.D{{"taskID", taskID}}

	// Find documents
	cursor, err := coll.Find(context.Background(), filter)
	if err != nil {
		return nil, fmt.Errorf("error while getting activities: %w", err)
	}

	// Iterate through the returned cursor.
	var activities []models.Activity
	for cursor.Next(context.Background()) {
		var activity models.Activity
		if err = cursor.Decode(&activity); err != nil {
			return nil, fmt.Errorf("error while decoding activity: %w", err)
		}
		activities = append(activities, activity)
	}

	return activities, nil
}
