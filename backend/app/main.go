package main

import (
	repository "projek/toko-retail/repository/config" // Import the repository package for database configuration
	"projek/toko-retail/routes"                       // Import the routes package for setting up API routes

	"github.com/gofiber/fiber/v2" // Import Fiber for creating and managing HTTP server
	"github.com/joho/godotenv"    // Import godotenv for loading environment variables
	"github.com/sirupsen/logrus"  // Import logrus for logging
)

// InitEnv initializes environment variables from a .env file
func InitEnv() {
	err := godotenv.Load("../.env") // Load environment variables from the .env file
	if err != nil {                 // Check if there was an error loading the .env file
		logrus.Warn("Cannot load env file, using system env") // Log a warning if the .env file could not be loaded
	}
}

func main() {
	// Initialize environment variables
	InitEnv()

	// Open a connection to the database
	repository.OpenDB()

	// Create a new Fiber application instance
	app := fiber.New()

	// Setup API routes
	routes.RouteSetup(app)

	// Start the Fiber server and listen on port 3000
	err := app.Listen(":3000") // Listen on port 3000
	if err != nil {            // Check if there was an error starting the server
		logrus.Fatal(
			"Error on running fiber, ",
			err.Error()) // Log a fatal error if the server fails to start
	}
}
