package sqls

const (

	// GetAllProducts is the query to get all products
	GetAllProducts = "SELECT * FROM %s.%s"
	// GetProduct is the query to get a product
	GetProduct = "SELECT * FROM %s.%s WHERE product_id = '%s'"
	// UpdateProduct is the query to update a product
	UpdateProductQuantity = "UPDATE %s.%s SET quantity = %d WHERE product_id = '%s'"
	// InsertOrder is the query to insert an order
	InsertOrder = "INSERT INTO %s.%s (order_id, user_id, cart, discount, total_price_with_discount, total_price_without_discount, time,timestamp) values ('%s','%s','%s',%f,%f,%f,%d,'%s')"
	// GetOrder is the query to get an order
	GetOrder = "SELECT total_price_without_discount,timestamp from %s.%s WHERE user_id = '%s' AND total_price_without_discount > %f AND timestamp LIKE %s ORDER BY timestamp ASC LIMIT 1;"
	// DeleteProduct is the query to delete a product
	DeleteProduct = "DELETE FROM %s.%s WHERE product_id = '%s'"
)
