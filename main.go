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

	config.LoadEnv()

	db, err := config.ConnectDB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := db.AutoMigrate(&models.Discount{}); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	discountRepository := repository.NewDiscountRepository(db)

	apiHandler := api.NewDiscountAPI(discountRepository)

	r := gin.Default()

	r.POST("/discounts", apiHandler.ApplyDiscount)

	if err := r.Run(":8083"); err != nil {
		log.Fatalf("Failed to start API server: %v", err)
	}
}
