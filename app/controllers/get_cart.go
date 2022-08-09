package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/vars"
)

// ShowBasket returns the cart of the user
func ShowBasket(c *fiber.Ctx) error {

	// return cart if cart is not empty
	if len(vars.Cart) > 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Cart retrieved successfully",
			"cart":    vars.CartResponse,
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cart is empty",
		})
	}

}
