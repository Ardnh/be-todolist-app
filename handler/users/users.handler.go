package users

import (
	userRepository "todolist-app/repository/users"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UsersHandler interface {
	Login(c *fiber.Ctx) error
	Register(c *fiber.Ctx) error
	FindFollowersByUserId(c *fiber.Ctx) error
	FindFollowingByUserId(c *fiber.Ctx) error
	FindUserProfileById(c *fiber.Ctx) error
	UpdateProfileById(c *fiber.Ctx) error
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

func (handler *UsersHandlerImpl) FindFollowersByUserId(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from Find followers by id",
	})
}

func (handler *UsersHandlerImpl) FindFollowingByUserId(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from find following by id",
	})
}

func (handler *UsersHandlerImpl) FindUserProfileById(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from",
	})
}

func (handler *UsersHandlerImpl) UpdateProfileById(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from find following by id",
	})
}
