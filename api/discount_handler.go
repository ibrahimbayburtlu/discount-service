package api

import (
	"discount-service/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DiscountAPI struct {
	service *service.DiscountService
}

func NewDiscountAPI(service *service.DiscountService) *DiscountAPI {
	return &DiscountAPI{service: service}
}

func (api *DiscountAPI) GetCustomerDiscounts(c *gin.Context) {
	customerID, err := strconv.Atoi(c.Param("customerID"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Geçersiz müşteri ID"})
		return
	}

	discounts, err := api.service.GetDiscountsByCustomerID(uint(customerID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "İndirimler getirilemedi"})
		return
	}

	c.JSON(http.StatusOK, discounts)
}
