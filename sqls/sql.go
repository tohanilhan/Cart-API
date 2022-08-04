package sqls

const (
	// CreateTableQuery is the query to create table
	CreateTableQuery = `CREATE TABLE IF NOT EXISTS pf_schema.products
	(
		product_id               uuid primary key,
		product_name             varchar(200),
		price                    float,
		vat                      varchar(4),
		total                    int2
	);`

	// GetAllProductsQuery is the query to get all products
	GetAllProductsQuery = `SELECT * FROM %s.%s`
	
	

)
