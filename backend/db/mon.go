package db

import (
	"context"
	"fmt"

	"github.com/gdscduzceuniversity/todo-app-1/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client         *mongo.Client
	Context        context.Context
	CancelFunc     context.CancelFunc
	UserCollection *mongo.Collection
)

// Setup mongodb initial connection code block
func Setup() {
	// Use the SetServerAPIOptions() method to set the Stable API version to 1
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(config.Config("MONGODB_URI")).SetServerAPIOptions(serverAPI)

	var err error
	// Create a new Client and connect to the server
	Client, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	// Send a ping to confirm a successful connection
	if err = Client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		panic(err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	getCollections()
	UserCollection.Database().Client().Ping(Context, nil)
}

func getCollections() {
	databaseName := "todo-database"
	database := Client.Database(databaseName)

	UserCollection = database.Collection("users")
}

func Disconnect() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	fmt.Println("Connection to MongoDB closed.")
}
