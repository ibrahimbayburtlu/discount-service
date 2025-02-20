package main

import (
	"discount-service/api"
	"discount-service/config"
	"discount-service/models"
	"discount-service/repository"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	// Load environment variables
	config.LoadEnv()

	// Connect to the database
	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Run database migrations
	if err := db.AutoMigrate(&models.Discount{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	// Initialize repository
	discountRepository := repository.NewDiscountRepository(db)

	// Initialize API Handler
	apiHandler := api.NewDiscountAPI(discountRepository)

	// Create Gin router
	r := gin.Default()

	// Define API routes
	r.POST("/discounts", apiHandler.ApplyDiscount)

	// Start the server
	if err := r.Run(":8083"); err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}
