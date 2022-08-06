package controllers

import (
	"fmt"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/db"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/sqls"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/structs"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/vars"
)

func AddProductToBasket(c *fiber.Ctx) error {
	// to store one product
	var oneProduct structs.Cart
	var isInCart bool

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

	query := fmt.Sprintf(sqls.GetProduct, db.DbConf.Schema, db.DbConf.TableNameProduct, request.ProductId)
	rows, err := db.Db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// iterate over rows
	for rows.Next() {

		err := rows.Scan(&vars.Product.ProductId, &vars.Product.ProductName, &vars.Product.Price, &vars.Product.Vat, &vars.Product.Quantity)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Error while getting products",
				"error":   err.Error(),
			})
		}

		// if there is not any product in cart, add new product to cart
		if len(vars.Cart) == 0 {
			oneProduct.Quantity = 1
			oneProduct.Price = vars.Product.Price
			oneProduct.ProductId = vars.Product.ProductId
			oneProduct.ProductName = vars.Product.ProductName
			oneProduct.Vat = vars.Product.Vat
			vars.CartResponse.TotalPrice = vars.Product.Price
			// add product to cart
			vars.Cart = append(vars.Cart, oneProduct)

		} else {
			// check if product is already in cart
			for i := range vars.Cart {

				if vars.Cart[i].ProductId == vars.Product.ProductId {

					// add product to cart
					vars.Cart[i].Quantity++
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
				oneProduct.Quantity = 1
				oneProduct.Price = vars.Product.Price
				oneProduct.ProductId = vars.Product.ProductId
				oneProduct.ProductName = vars.Product.ProductName
				oneProduct.Vat = vars.Product.Vat
				vars.Cart = append(vars.Cart, oneProduct)
			}

			// calculate total price
			vars.CartResponse.TotalPrice += vars.Product.Price

			// TotalPrice should be rounded to 2 decimal places
			vars.CartResponse.TotalPrice = math.Round(float64(vars.CartResponse.TotalPrice)*100) / 100

		}
	}

	// if there is no product with the given id in the db then return error
	if vars.Product.ProductId.String() == uuid.Nil.String() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Product not found",
		})
	}

	vars.CartResponse.Cart = vars.Cart
	vars.Product = structs.Product{}

	// return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Product added to your Cart successfully",
	})
}
