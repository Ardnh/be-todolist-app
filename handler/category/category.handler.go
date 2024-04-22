package category

import (
	categoryRepository "todolist-app/repository/category"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type CategoryHandler interface {
	Create(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	FindById(c *fiber.Ctx) error
}

type CategoryHandlerImpl struct {
	CategoryRepository categoryRepository.CategoryRepository
}

func NewCategoryHandler(db *gorm.DB) CategoryHandler {
	categoryRepository := categoryRepository.NewCategoryRepository(db)
	return &CategoryHandlerImpl{
		CategoryRepository: categoryRepository,
	}
}

func (handler *CategoryHandlerImpl) Create(c *fiber.Ctx) error
func (handler *CategoryHandlerImpl) Update(c *fiber.Ctx) error
func (handler *CategoryHandlerImpl) Delete(c *fiber.Ctx) error
func (handler *CategoryHandlerImpl) FindAll(c *fiber.Ctx) error
func (handler *CategoryHandlerImpl) FindById(c *fiber.Ctx) error
