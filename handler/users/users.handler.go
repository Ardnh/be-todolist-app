package users

import (
	userRepository "todolist-app/repository/users"

	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

type UsersHandler interface {
	Login(c fiber.Ctx) error
	Register(c fiber.Ctx) error
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

func (handler *UsersHandlerImpl) Login(c fiber.Ctx) error
func (handler *UsersHandlerImpl) Register(c fiber.Ctx) error
