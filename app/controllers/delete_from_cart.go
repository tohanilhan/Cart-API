package controllers

import (
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/vars"
)

func RemoveProductFromBasket(c *fiber.Ctx) error {

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

	isInCart := DeleteProductFromBasket(request.ProductId)
	// // check if product is in vars.Cart

	if !isInCart {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Product is not in cart",
		})
	}

	// return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Product removed from your Cart successfully",
	})

}

func DeleteProductFromBasket(productId string) bool {

	var isInCart bool

	// check if product is in vars.Cart
	for i, item := range vars.Cart {

		if item.ProductId.String() == productId {
			isInCart = true
			if vars.Cart[i].Quantity > 1 {
				vars.Cart[i].Quantity--
			} else {
				// remove product from vars.CartResponse
				vars.Cart = append(vars.Cart[:i], vars.Cart[i+1:]...)

			}

			// update vars.CartResponse.Cart
			vars.CartResponse.Cart = vars.Cart
			// calculate total price
			vars.CartResponse.TotalPrice -= item.Price
			// TotalPrice should be rounded to 2 decimal places
			vars.CartResponse.TotalPrice = math.Round(float64(vars.CartResponse.TotalPrice)*100) / 100
		}

		if isInCart {
			break
		}

	}
	return isInCart

}
