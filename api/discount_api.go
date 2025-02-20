package api

import (
	"discount-service/models"
	"discount-service/repository"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type DiscountAPI struct {
	discountRepo *repository.DiscountRepository
}

// Constructor function
func NewDiscountAPI(discountRepo *repository.DiscountRepository) *DiscountAPI {
	return &DiscountAPI{discountRepo: discountRepo}
}

// Apply discount function
func (api *DiscountAPI) ApplyDiscount(c *gin.Context) {
	var request models.DiscountRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		log.Printf("🚨 Invalid JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON format"})
		return
	}

	// Validate required fields
	customerTier := strings.ToLower(request.CustomerTier)
	if customerTier == "" {
		log.Println("🚨 Missing customerTier in request")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Customer tier is required"})
		return
	}

	if request.CustomerID == 0 || request.OrderID == 0 || request.AmountBeforeDiscount <= 0 {
		log.Println("🚨 Invalid input data:", request)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid customer ID, order ID, or amount"})
		return
	}

	// Get discount percentage
	discountPercent := getDiscountByTier(customerTier)
	discountAmount := request.AmountBeforeDiscount * (discountPercent / 100)
	amountAfterDiscount := request.AmountBeforeDiscount - discountAmount

	// Save discount to DB
	discount := models.Discount{
		CustomerID:           request.CustomerID,
		OrderID:              request.OrderID,
		DiscountPercent:      discountPercent,
		AmountBeforeDiscount: request.AmountBeforeDiscount,
		AmountAfterDiscount:  amountAfterDiscount,
	}

	if err := api.discountRepo.Save(&discount); err != nil {
		log.Printf("🚨 Failed to save discount: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to apply discount"})
		return
	}

	log.Println("✅ Discount successfully saved!")

	// Return response
	c.JSON(http.StatusOK, discount)
}

// Helper function to get discount percentage
func getDiscountByTier(tier string) float64 {
	switch strings.ToLower(tier) {
	case "gold":
		return 10.0
	case "platinum":
		return 20.0
	default:
		return 0.0
	}
}
