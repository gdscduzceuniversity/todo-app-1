package repository

import (
	"context"
	"fmt"
	"github.com/gdscduzceuniversity/todo-app-1/db"
	"github.com/gdscduzceuniversity/todo-app-1/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const databaseName = "todo-database"
const collectionName = "users"

// InsertUser inserts a user to the database
func InsertUser(username, password string) {
	// Create a unique index on the "username" field
	indexModel := mongo.IndexModel{
		Keys:    bson.D{{"username", 1}},
		Options: options.Index().SetUnique(true),
	}

	// Get the collection
	coll := db.Client.Database(databaseName).Collection(collectionName)

	// Create the index
	_, err := coll.Indexes().CreateOne(context.Background(), indexModel)
	if err != nil {
		panic(err)
	}

	// Create a new user
	newUser := models.User{Username: username, Password: password}

	// Insert the user to the database
	result, err := coll.InsertOne(context.TODO(), newUser)
	if err != nil {
		panic(err)
	}
	// Print the inserted user's ID
	fmt.Println("Inserted a user, id:", result.InsertedID)
}
