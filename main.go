package main

import (
	"os"
	"todolist-app/app"
	"todolist-app/helper"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

func main() {

	newApp := fiber.New()
	db := app.DbConnection()

	app.SetupRoutes(newApp, db)

	port := os.Getenv("APP_PORT")
	err := newApp.Listen(port)
	helper.PanicIfError(err)
}
