package repository_test

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"projek/toko-retail/repository/config"
)

// Init loads environment variables from a .env file for testing
func Init() {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("env not found, using system env")
	}
}

// TestKoneksi tests the database connection
func TestKoneksi(t *testing.T) {
	Init()  // Load environment variables

	db, err := repository.OpenDB()
	if err != nil {
		t.Fatalf("Failed to connect to the database: %v", err)
	}

	// Optionally, you can add a simple query to test if the connection is working
	// Example: checking if the connection is alive
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("Failed to get raw database connection: %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("Failed to ping the database: %v", err)
	}

	t.Log("Database connection successful")
}
