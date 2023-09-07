package controllers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tohanilhan/Cart-API/app/models"
	"github.com/tohanilhan/Cart-API/database"
	"github.com/tohanilhan/Cart-API/pkg/utils"
	"github.com/tohanilhan/Cart-API/vars"
)

var (
	OrderTotalAmount float64
	Month            string
)

// CompleteOrder function is used to complete the order.
func CompleteOrder(c *fiber.Ctx) error {

	req := models.CompleteOrderRequest{}

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

	// check if vars.Cart is empty
	if len(models.UserCart.Products) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cart is empty",
		})
	}

	// update products quantity in database

	for _, item := range models.UserCart.Products {

		// decrease quantity for each product in cart by 1
		if item.Quantity == 1 {
			err := db.DecreaseQuantity(item.ProductId, 1)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"success": false,
					"message": err.Error(),
				})
			}
		} else {
			// decrease quantity for each product in cart by quantity of that product
			err := db.DecreaseQuantity(item.ProductId, int(item.Quantity))
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"success": false,
					"message": err.Error(),
				})
			}
		}

	}

	vars.TotalOrder, err = db.GetTotalOrderCount(req.UserId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Create discounts
	discount := []models.Discount{
		models.DiscountOnFourthOrderMoreThanGivenAmount{},
		models.DiscountOnSameThirdProducts{},
		models.DiscountOnOrderMoreThanGivenAmountInAMonth{},
	}

	timestamp, timex := utils.GetMoment()

	// set order struct
	models.UserOrder = models.Order{
		Cart:                      models.UserCart,
		OrderId:                   uuid.New(),
		UserId:                    req.UserId,
		TotalPriceWithoutDiscount: models.UserCart.TotalPrice,
		Timestamp:                 timestamp,
		Timex:                     timex,
	}

	// Apply the best promotion to cart
	models.UserOrder.ApplyBestDiscounts(discount)

	// marshal order struct to json
	orderJson, err := json.Marshal(models.UserOrder.Cart)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// insert order to db
	err = db.SaveOrder(models.UserOrder, orderJson)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// Update this cart in db as completed
	err = db.UpdateCartAsCompleted(models.UserCart.UserId)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Order completed successfully",
		"order":   models.UserOrder,
	})
}
