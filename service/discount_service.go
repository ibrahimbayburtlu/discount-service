package service

import (
	"discount-service/models"
	"discount-service/repository"
	"log"
)

type DiscountService struct {
	DiscountRepo repository.DiscountRepository
}

func NewDiscountService(repo repository.DiscountRepository) *DiscountService {
	return &DiscountService{DiscountRepo: repo}
}

func (s *DiscountService) CalculateDiscount(customerTier string) float64 {
	switch customerTier {
	case "Gold":
		return 0.10
	case "Platinum":
		return 0.20
	default:
		return 0.0
	}
}

func (s *DiscountService) GetDiscountsByCustomerID(customerID uint) ([]models.Discount, error) {
	return s.DiscountRepo.GetByCustomerID(customerID)
}

func (s *DiscountService) ApplyDiscount(customerID uint, orderID uint, amount float64, customerTier string) (models.Discount, error) {
	discountRate := s.CalculateDiscount(customerTier)
	discountAmount := amount * discountRate
	finalAmount := amount - discountAmount

	discount := models.Discount{
		CustomerID:           customerID,
		OrderID:              orderID,
		DiscountPercent:      discountRate * 100,
		AmountBeforeDiscount: amount,
		AmountAfterDiscount:  finalAmount,
	}

	// 📌 İndirim veritabanına kaydediliyor
	err := s.DiscountRepo.Save(&discount)
	if err != nil {
		log.Printf("Veritabanına indirim kaydedilemedi: %v", err)
		return discount, err
	}

	log.Printf("İndirim başarıyla uygulandı: %+v", discount)
	return discount, nil
}
