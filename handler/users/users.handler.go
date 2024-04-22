package users

import (
	userRepository "todolist-app/repository/users"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UsersHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
}

type UsersHandlerImpl struct {
	UsersRepository userRepository.UsersRepository
}

func NewUsersHandler(db *gorm.DB) UsersHandler {
	user := userRepository.NewUsersRepository(db)
	return &UsersHandlerImpl{
		UsersRepository: user,
	}
}

func (handler *UsersHandlerImpl) Login(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from login",
	})
}

func (handler *UsersHandlerImpl) Register(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from register",
	})
}
