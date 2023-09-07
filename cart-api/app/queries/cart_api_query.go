package queries

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/jmoiron/sqlx"
	"github.com/tohanilhan/Cart-API/app/models"
)

type CartApiQueries struct {
	*sqlx.DB
}

// GetProductById func for get product by id.
func (q *CartApiQueries) GetProductById(productId string) (models.Product, error) {

	var product models.Product

	query := "SELECT * FROM ecomm_schema.products WHERE product_id = $1;"

	stmt, err := q.Preparex(query)
	if err != nil {
		return product, err
	}

	err = stmt.Get(&product, productId)
	if err != nil {
		return product, err
	}

	return product, nil
}

// GetProducts func for get all products.
func (q *CartApiQueries) GetProducts() ([]models.Product, error) {

	var products []models.Product

	query := "SELECT * FROM ecomm_schema.products;"

	stmt, err := q.Preparex(query)
	if err != nil {
		return products, err
	}

	rows, err := stmt.Query()
	if err != nil {
		return products, err
	}

	// iterate over rows
	for rows.Next() {
		var product models.Product
		err := rows.Scan(
			&product.ProductId,
			&product.ProductName,
			&product.Price,
			&product.Vat,
			&product.Quantity,
		)
		if err != nil {
			return products, err
		}
		// if product quantity is greater than 0, then add product to products array
		if product.Quantity > 0 {
			products = append(products, product)
		}

	}
	return products, nil
}

// CreateCart func for create cart in db
func (q *CartApiQueries) CreateCart(cart models.Cart, cartJson []byte) error {

	query := "INSERT INTO ecomm_schema.carts VALUES ($1, $2, $3, $4, $5);"

	stmt, err := q.Preparex(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		cart.CartId,
		cart.UserId,
		cartJson,
		cart.TotalPrice,
		cart.Status,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetCartByUserId func for get chart by user id
func (q *CartApiQueries) GetCartByUserId(userID string) (models.Cart, error) {

	var cartFromDb models.CartFromDb
	var cart models.Cart

	query := "SELECT * FROM ecomm_schema.carts WHERE user_id = $1 and status = 'NOT_COMPLETED';"

	stmt, err := q.Preparex(query)
	if err != nil {
		return cart, err
	}

	err = stmt.Get(&cartFromDb, userID)
	if err != nil {
		return cart, nil
	}

	// convert string to json
	err = json.Unmarshal([]byte(cartFromDb.Products), &cart)
	if err != nil {
		return cart, err
	}

	return cart, nil
}

// UpdateCart func for update cart in db
func (q *CartApiQueries) UpdateCart(cart models.Cart, cartJson []byte) error {

	query := "UPDATE ecomm_schema.carts SET products = $1, total_price = $2 WHERE user_id = $3;"

	stmt, err := q.Preparex(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(cartJson, cart.TotalPrice, cart.UserId)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCartAsCompleted func for update cart as completed
func (q *CartApiQueries) UpdateCartAsCompleted(userID string) error {

	query := "UPDATE ecomm_schema.carts SET status = 'COMPLETED' WHERE user_id = $1;"

	stmt, err := q.Preparex(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(userID)
	if err != nil {
		return err
	}

	return nil
}

func (q *CartApiQueries) SaveOrder(order models.Order, orderJson []byte) error {

	query := "INSERT INTO ecomm_schema.orders VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);"

	stmt, err := q.Preparex(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(
		order.OrderId,
		order.UserId,
		orderJson,
		order.Discount,
		order.DiscountReason,
		order.TotalPriceWithDiscount,
		order.TotalPriceWithoutDiscount,
		order.Timex,
		order.Timestamp,
	)

	if err != nil {
		return err
	}

	return nil

}

func (q *CartApiQueries) DecreaseQuantity(productId string, quantityCount int) error {

	quantityCountString := strconv.Itoa(quantityCount)

	query := fmt.Sprintf("UPDATE ecomm_schema.products SET quantity = quantity - %s WHERE product_id = $1;", quantityCountString)

	stmt, err := q.Preparex(query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(productId)
	if err != nil {
		return err
	}

	return nil
}

func (q *CartApiQueries) GetTotalOrderCount(userID string) (int, error) {
	count := 0
	query := "SELECT COUNT(*) FROM ecomm_schema.orders where user_id=$1"

	stmt, err := q.Preparex(query)
	if err != nil {
		return 0, err
	}

	err = stmt.Get(&count, userID)
	if err != nil {
		return 0, err
	}

	return count, nil
}
