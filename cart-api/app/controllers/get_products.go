package controllers

import (
	"github.com/tohanilhan/Cart-API/app/models"
	"github.com/tohanilhan/Cart-API/database"

	"github.com/gofiber/fiber/v2"
)

// ListAllProducts returns all products
func ListAllProducts(c *fiber.Ctx) error {

	// get db connection
	db, err := database.GetDBConnection()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "db_connection_error",
			"data":    nil,
		})
	}

	models.UserProduct, err = db.GetProducts()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Error getting products from db",
			"data":    nil,
		})
	}

	// return products
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"message":  "Products retrieved successfully",
		"products": models.UserProduct,
	})

}
