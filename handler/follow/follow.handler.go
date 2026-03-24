package follow

import (
	"fmt"
	"strconv"
	followRepository "todolist-app/repository/follow"

	"todolist-app/model"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type FollowHandler interface {
	FindFollowersByUserId(c *fiber.Ctx) error
	FindFollowingByUserId(c *fiber.Ctx) error
	CreateFollowingUserById(c *fiber.Ctx) error
	UnfollowUserById(c *fiber.Ctx) error
}

type FollowHandlerImpl struct {
	FollowRepository followRepository.FollowRepository
	Validate         *validator.Validate
}

func NewFollowHandler(db *gorm.DB, validate *validator.Validate) FollowHandler {
	user := followRepository.NewFollowRepository(db)
	return &FollowHandlerImpl{
		FollowRepository: user,
		Validate:         validate,
	}
}

// Find follower by user id
// @Summary Find follower by user id
// @Description Find follower by user id
// @Tags Following / Followers
// @Accept json
// @Produce json
// @Param user_id path string true "user_id"
// @Param page path string false "page"
// @Param pageSize path string false "pageSize"
// @Param username path string false "username"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security Bearer
// @Router /user/followers/{user_id} [get]
func (handler *FollowHandlerImpl) FindFollowersByUserId(c *fiber.Ctx) error {

	// desc: find account that follow your account by your user id
	// 1. get userId from jwt token
	userId := c.Params("user_id", "")

	fmt.Println(userId)

	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid user id",
		})
	}

	userIdVal, err := strconv.ParseUint(userId, 10, 32) // basis 10, 32-bit
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid user id",
		})
	}

	page := c.Query("page", "1")
	pageInt, _ := strconv.Atoi(page)
	pageSize := c.Query("pageSize", "100")
	pageSizeInt, _ := strconv.Atoi(pageSize)
	username := c.Query("username", "")

	// 2. get follower from repository
	result, totalEntries, err := handler.FollowRepository.FindFollowersByUserId(c, uint(userIdVal), pageInt, pageSizeInt, username)

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
// @Tags Following / Followers
// @Accept json
// @Produce json
// @Param page path string false "page"
// @Param pageSize path string false "pageSize"
// @Param username path string false "username"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Security Bearer
// @Router /user/following/:userId [get]
func (handler *FollowHandlerImpl) FindFollowingByUserId(c *fiber.Ctx) error {

	userId := c.Params("userId", "")

	if userId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid user id",
		})
	}

	userIdVal, err := strconv.ParseUint(userId, 10, 32) // basis 10, 32-bit
	fmt.Println(userIdVal)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code":    fiber.StatusBadRequest,
			"message": "Invalid user id",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Hello from find following by id",
	})
}

// Create follow user by id
// @Summary Create follow user by id
// @Description Create follow user by id
// @Tags Following / Followers
// @Accept json
// @Produce json
// @Param body body model.FollowUserCreateRequest true "Follow user by id"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/following [post]
func (handler *FollowHandlerImpl) CreateFollowingUserById(c *fiber.Ctx) error {

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
	errCreate := handler.FollowRepository.CreateFollowUserById(c, request.UserId, request.Following)
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
// @Tags Following / Followers
// @Accept json
// @Produce json
// @Param body body model.UnfollowUserRequest true "Unfollow user"
// @Success 200 {object} map[string]interface{} "Success update category"
// @Failure 400 {object} map[string]interface{} "Invalid request body or missing required fields"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /user/unfollow [delete]
func (handler *FollowHandlerImpl) UnfollowUserById(c *fiber.Ctx) error {

	// 1. Parser body request
	var request model.UnfollowUserRequest
	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"code":    fiber.StatusInternalServerError,
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"code":    fiber.StatusOK,
		"message": "Hello from find following by id",
	})
}
