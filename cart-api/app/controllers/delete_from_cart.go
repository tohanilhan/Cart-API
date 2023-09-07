package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/Cart-API/app/models"
	"github.com/tohanilhan/Cart-API/database"
)

// RemoveProductFromBasket removes product from cart
func RemoveProductFromBasket(c *fiber.Ctx) error {

	// get request body
	var req models.RemoveFromCartRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

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

	err = models.UserCart.RemoveItem(req.ProductId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// marshal cart to json
	cartJson, err := json.Marshal(models.UserCart)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// update cart in db
	err = db.UpdateCart(models.UserCart, cartJson)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Product removed from your Cart successfully",
	})

}
