package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tohanilhan/Cart-API/app/models"
	"github.com/tohanilhan/Cart-API/database"
)

func AddProductToBasket(c *fiber.Ctx) error {

	var req models.AddToChartRequest

	// parse request body
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Get db connection
	db, err := database.GetDBConnection()
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "DB connection error",
			"data":    nil,
		})
	}

	// get product details from db
	productFromDb, err := db.GetProductById(req.ProductId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": "Error getting product from db",
			"data":    nil,
		})
	}

	// if there is no product with the given id in the db then return error
	if productFromDb.ProductId == uuid.Nil.String() {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Product not found",
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

	if len(models.UserCart.Products) == 0 {
		// if cart is empty then create a new cart
		models.UserCart = models.Cart{
			CartId:     uuid.New().String(),
			UserId:     req.UserId,
			Products:   []models.Product{},
			TotalPrice: 0,
			Status:     "NOT_COMPLETED",
		}

		err = db.CreateCart(models.UserCart, nil)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}
	}

	err = models.UserCart.AddProduct(productFromDb)
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

	models.UserCart.UserId = req.UserId

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
		"message": "Product added to your Cart successfully",
	})
}
