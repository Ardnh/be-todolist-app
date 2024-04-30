package users

import (
	userRepository "todolist-app/repository/users"

	"todolist-app/model"

	"github.com/go-playground/validator/v10"
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
	CreateFollowUserById(c *fiber.Ctx) error
	DeleteFollowUserById(c *fiber.Ctx) error
}

type UsersHandlerImpl struct {
	UsersRepository userRepository.UsersRepository
	Validate        *validator.Validate
}

func NewUsersHandler(db *gorm.DB, validate *validator.Validate) UsersHandler {
	user := userRepository.NewUsersRepository(db)
	return &UsersHandlerImpl{
		UsersRepository: user,
		Validate:        validate,
	}
}

// Login user
// @Summary Login user
// @Description Login user
// @Tags Users
// @Accept json
// @Produce json
// @Param body body model.LoginRequest true "Login"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/login [post]
func (handler *UsersHandlerImpl) Login(c *fiber.Ctx) error {

	// Read body request
	var request model.LoginRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// Validate incoming request
	errValidate := handler.Validate.Struct(request)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": errValidate.Error(),
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from login",
	})
}

// Register user
// @Summary Register user
// @Description Register user
// @Tags Users
// @Accept json
// @Produce json
// @Param body body model.RegisterRequest true "Login"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/register [post]
func (handler *UsersHandlerImpl) Register(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from register",
	})
}

// Find follower by user id
// @Summary Find follower by user id
// @Description Find follower by user id
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/followers/:user_id [get]
func (handler *UsersHandlerImpl) FindFollowersByUserId(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from Find followers by id",
	})
}

// Find following by user id
// @Summary Find following by user id
// @Description Find following by user id
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/following/:user_id [get]
func (handler *UsersHandlerImpl) FindFollowingByUserId(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from find following by id",
	})
}

// Find user profile by id
// @Summary Find user profile by id
// @Description Find user profile by id
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/following [get]
func (handler *UsersHandlerImpl) FindUserProfileById(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from",
	})
}

// Update profile by id
// @Summary Update profile by id
// @Description Update profile by id
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Param body body model.UserUpdateRequest true "Update profile"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/profile/:userId [put]
func (handler *UsersHandlerImpl) UpdateProfileById(c *fiber.Ctx) error {

	return c.Status(200).JSON(fiber.Map{
		"code":    200,
		"message": "Hello from find following by id",
	})
}

// Create follow user by id
// @Summary Create follow user by id
// @Description Create follow user by id
// @Tags Users
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Param body body model.FollowUserCreateRequest true "Follow user by id"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/followers [post]
func (handler *UsersHandlerImpl) CreateFollowUserById(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Hello from find following by id",
	})
}

// Unfollow user
// @Summary Unfollow user
// @Description Unfollow user
// @Tags Users
// @Accept json
// @Produce json
// @Param following_id path string true "following_id"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/followers [delete]
func (handler *UsersHandlerImpl) DeleteFollowUserById(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Hello from find following by id",
	})
}
