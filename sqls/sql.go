package sqls

const (

	// GetAllProducts is the query to get all products
	GetAllProducts = "SELECT * FROM %s.%s"
	// GetProduct is the query to get a product
	GetProduct = "SELECT * FROM %s.%s WHERE product_id = '%s'"
	// UpdateProduct is the query to update a product
	UpdateProductQuantity = "UPDATE %s.%s SET quantity = %d WHERE product_id = '%s'"
	// InsertOrder is the query to insert an order
	InsertOrder = "INSERT INTO %s.%s (order_id, user_id, cart, discount, total_price_with_discount, total_price_without_discount, time,timestamp,discount_reason) values ('%s','%s','%s',%f,%f,%f,%d,'%s','%s')"
	// GetOrder is the query to get an order
	GetOrder = "SELECT COALESCE(SUM (total_price_without_discount),0) from %s.%s WHERE user_id = '%s' AND timestamp LIKE %s "
	// DeleteProduct is the query to delete a product
	DeleteProduct = "DELETE FROM %s.%s WHERE product_id = '%s'"
)
