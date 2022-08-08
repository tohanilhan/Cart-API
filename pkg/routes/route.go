package routes

import (
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoute(route fiber.Router) {

	//Get all products
	route.Get("/products/", controllers.ListAllProducts)

	//Add product to basket
	route.Post("/products/", controllers.AddProductToBasket)

	//Show basket
	route.Get("/products/basket/", controllers.ShowBasket)

	//Remove product from basket
	route.Delete("/products/basket/", controllers.RemoveProductFromBasket)

	//Complete order
	route.Get("/products/order", controllers.CompleteOrder)

}
