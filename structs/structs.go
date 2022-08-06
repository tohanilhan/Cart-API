package structs

// Struct declarations
type DbConfig struct {
	Host      string
	Port      string
	Db        string
	Schema    string
	User      string
	Pass      string
	TableName string
	SslMode   string
}

type Product struct {
	ProductId   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float32 `json:"price"`
	Vat         int32   `json:"vat"`
	Quantity    int32   `json:"quantity"`
}

type Cart struct {
	ProductId   string  `json:"product_id"`
	ProductName string  `json:"product_name"`
	Price       float32 `json:"price"`
	Vat         int32   `json:"vat"`
	Quantity    int32   `json:"quantity"`
}

type CartResponse struct {
	TotalPrice float32 `json:"total_price"`
	Discount   float32 `json:"discount"`
	Cart       []Cart  `json:"cart"`
}
