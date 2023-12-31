package models

import (
	"context"
	"fmt"
	"github.com/gdscduzceuniversity/todo-app-1/db"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

type Activity struct {
	ID     string    `bson:"_id,omitempty" json:"id"`
	Time   time.Time `bson:"time,omitempty" json:"time"`
	TaskID string    `bson:"taskID,omitempty" json:"taskID"`
	Type   int       `bson:"type,omitempty" json:"type"` // 0: status change, 1: title, description change, 2: due date change
	Old    string    `bson:"old,omitempty" json:"old"`
	New    string    `bson:"new,omitempty" json:"new"`
}

const activityCollection = "activities"

// InsertActivity inserts an activity to the database
func InsertActivity(taskID string, activityType int, old, new string) {
	// Get the collection
	coll := db.Client.Database(databaseName).Collection(activityCollection)

	// Create a new activity
	newActivity := Activity{Time: time.Now(), TaskID: taskID, Type: activityType, Old: old, New: new}

	// Insert the activity to the database
	result, err := coll.InsertOne(context.Background(), newActivity)
	if err != nil {
		panic(err)
	}
	// Print the inserted activity's ID
	fmt.Println("Inserted an activity, id:", result.InsertedID)
}

// ReadActivities reads activities by taskID from the database
func ReadActivities(taskID string) ([]Activity, error) {
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
	var activities []Activity
	for cursor.Next(context.Background()) {
		var activity Activity
		if err = cursor.Decode(&activity); err != nil {
			return nil, fmt.Errorf("error while decoding activity: %w", err)
		}
		activities = append(activities, activity)
	}

	return activities, nil
}
