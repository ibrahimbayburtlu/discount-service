package models

import "time"

type Discount struct {
	ID                   uint    `gorm:"primaryKey"`
	CustomerID           uint    `gorm:"index"`
	OrderID              uint    `gorm:"index"`
	DiscountPercent      float64 `gorm:"type:decimal(5,2)"`
	AmountBeforeDiscount float64 `gorm:"type:decimal(10,2)"`
	AmountAfterDiscount  float64 `gorm:"type:decimal(10,2)"`
	CreatedAt            time.Time
}
