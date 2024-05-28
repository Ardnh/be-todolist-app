package routes

import (
	"fmt"
	"todolist-app/handler/category"
	"todolist-app/handler/todolist"
	"todolist-app/handler/users"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, validate *validator.Validate) {

	app.Use(func(c *fiber.Ctx) error {
		fmt.Printf("Request: %s %s \n", c.Method(), c.OriginalURL())
		return c.Next()
	})

	userHandler := users.NewUsersHandler(db, validate)
	categoryHandler := category.NewCategoryHandler(db, validate)
	todolistHandler := todolist.NewTodolistHandler(db, validate)

	appGroup := app.Group("/api/v1")

	// Swagger
	appGroup.Get("/swagger/*", swagger.HandlerDefault)

	// Users
	usersGroup := appGroup.Group("user")
	usersGroup.Post("/login", userHandler.Login)
	usersGroup.Post("/register", userHandler.Register)
	usersGroup.Get("/followers", userHandler.FindFollowersByUserId)
	usersGroup.Get("/following", userHandler.FindFollowingByUserId)
	usersGroup.Put("/profile/:userId", userHandler.UpdateProfileById)
	usersGroup.Post("/following", userHandler.CreateFollowUserById)
	usersGroup.Delete("/following", userHandler.DeleteFollowUserById)

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
	todolistGroup.Post("/follow/:todolistId", todolistHandler.FollowTodolistByTodolistId)
	todolistGroup.Delete("/follow/:todolistId", todolistHandler.UnfollowTodolistByTodolistId)

	// Category
	categoryGroup := appGroup.Group("category")
	categoryGroup.Post("/", categoryHandler.Create)
	categoryGroup.Put("/", categoryHandler.Update)
	categoryGroup.Delete("/:id", categoryHandler.Delete)
	categoryGroup.Get("/", categoryHandler.FindAll)

}
