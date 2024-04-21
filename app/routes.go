package app

import (
	"fmt"
	"todolist-app/handler/users"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	app.Use(func(c fiber.Ctx) error {
		fmt.Printf("Request: %s %s \n", c.Method(), c.OriginalURL())
		return c.Next()
	})

	userHandler := users.NewUsersHandler(db)

	usersGroup := app.Group("user")
	usersGroup.Post("/login", userHandler.Login)
	usersGroup.Post("/register", userHandler.Register)

}
