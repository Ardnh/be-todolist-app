package todolist

import (
	todolistRepository "todolist-app/repository/todolist"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type TodolistHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindByUserId(c *fiber.Ctx) error
	FindByCategory(c *fiber.Ctx) error
}

type TodolistHandlerImpl struct {
	TodolistRepository todolistRepository.TodolistRepository
}

func NewTodolistHandler(db *gorm.DB) TodolistHandler {
	repository := todolistRepository.NewTodolistRepository(db)
	return &TodolistHandlerImpl{
		TodolistRepository: repository,
	}
}

func (handler *TodolistHandlerImpl) Create(c *fiber.Ctx) error
func (handler *TodolistHandlerImpl) Update(c *fiber.Ctx) error
func (handler *TodolistHandlerImpl) Delete(c *fiber.Ctx) error
func (handler *TodolistHandlerImpl) FindById(c *fiber.Ctx) error
func (handler *TodolistHandlerImpl) FindAll(c *fiber.Ctx) error
func (handler *TodolistHandlerImpl) FindByUserId(c *fiber.Ctx) error
func (handler *TodolistHandlerImpl) FindByCategory(c *fiber.Ctx) error
