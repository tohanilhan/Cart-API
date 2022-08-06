package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/db"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/sqls"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/structs"
)

var (
	cart         []structs.Cart // object to store cart
	cartResponse structs.CartResponse
	product      structs.Product
)

func ListAllProducts(c *fiber.Ctx) error {

	var products []structs.Product // object to store products

	// getting query params
	query := fmt.Sprintf(sqls.GetAllProducts, db.DbConf.Schema, db.DbConf.TableName)

	rows, err := db.Db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})

	}

	// iterate over rows
	for rows.Next() {
		var product structs.Product
		err := rows.Scan(&product.ProductId, &product.ProductName, &product.Price, &product.Vat, &product.Quantity)
		if err != nil {

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Error while getting products",
			})
		}
		products = append(products, product)
	}

	// return products
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"message":  "Products retrieved successfully",
		"products": products,
	})

}

func AddProductToBasket(c *fiber.Ctx) error {

	type Request struct {
		ProductId string `json:"product_id"`
	}

	// get request body
	var request Request
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	query := fmt.Sprintf(sqls.GetProduct, db.DbConf.Schema, db.DbConf.TableName, request.ProductId)

	rows, err := db.Db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}
	// iterate over rows
	for rows.Next() {

		var oneProduct structs.Cart

		err := rows.Scan(&product.ProductId, &product.ProductName, &product.Price, &product.Vat, &product.Quantity)
		if err != nil {

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Error while getting products",
				"error":   err.Error(),
			})
		}

		if len(cart) == 0 {
			oneProduct.Quantity = 1
			oneProduct.Price = product.Price
			oneProduct.ProductId = product.ProductId
			oneProduct.ProductName = product.ProductName
			oneProduct.Vat = product.Vat
			cartResponse.TotalPrice = product.Price
			cart = append(cart, oneProduct)
		} else {
			// check if product is already in cart
			for i, item := range cart {

				if item.ProductId == product.ProductId {
					// add product to cart
					cart[i].Quantity++

				} else {
					// if product is not in cart, set quantity to 1
					oneProduct.Quantity = 1
					oneProduct.Price = product.Price
					oneProduct.ProductId = product.ProductId
					oneProduct.ProductName = product.ProductName
					oneProduct.Vat = product.Vat
					// add product to cart
					cart = append(cart, oneProduct)
				}
				// calculate total price
				cartResponse.TotalPrice += product.Price
			}

		}

		cartResponse.Cart = cart

	}

	// return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Product added to your cart successfully",
	})
}

func ShowBasket(c *fiber.Ctx) error {

	// return cart
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Cart retrieved successfully",
		"cart":    cartResponse,
	})
}
