package main

import (
	"fmt"
	"github.com/gdscduzceuniversity/todo-app-1/db"
	"github.com/gdscduzceuniversity/todo-app-1/routes"
	"os"
)

// init function to setup the database connection
func init() {
	db.Setup()
}

func main() {
	// Defer the database disconnect
	defer db.Disconnect()

	// Start the server
	err := routes.StartServer()
	if err != nil {
		fmt.Errorf("error starting server: %w", err)
		os.Exit(1)
	}
}
