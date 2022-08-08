package main

import (
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/google/uuid"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/db"
	configs "github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/pkg/config"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/pkg/middleware"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/pkg/routes"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/pkg/utils"
	"github.com/tohanilhan/Patika.dev-Property-Finder-Go-Bootcamp-Final-Project/vars"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

var (
	app    *fiber.App
	config fiber.Config
	err    error
	once   sync.Once
)

func init() {
	// load env file
	godotenv.Load()

	// init user id
	initUserID()

	// init given amount
	initGivenAmount()

	// verify db connection
	initDb()
	initFiber()
}

func main() {

	// setup routes
	setupRoutes(app) // new

	// Listen on server 8000 and catch error if any
	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}

}

// setupRoutes function sets up routes for fiber app
func setupRoutes(app *fiber.App) {

	api := app.Group("/api/v1")

	// give response when at /api/v1
	api.Get("", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"success": true,
			"message": "Server is up and running!",
		})
	})

	// setup routes
	routes.PublicRoute(api.Group("/products-api"))

}

// initFiber function initializes fiber app
func initFiber() {

	// setup config
	config = configs.FiberConfig()

	// Fiber instance
	app = fiber.New(config)

	// Middleware
	middleware.FiberMiddleware(app)
	log.Println("Fiber init OK.")
}

// initDb function initializes database connection
func initDb() {
	db.InitDb()
	log.Println("Db init OK.")
}

// initUserID function initializes user id
func initUserID() {
	// UserId will be created only once when program starts by using singleton pattern
	// If you want to create new user id, you need to restart the program
	once.Do(func() {
		vars.UserId = uuid.New()
	})

}

// initGivenAmount function initializes given amount
func initGivenAmount() {

	vars.GivenAmount, err = strconv.ParseFloat(os.Getenv("GIVEN_AMOUNT"), 64)
	if err != nil {
		log.Fatal(err)
	}

}
