package database

import "github.com/tohanilhan/Cart-API/app/queries"

// Queries struct for collect all app queries.
type Queries struct {
	*queries.CartApiQueries // load queries from User model
}

// GetDBConnection func for opening database connection.
func GetDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.

	return &Queries{
		// Set queries from models:
		CartApiQueries: &queries.CartApiQueries{DB: Db}, // from UserJob model
	}, nil
}
