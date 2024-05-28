package users

import (
	"fmt"
	"strconv"
	"todolist-app/helper"
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
// @Router /user/login [post]
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

	// Get user by email
	userResult, errUserFindByEmail := handler.UsersRepository.FindByEmail(c, request.Email)
	if errUserFindByEmail != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": errUserFindByEmail.Error(),
		})
	}

	// compare password from body request and from database
	errComparePassword := helper.CheckPasswordHash(request.Password, userResult.Password)
	if !errComparePassword {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Wrong password!",
		})
	}

	// generate jwt token
	token, errGenerateToken := helper.GenerateToken(userResult.ID)
	if errGenerateToken != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": errGenerateToken.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":     fiber.StatusOK,
		"message":  "Login successfully",
		"token":    token,
		"username": userResult.Username,
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
// @Router /user/register [post]
func (handler *UsersHandlerImpl) Register(c *fiber.Ctx) error {

	// 1. Parser body request
	var request model.RegisterRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// 2. Validasi json yang dikirim
	errValidate := handler.Validate.Struct(request)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": errValidate.Error(),
		})
	}

	// 3. Cek apakah user dengan email yang dikirim sudah ada di database
	result, err := handler.UsersRepository.FindByEmail(c, request.Email)
	if result.Email == request.Email {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "User already exist",
		})
	}

	// 4. Hash user password
	hashResult, errHash := helper.HashPassword(request.Password)
	if errHash != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": "Error hashing password",
		})
	}
	request.Password = hashResult

	// 5. Save user to database
	req := model.User{
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
	}

	errRegister := handler.UsersRepository.Register(c, &req)
	if errRegister != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Successfully register user",
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
// @Router /user/followers [get]
func (handler *UsersHandlerImpl) FindFollowersByUserId(c *fiber.Ctx) error {

	// desc: find account that follow your account by your user id
	// 1. get userId from jwt token
	var userId uint = helper.UserId
	fmt.Println(userId)
	page := c.Query("page", "1")
	pageInt, _ := strconv.Atoi(page)
	pageSize := c.Query("pageSize", "100")
	pageSizeInt, _ := strconv.Atoi(pageSize)
	username := c.Query("username", "")

	// 2. get follower from repository
	result, totalEntries, err := handler.UsersRepository.FindFollowersByUserId(c, userId, pageInt, pageSizeInt, username)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":       fiber.StatusOK,
		"message":    "Success get followers",
		"page":       page,
		"pageSize":   pageSize,
		"totalPages": totalEntries,
		"data":       result,
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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
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
// @Param body body model.ProfileUpdateRequest true "Update profile"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /api/v1/user/profile/:userId [put]
func (handler *UsersHandlerImpl) UpdateProfileById(c *fiber.Ctx) error {

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Hello from find following by id",
	})
}

// Create follow user by id
// @Summary Create follow user by id
// @Description Create follow user by id
// @Tags Users
// @Accept json
// @Produce json
// @Param body body model.FollowUserCreateRequest true "Follow user by id"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/following [post]
func (handler *UsersHandlerImpl) CreateFollowUserById(c *fiber.Ctx) error {

	// 1. Parser body request
	var request model.FollowUserCreateRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	// 2. Validasi json yang dikirim
	errValidate := handler.Validate.Struct(request)
	if errValidate != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": errValidate.Error(),
		})
	}

	// 2. Create follow user
	errCreate := handler.UsersRepository.CreateFollowUserById(c, request.UserId, request.TargetId)
	if errCreate != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": errCreate.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "successfully follow user",
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
