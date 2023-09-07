package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/Cart-API/app/models"
	"github.com/tohanilhan/Cart-API/database"
)

// ShowBasket returns the cart of the user
func ShowBasket(c *fiber.Ctx) error {

	req := models.GetCartRequest{}

	// parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// get db connection
	db, err := database.GetDBConnection()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "DB connection error",
			"data":    nil,
		})
	}

	// Get cart of requested user from db
	models.UserCart, err = db.GetCartByUserId(req.UserId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Error getting cart from db",
			"data":    nil,
		})
	}

	// return cart if cart is not empty
	if len(models.UserCart.Products) > 0 {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Cart retrieved successfully",
			"cart":    models.UserCart,
		})
	} else {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cart is empty",
		})
	}

}
