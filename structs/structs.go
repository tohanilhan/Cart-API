package structs

import "github.com/google/uuid"

// Struct declarations
type DbConfig struct {
	Host             string
	Port             string
	Db               string
	Schema           string
	User             string
	Pass             string
	TableNameProduct string
	TableNameOrder   string

	SslMode string
}

type Product struct {
	ProductId   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name"`
	Price       float64   `json:"price"`
	Vat         int16     `json:"vat"`
	Quantity    int32     `json:"quantity"`
}

type Cart struct {
	ProductId   uuid.UUID `json:"product_id"`
	ProductName string    `json:"product_name"`
	Price       float64   `json:"price"`
	Vat         int16     `json:"vat"`
	Quantity    int32     `json:"quantity"`
}

type CartResponse struct {
	TotalPrice float64 `json:"total_price"`
	Cart       []Cart  `json:"cart"`
}

type Order struct {
	OrderId                   uuid.UUID `json:"order_id"`
	UserId                    uuid.UUID `json:"user_id"`
	Cart                      []Cart    `json:"cart"`
	Discount                  float64   `json:"discount"`
	TotalPriceWithDiscount    float64   `json:"total_price_with_discount"`
	TotalPriceWithoutDiscount float64   `json:"total_price_without_discount"`
	Timestamp                 string    `json:"timesamp"`
	Timex                     int64     `json:"timex"`
}
