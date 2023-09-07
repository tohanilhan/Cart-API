package routes

import (
	"github.com/tohanilhan/Cart-API/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoutes(a *fiber.App) {

	// Create routes group.
	routeV1 := a.Group("/api/v1")

	routeV1.Get("/ping", controllers.Ping)

	//Get all products
	routeV1.Get("/products/", controllers.ListAllProducts)

	//Add product to cart
	routeV1.Post("/products/", controllers.AddProductToBasket)

	//Show cart
	routeV1.Post("/cart/", controllers.ShowBasket)

	//Remove product from cart
	routeV1.Delete("/cart/", controllers.RemoveProductFromBasket)

	//Complete order
	routeV1.Post("/order", controllers.CompleteOrder)

}
