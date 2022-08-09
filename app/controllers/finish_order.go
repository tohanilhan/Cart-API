package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/db"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/pkg/utils"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/sqls"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/structs"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/vars"
)

var (
	OrderTotalAmount float64
	Month            string
)

// CompleteOrder function is used to complete the order.
func CompleteOrder(c *fiber.Ctx) error {

	// check if vars.Cart is empty
	if len(vars.CartResponse.Cart) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cart is empty",
		})
	}

	// check the stock of the products
	for _, item := range vars.CartResponse.Cart {
		if item.Quantity > vars.Product.Quantity {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Not enough stock",
			})
		}
	}

	timestamp, timex, Month := utils.GetTimestamp()

	// get last order from db
	query := fmt.Sprintf(sqls.GetOrder, db.DbConf.Schema, db.DbConf.TableNameOrder, vars.UserId, "'%"+Month+"%'")
	rows, err := db.Db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error while getting orders",
			"error":   err.Error(),
			"query":   query,
		})
	}

	// iterate over rows
	for rows.Next() {
		err := rows.Scan(&OrderTotalAmount)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Error while getting orders",
				"error":   err.Error(),
			})
		}
	}

	// calculate discount
	totalPriceWithoutDiscount, reason, totalPriceWitDiscount := CalculateDiscount()

	// totalPriceWitDiscount should be rounded to 2 decimal places
	totalPriceWitDiscount = math.Round(float64(totalPriceWitDiscount)*100) / 100

	// discountPrice should be rounded to 2 decimal places
	vars.FinalDiscount = math.Round(float64(vars.FinalDiscount)*100) / 100

	// totalPriceWithoutDiscount should be rounded to 2 decimal places
	totalPriceWithoutDiscount = math.Round(float64(totalPriceWithoutDiscount)*100) / 100

	// set order struct
	order := structs.Order{
		TotalPriceWithDiscount:    totalPriceWitDiscount,
		TotalPriceWithoutDiscount: totalPriceWithoutDiscount,
		Cart:                      vars.CartResponse.Cart,
		OrderId:                   uuid.New(),
		DiscountReason:            reason,
		Discount:                  vars.FinalDiscount,
		UserId:                    vars.UserId,
		Timestamp:                 timestamp,
		Timex:                     timex,
	}

	// marshal order struct to json
	orderJson, err := json.Marshal(order.Cart)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// insert order to database
	// (order_id, user_id, cart, discount, total_price_with_discount, total_price_without_discount, time,timestamp) values ('%s','%s','%s',%f,%f,%f,%d,'%s')
	query = fmt.Sprintf(sqls.InsertOrder, db.DbConf.Schema, db.DbConf.TableNameOrder, order.OrderId, order.UserId, string(orderJson), order.Discount, order.TotalPriceWithDiscount, order.TotalPriceWithoutDiscount, order.Timex, order.Timestamp, reason)
	_, err = db.Db.Exec(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// update products quantity in database

	for _, item := range vars.Cart {
		var query string
		// decrease quantity for each product in cart by 1
		if item.Quantity == 1 {
			query = fmt.Sprintf(sqls.UpdateProductQuantity, db.DbConf.Schema, db.DbConf.TableNameProduct, vars.Product.Quantity-1, item.ProductId)
		} else {
			// decrease quantity for each product in cart by quantity of that product
			query = fmt.Sprintf(sqls.UpdateProductQuantity, db.DbConf.Schema, db.DbConf.TableNameProduct, vars.Product.Quantity-item.Quantity, item.ProductId)
		}
		_, err := db.Db.Exec(query)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": err.Error(),
			})
		}

	}

	// reset cart
	vars.Cart = nil
	vars.CartResponse = structs.CartResponse{}
	// reset discount
	vars.DiscountOnFourthOrderMoreThanGivenAmount = 0
	vars.DiscountOnSameThirdProducts = 0
	vars.DiscountOnOrderMoreThanGivenAmountInAMonth = 0
	vars.FinalDiscount = 0

	// return success message
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success": true,
		"message": "Order completed successfully",
		"order":   order,
	})
}

// CalculateDiscount function is used to calculate discount.
func CalculateDiscount() (float64, string, float64) {

	var totalPriceWithoutDiscount float64 // total price for the order
	var reason string                     // reason for discount

	// If there are more than 3 items of the same product, then fourth and subsequent ones would have %8 off.
	for _, item := range vars.Cart {
		if item.Quantity > 3 {
			vars.DiscountOnSameThirdProducts += item.Price * float64(item.Quantity-3) * 0.08

		} else {
			vars.DiscountOnSameThirdProducts = 0
		}
		totalPriceWithoutDiscount += item.Price * float64(item.Quantity)
	}

	// check if this is the fourth order. Every fourth order whose total is more than given amount may have discount depending on products.
	if totalPriceWithoutDiscount >= vars.GivenAmount && vars.TotalOrder == 3 {

		// Products whose VAT is %1 don’t have any discount but products whose VAT is %8 and %18 have discount of %10 and %15 respectively.
		for _, item := range vars.Cart {
			if item.Vat == 1 {
				vars.DiscountOnFourthOrderMoreThanGivenAmount += 0

			} else if item.Vat == 8 {

				vars.DiscountOnFourthOrderMoreThanGivenAmount += item.Price * float64(item.Quantity) * 0.1

			} else if item.Vat == 18 {

				vars.DiscountOnFourthOrderMoreThanGivenAmount += item.Price * float64(item.Quantity) * 0.15

			}

		}
		// reset order count
		vars.TotalOrder = 0

	} else {
		// increment order count
		vars.TotalOrder++
	}

	// If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off.
	if OrderTotalAmount >= vars.GivenAmount {

		log.Println(vars.GivenAmount)
		vars.DiscountOnOrderMoreThanGivenAmountInAMonth = totalPriceWithoutDiscount * 0.1

	}

	// Check if there is any discount
	if vars.DiscountOnFourthOrderMoreThanGivenAmount > 0 || vars.DiscountOnSameThirdProducts > 0 || vars.DiscountOnOrderMoreThanGivenAmountInAMonth > 0 {
		// check which discount is bigger and apply it
		if vars.DiscountOnFourthOrderMoreThanGivenAmount > vars.DiscountOnSameThirdProducts {
			if vars.DiscountOnFourthOrderMoreThanGivenAmount > vars.DiscountOnOrderMoreThanGivenAmountInAMonth {
				vars.FinalDiscount = vars.DiscountOnFourthOrderMoreThanGivenAmount
				reason = "Every fourth order whose total is more than given amount may have discount depending on products. Products whose VAT is %1 don’t have any discount but products whose VAT is %8 and %18 have discount of %10 and %15 respectively."
			} else {
				vars.FinalDiscount = vars.DiscountOnOrderMoreThanGivenAmountInAMonth
				reason = "If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off."
			}
		} else {
			if vars.DiscountOnSameThirdProducts > vars.DiscountOnOrderMoreThanGivenAmountInAMonth {
				vars.FinalDiscount = vars.DiscountOnSameThirdProducts
				reason = "If there are more than 3 items of the same product, then fourth and subsequent ones would have %8 off."
			} else {
				vars.FinalDiscount = vars.DiscountOnOrderMoreThanGivenAmountInAMonth
				reason = "If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off."
			}
		}
	} else {
		vars.FinalDiscount = 0
		reason = "No discount"
	}

	totalPriceWithDiscount := totalPriceWithoutDiscount - vars.FinalDiscount

	// empty vars
	vars.DiscountOnFourthOrderMoreThanGivenAmount = 0
	vars.DiscountOnSameThirdProducts = 0
	vars.DiscountOnOrderMoreThanGivenAmountInAMonth = 0

	return totalPriceWithoutDiscount, reason, totalPriceWithDiscount
}
