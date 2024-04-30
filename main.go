package main

import (
	"os"
	"todolist-app/app"
	"todolist-app/helper"
	"todolist-app/routes"

	_ "todolist-app/docs"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

// @title           Todolist API
// @version         1.0
// @description     API Documentation for Todolist API.

// @contact.name   Muhammad Ardan Hilal
// @contact.url    ardn.h79@gmail.com
// @contact.email  ardn.h79@gmail.com

// @host      localhost:8080
// @BasePath  /api/v1

func main() {

	newApp := fiber.New()
	db := app.DbConnection()
	validate := validator.New(validator.WithRequiredStructEnabled())

	routes.SetupRoutes(newApp, db, validate)

	port := os.Getenv("APP_PORT")
	err := newApp.Listen(port)
	helper.PanicIfError(err)
}
