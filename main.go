package main

import (
	"os"
	"todolist-app/app"
	"todolist-app/helper"
	"todolist-app/routes"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	newApp := fiber.New()
	db := app.DbConnection()

	routes.SetupRoutes(newApp, db)

	port := os.Getenv("APP_PORT")
	err := newApp.Listen(port)
	helper.PanicIfError(err)
}
