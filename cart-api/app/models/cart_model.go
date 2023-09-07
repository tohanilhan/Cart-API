package models

import (
	"fmt"
	"math"
)

type Cart struct {
	CartId     string    `json:"cart_id"`
	Products   []Product `json:"products"`
	TotalPrice float64   `json:"total_price"`
	UserId     string    `json:"user_id"`
	Status     string    `json:"status"`
}

type CartFromDb struct {
	CartId     string  `db:"cart_id" json:"cart_id"`
	Products   string  `db:"products" json:"products"`
	TotalPrice float64 `db:"total_price" json:"total_price"`
	UserId     string  `db:"user_id" json:"user_id"`
	Status     string  `db:"status" json:"status"`
}

type Product struct {
	ProductId   string  `db:"product_id" json:"product_id"`
	ProductName string  `db:"product_name" json:"product_name"`
	Price       float64 `db:"price" json:"price"`
	Vat         int16   `db:"vat" json:"vat"`
	Quantity    int32   `db:"quantity" json:"quantity"`
}

// AddProduct adds product to cart. It checks if product is already in cart. If it is, it increases quantity of product.
func (c *Cart) AddProduct(storedProduct Product) error {

	// to store product details
	var product Product

	// if there is not any product in cart, add new product to cart
	if len(UserCart.Products) == 0 {

		product.Quantity = 1
		product.Price = storedProduct.Price
		product.ProductId = storedProduct.ProductId
		product.ProductName = storedProduct.ProductName
		product.Vat = storedProduct.Vat
		UserCart.TotalPrice += storedProduct.Price

		// add product to cart
		c.Products = append(c.Products, product)

	} else {
		var isInCart = false

		// check if product is already in cart
		for i, v := range c.Products {

			if v.ProductId == storedProduct.ProductId {

				if v.Quantity >= storedProduct.Quantity {
					return fmt.Errorf("Product is out of stock")
				}

				// add product to cart if stock is available
				c.Products[i].Quantity++
				isInCart = true

			}

			// end loop if product is in cart
			if isInCart {
				break
			}
		}

		// if product is not in cart, add new product to cart
		if !isInCart {

			// add new product to cart
			product.Quantity = 1
			product.Price = storedProduct.Price
			product.ProductId = storedProduct.ProductId
			product.ProductName = storedProduct.ProductName
			product.Vat = storedProduct.Vat
			c.Products = append(c.Products, product)
		}

	}

	// calculate total price
	c.CalculateTotalPrice()
	c.Status = "NOT_COMPLETED"
	return nil
}

// RemoveItem removes product from cart. It checks if product is in cart. If it is, it decreases quantity of product.
func (c *Cart) RemoveItem(productId string) error {

	// check if product is in models.UserCart
	for i, item := range c.Products {

		if item.ProductId == productId {

			if item.Quantity > 1 {
				c.Products[i].Quantity--
			} else {
				// remove product from models.UserCartResponse
				c.Products = append(c.Products[:i], c.Products[i+1:]...)

			}

			// calculate total price
			c.CalculateTotalPrice()

			return nil
		}

	}
	return fmt.Errorf("Product is not in cart")
}

// CalculateTotalPrice calculates total price of cart. It iterates over items and calculates total price.
func (c *Cart) CalculateTotalPrice() {

	// Calculate total price
	totalPrice := 0.0
	totalVat := 0.0
	for _, cartItem := range c.Products {
		totalPrice += cartItem.Price * float64(cartItem.Quantity)
		totalVat += cartItem.Price * float64(cartItem.Quantity) * float64(cartItem.Vat) / 100

		totalPrice = totalPrice + totalVat
	}

	c.TotalPrice = totalPrice

	// TotalPrice should be rounded to 2 decimal places
	c.TotalPrice = math.Round(float64(c.TotalPrice)*100) / 100
}

// ResetCart resets cart.
func (c *Cart) ResetCart() {
	c.Products = nil
	c.TotalPrice = 0
}
