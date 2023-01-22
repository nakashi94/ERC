package config

import (
	"fmt"
	"os"
)

func GetDBURL() (string, error) {
	dbUser := os.Getenv("USER_NAME")
	dbPassword := os.Getenv("PASSWORD")
	dbHost := os.Getenv("HOST")
	dbPort := os.Getenv("PORT")
	dbDatabaseName := os.Getenv("DATABASE_NAME")
	if dbUser == "" || dbPassword == "" || dbHost == "" || dbPort == "" || dbDatabaseName == "" {
		return "", fmt.Errorf("Error: required environment variable not found")
	}
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPassword, dbHost, dbPort, dbDatabaseName), nil
}
