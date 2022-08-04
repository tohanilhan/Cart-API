package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/db"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/sqls"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/structs"
)

var (
	products []structs.Product // object to store products
)

func ListAllProducts(c *fiber.Ctx) error {

	// getting query params
	query := fmt.Sprintf(sqls.GetAllProductsQuery, db.DbConf.Schema, db.DbConf.TableName)

	rows, err := db.Db.Query(query)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	// iterate over rows
	for rows.Next() {
		var product structs.Product
		err := rows.Scan(&product.ProductId, &product.ProductName, &product.Price, &product.Vat, &product.Total)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}
		products = append(products, product)
	}

	// return products
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"success":  true,
		"message":  "Products retrieved successfully",
		"products": products,
	})
}
