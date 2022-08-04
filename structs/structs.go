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
	ProductId   string `json:"product_id"`
	ProductName string `json:"product_name"`
	Price       string `json:"price"`
	Vat         string `json:"vat"`
	Total       string `json:"total"`
}
