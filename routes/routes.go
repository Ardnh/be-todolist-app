package routes

import (
	"fmt"
	"todolist-app/handler/category"
	"todolist-app/handler/todolist"
	"todolist-app/handler/users"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Request: %s %s \n", c.Method(), c.OriginalURL())
		return c.Next()
	})

	userHandler := users.NewUsersHandler(db)
	categoryHandler := category.NewCategoryHandler(db)
	todolistHandler := todolist.NewTodolistHandler(db)

	appGroup := app.Group("api/v1")

	// Users
	usersGroup := appGroup.Group("user")
	usersGroup.Post("/login", userHandler.Login)
	usersGroup.Post("/register", userHandler.Register)
	usersGroup.Get("/followers", userHandler.FindFollowersByUserId)
	usersGroup.Get("/following", userHandler.FindFollowingByUserId)
	usersGroup.Put("/profile/:userId", userHandler.UpdateProfileById)

	// Todolist
	todolistGroup := appGroup.Group("todolist")
	todolistGroup.Post("/", todolistHandler.Create)
	todolistGroup.Put("/:id", todolistHandler.Update)
	todolistGroup.Put("/:todolistId", todolistHandler.UpdateTodolistItem)
	todolistGroup.Delete("/:todolistId", todolistHandler.Delete)
	todolistGroup.Get("/:todolistId", todolistHandler.FindTodolistById)
	todolistGroup.Get("/", todolistHandler.FindAll)
	todolistGroup.Get("/:userId", todolistHandler.FindTodolistByUserId)
	todolistGroup.Get("/categoryId", todolistHandler.FindTodolistByCategoryId)

	// Category
	categoryGroup := appGroup.Group("category")
	categoryGroup.Post("/", categoryHandler.Create)
	categoryGroup.Put("/", categoryHandler.Update)

	// Super admin

}
