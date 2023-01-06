package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinlee0/fiber-gorm-sample/database"
	"github.com/jinlee0/fiber-gorm-sample/userRouter"
	"log"
)

func setupRoutes(app *fiber.App) {
	api := app.Group("/api")
	userRouter.Route(api, "/users")
}

func main() {
	database.ConnectDb()
	app := fiber.New()
	setupRoutes(app)
	log.Fatal(app.Listen(":3000"))
}
