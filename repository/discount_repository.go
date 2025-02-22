package repository

import (
	"discount-service/models"
	"log"

	"gorm.io/gorm"
)

type DiscountRepository struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) *DiscountRepository {
	return &DiscountRepository{db: db}
}

func (r *DiscountRepository) Save(discount *models.Discount) error {
	result := r.db.Debug().Create(discount)

	if result.Error != nil {
		log.Printf("Discount save failed: %v", result.Error)
		return result.Error
	}

	log.Printf("Discount saved: %+v", discount)
	return nil
}
