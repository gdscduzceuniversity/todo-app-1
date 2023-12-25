package main

import (
	"github.com/gdscduzceuniversity/todo-app-1/db"
	"github.com/gdscduzceuniversity/todo-app-1/repository"
)

// init function to setup the database connection
func init() {
	db.Setup()
}

func main() {
	defer db.Disconnect()
	repository.InsertUser("test2", "test")
}
