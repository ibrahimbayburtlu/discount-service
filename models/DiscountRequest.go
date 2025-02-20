package models

type DiscountRequest struct {
	CustomerID           uint    `json:"customerID" binding:"required"`
	OrderID              uint    `json:"orderID" binding:"required"`
	CustomerTier         string  `json:"customerTier" binding:"required"`
	AmountBeforeDiscount float64 `json:"amountBeforeDiscount" binding:"required,gt=0"`
}
