package models

import (
	"github.com/google/uuid"
)

type AddToChartRequest struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}

type CompleteOrderRequest struct {
	UserId string `json:"user_id"`
}

type GetCartRequest struct {
	UserId string `json:"user_id"`
}

type RemoveFromCartRequest struct {
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}

var (
	UserCart    Cart
	UserProduct []Product
	UserOrder   Order
)

type Order struct {
	OrderId                   uuid.UUID `json:"order_id"`
	UserId                    string    `json:"user_id"`
	Cart                      Cart      `json:"cart"`
	Discount                  float64   `json:"discount"`
	DiscountReason            string    `json:"discount_reason"`
	TotalPriceWithDiscount    float64   `json:"total_price_with_discount"`
	TotalPriceWithoutDiscount float64   `json:"total_price_without_discount"`
	Timestamp                 string    `json:"timesamp"`
	Timex                     int64     `json:"timex"`
}
