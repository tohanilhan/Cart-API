package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/db"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/sqls"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/structs"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/util"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/vars"
)

var (
	orderTotalAmount float64
	orderTimestamp   string
)

func CompleteOrder(c *fiber.Ctx) error {

	// check if vars.Cart is empty
	if len(vars.CartResponse.Cart) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Cart is empty",
		})
	}

	timestamp, timex, month := util.GetTimestamp()

	// get last order from db
	query := fmt.Sprintf(sqls.GetOrder, db.DbConf.Schema, db.DbConf.TableNameOrder, vars.UserId, vars.GivenAmount, "'%"+month+"%'")
	rows, err := db.Db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "Error while getting orders",
			"error":   err.Error(),
		})
	}

	// iterate over rows
	for rows.Next() {
		err := rows.Scan(&orderTotalAmount, &orderTimestamp)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Error while getting orders",
				"error":   err.Error(),
			})
		}
	}

	var totalPrice float64 // total price for the order

	// If there are more than 3 items of the same product, then fourth and subsequent ones would have %8 off.
	for _, item := range vars.Cart {
		if item.Quantity > 3 {
			vars.DiscountOnSameThirdProducts += item.Price * float64(item.Quantity-3) * 0.08

		}
		totalPrice += item.Price * float64(item.Quantity)
	}
	// totalPrice -= discountPrice

	// check if this is the fourth order. Every fourth order whose total is more than given amount may have discount depending on products.

	if vars.TotalOrder == 3 && totalPrice > vars.GivenAmount {

		// Products whose VAT is %1 don’t have any discount but products whose VAT is %8 and %18 have discount of %10 and %15 respectively.
		for _, item := range vars.Cart {
			if item.Vat == 1 {
				vars.DiscountOnFourthOrderMoreThanGivenAmount += 0
				totalPrice += item.Price * float64(item.Quantity)
			} else if item.Vat == 8 {
				vars.DiscountOnFourthOrderMoreThanGivenAmount += item.Price * float64(item.Quantity) * 0.1
				//totalPrice -= discountPrice
			} else if item.Vat == 18 {
				vars.DiscountOnFourthOrderMoreThanGivenAmount += item.Price * float64(item.Quantity) * 0.15
				//totalPrice -= discountPrice
			}
		}
		// reset order count
		vars.TotalOrder = 0

	} else {
		// increment order count
		vars.TotalOrder++
	}

	// If the customer made purchase which is more than given amount in a month then all subsequent purchases should have %10 off.
	if totalPrice > vars.DiscountOnOrderMoreThanGivenAmountInAMonth && strings.Contains(orderTimestamp, month) {
		vars.DiscountOnOrderMoreThanGivenAmountInAMonth += totalPrice * 0.1
	}

	// check which discount is bigger and apply it
	if vars.DiscountOnFourthOrderMoreThanGivenAmount > vars.DiscountOnSameThirdProducts {
		if vars.DiscountOnFourthOrderMoreThanGivenAmount > vars.DiscountOnOrderMoreThanGivenAmountInAMonth {
			vars.FinalDiscount = vars.DiscountOnFourthOrderMoreThanGivenAmount
		} else {
			vars.FinalDiscount = vars.DiscountOnOrderMoreThanGivenAmountInAMonth
		}
	} else {
		if vars.DiscountOnSameThirdProducts > vars.DiscountOnOrderMoreThanGivenAmountInAMonth {
			vars.FinalDiscount = vars.DiscountOnSameThirdProducts
		} else {
			vars.FinalDiscount = vars.DiscountOnOrderMoreThanGivenAmountInAMonth
		}
	}

	// apply discount to total price
	totalPrice -= vars.FinalDiscount

	// discountPrice should be rounded to 2 decimal places
	vars.FinalDiscount = math.Round(float64(vars.FinalDiscount)*100) / 100

	// totalPrice should be rounded to 2 decimal places
	totalPrice = math.Round(float64(totalPrice)*100) / 100

	// set order struct
	order := structs.Order{
		TotalPriceWithDiscount:    totalPrice - vars.FinalDiscount,
		TotalPriceWithoutDiscount: totalPrice,
		Cart:                      vars.CartResponse.Cart,
		OrderId:                   uuid.New(),
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
	query = fmt.Sprintf(sqls.InsertOrder, db.DbConf.Schema, db.DbConf.TableNameOrder, order.OrderId, order.UserId, string(orderJson), order.Discount, order.TotalPriceWithDiscount, order.TotalPriceWithoutDiscount, order.Timex, order.Timestamp)
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

	log.Println("Same third products discount: ", vars.DiscountOnSameThirdProducts)
	log.Println("Discount on more than given amount in a month: ", vars.DiscountOnOrderMoreThanGivenAmountInAMonth)
	log.Println("Discount on fourth order: ", vars.DiscountOnFourthOrderMoreThanGivenAmount)
	log.Println("Final Discount: ", vars.FinalDiscount)

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
