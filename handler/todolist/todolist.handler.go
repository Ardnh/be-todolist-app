package todolist

import (
	todolistRepository "todolist-app/repository/todolist"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TodolistHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	UpdateTodolistItem(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindTodolistById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindTodolistByUserId(c *fiber.Ctx) error
	FindTodolistByCategoryId(c *fiber.Ctx) error
}

type TodolistHandlerImpl struct {
	TodolistRepository todolistRepository.TodolistRepository
	Validate           *validator.Validate
}

func NewTodolistHandler(db *gorm.DB, validator *validator.Validate) TodolistHandler {
	repository := todolistRepository.NewTodolistRepository(db)
	return &TodolistHandlerImpl{
		TodolistRepository: repository,
		Validate:           validator,
	}
}

func (handler *TodolistHandlerImpl) Create(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from create todolist",
	})
}

func (handler *TodolistHandlerImpl) Update(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from update todolist",
	})
}

func (handler *TodolistHandlerImpl) UpdateTodolistItem(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from update todolist item",
	})
}

func (handler *TodolistHandlerImpl) Delete(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from delete todolist",
	})
}

func (handler *TodolistHandlerImpl) FindTodolistById(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from find todolist by id",
	})
}

func (handler *TodolistHandlerImpl) FindAll(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from find all",
	})
}

func (handler *TodolistHandlerImpl) FindTodolistByUserId(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from find todolist by user id",
	})
}

func (handler *TodolistHandlerImpl) FindTodolistByCategoryId(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from todolist by category id",
	})
}
