package config

import (
	"github.com/joho/godotenv"

	"os"

	"fmt"
)

// Config func to get env value from key ---
func Config(key string) string {

	// load .env file
	err := godotenv.Load("connection.env")

	if err != nil {
		fmt.Print("Error loading connection.env file")
	}

	return os.Getenv(key)

}
