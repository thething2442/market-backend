// main.go

package main

import (
	"log"
	"os"

	"market-backend/controllers" // Import the controllers package
	"market-backend/models"    // Import the models package

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// ... (Your existing database connection and migration code) ...
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Successfully connected to NeonDB!")

	err = db.AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate models: %v", err)
	}

	log.Println("Database migration complete.")

	// Create a new Gin router
	router := gin.Default()

	// -------------------------------------------------------------
	// Define all API routes
	// -------------------------------------------------------------

	// User Routes
	router.GET("/users", controllers.GetUser(db))
	router.GET("/users/:userID", controllers.GetUserById(db))

	// Product Routes
	router.GET("/products", controllers.GetProducts(db))
	router.GET("/products/:productID", controllers.GetProductById(db))

	// Cart Routes
	router.GET("/carts", controllers.GetCarts(db))
	router.GET("/carts/:cartID", controllers.GetCartById(db))

	// Run the server on port 3000
	if err := router.Run(":3000"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}