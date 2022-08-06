package sqls

const (

	// GetAllProducts is the query to get all products
	GetAllProducts = `SELECT * FROM %s.%s`
	// GetProduct is the query to get a product
	GetProduct = `SELECT * FROM %s.%s WHERE product_id = '%s'`
	// UpdateProduct is the query to update a product
	UpdateProduct = `UPDATE %s.%s SET quantity = %d WHERE product_id = '%s'`
)
