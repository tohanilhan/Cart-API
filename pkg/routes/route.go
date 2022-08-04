package routes

import (
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/app/controllers"

	"github.com/gofiber/fiber/v2"
)

func PublicRoute(route fiber.Router) {

	route.Get("/products/", controllers.ListAllProducts) //Get all products

	//// add product to basket
	//route.Post("/products/basket", controllers.AddProductToBasket) //Add product to basket

	// // show cart of products in basket
	// route.Get("/products/basket", controllers.ShowBasket) //Show basket

	// // remove product from basket
	// route.Delete("/products/basket/:id", controllers.RemoveProductFromBasket) //Remove product from basket

	// // complete order
	// route.Post("/products/order", controllers.CompleteOrder) //Complete order

}
