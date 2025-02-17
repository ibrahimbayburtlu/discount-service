package repository

import (
	"discount-service/models"

	"gorm.io/gorm"
)

type DiscountRepository interface {
	Save(discount *models.Discount) error
	GetByCustomerID(customerID uint) ([]models.Discount, error)
}

type discountRepositoryImpl struct {
	db *gorm.DB
}

func NewDiscountRepository(db *gorm.DB) DiscountRepository {
	return &discountRepositoryImpl{db: db}
}

func (r *discountRepositoryImpl) Save(discount *models.Discount) error {
	return r.db.Create(discount).Error
}

func (r *discountRepositoryImpl) GetByCustomerID(customerID uint) ([]models.Discount, error) {
	var discounts []models.Discount
	err := r.db.Where("customer_id = ?", customerID).Find(&discounts).Error
	return discounts, err
}
