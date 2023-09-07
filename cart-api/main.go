package main

import (
	"fmt"

	"github.com/caarlos0/env/v6"
	"github.com/tohanilhan/Cart-API/database"

	"github.com/tohanilhan/Cart-API/pkg/config"
	"github.com/tohanilhan/Cart-API/pkg/middleware"
	"github.com/tohanilhan/Cart-API/pkg/routes"
	"github.com/tohanilhan/Cart-API/pkg/utils"
	"github.com/tohanilhan/Cart-API/vars"

	"github.com/gofiber/fiber/v2"
)

func init() {
	// parse env file
	parseEnv()

	// verify db connection
	initDb()

}

func main() {

	config := config.FiberConfig()
	config.BodyLimit = 32 * 1024 * 1024 // 512 MB file upload permitted

	// config.StreamRequestBody !!!!! bu konuyu araştır.
	// Define new Fiber app with config here:
	app := fiber.New(config)

	// Middlewares here:
	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	// Routes here:
	// routes.SwaggerRoute(app)  // Register a route for API Docs (Swagger).
	routes.PublicRoutes(app)  // Register a public routes for app.
	routes.NotFoundRoute(app) // Register route for 404 Error.

	// Start fiber server (with graceful shutdown).
	utils.StartServerWithGracefulShutdown(app)

}

// initDb function initializes database connection
func initDb() {
	err := database.PostgreSQLConnection()
	if err != nil {
		// TODO:  call log service
		panic("err while db connection: " + err.Error())
	}
}

func parseEnv() {
	err := env.Parse(&vars.AppConfigs)
	if err != nil {
		fmt.Printf("%+v\n", err)
	}
}
