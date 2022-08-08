package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/db"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/sqls"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/structs"
)

// ListAllProducts returns all products
func ListAllProducts(c *fiber.Ctx) error {

	var products []structs.Product // object to store products

	// getting query params
	query := fmt.Sprintf(sqls.GetAllProducts, db.DbConf.Schema, db.DbConf.TableNameProduct)

	rows, err := db.Db.Query(query)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})

	}

	// iterate over rows
	for rows.Next() {
		var product structs.Product
		err := rows.Scan(&product.ProductId, &product.ProductName, &product.Price, &product.Vat, &product.Quantity)
		if err != nil {

			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"success": false,
				"message": "Error while getting products",
			})
		}
		// if product quantity is greater than 0, then add product to products array
		if product.Quantity > 0 {
			products = append(products, product)
		} else {
			// if product quantity is 0, then remove product from database
			query := fmt.Sprintf(sqls.DeleteProduct, db.DbConf.Schema, db.DbConf.TableNameProduct, product.ProductId)
			_, err := db.Db.Exec(query)
			if err != nil {
				return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
					"success": false,
					"message": err.Error(),
				})
			}

		}

	}

	// return products
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"message":  "Products retrieved successfully",
		"products": products,
	})

}
