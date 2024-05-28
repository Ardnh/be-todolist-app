package users

import (
	"errors"
	"strings"
	"todolist-app/helper"
	"todolist-app/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UsersRepository interface {
	CreateFollowUserById(ctx *fiber.Ctx, userId int, userTargetId int) error
	Register(ctx *fiber.Ctx, req *model.User) error
	FindByUsernameOrEmail(ctx *fiber.Ctx, req string, isEmail bool) (*model.User, error)
	FindByEmail(ctx *fiber.Ctx, email string) (*model.User, error)
	FindFollowersByUserId(ctx *fiber.Ctx, userId uint, pageInt int, pageSizeInt int, username string) ([]model.UserWithProfile, int64, error)
	// FindFollowingByUserId(ctx *fiber.Ctx, userId int) ([]*model.UserWithProfile, error)
	// FindUserProfileById(ctx *fiber.Ctx, userId int) (*model.Profile, error)
	// UpdateProfileById(ctx *fiber.Ctx, userId int, req model.ProfileUpdateRequest) error
	// CreateFollowUserById(ctx *fiber.Ctx, userId int, userTargetId int) error
	// DeleteFollowUserById(ctx *fiber.Ctx, followId int) error
}

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {

	return &UsersRepositoryImpl{
		Db: db,
	}
}

var tableUser = "users"
var tableProfile = "user_profiles"
var tableFollowers = "follow_users"

func (repository *UsersRepositoryImpl) Register(ctx *fiber.Ctx, req *model.User) error {

	tx := repository.Db.Begin()
	defer helper.CommitOrRollback(tx)

	// 1. Insert user to user table
	result := tx.
		WithContext(ctx.Context()).
		Table(tableUser).
		Create(req)

	if result.Error != nil {
		return result.Error
	}

	// 2. get created user id then insert to user_profile table

	return nil
}

func (repository *UsersRepositoryImpl) FindByUsernameOrEmail(ctx *fiber.Ctx, req string, isEmail bool) (*model.User, error) {

	tx := repository.Db.Begin()
	defer helper.CommitOrRollback(tx)

	var result model.User

	query := tx.WithContext(ctx.Context()).Table(tableUser)

	if isEmail {
		query = query.Where("email = ?", req).Find(&result)
	}

	query = query.Where("username = ?", req).Find(&result)

	err := query.Error
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (repository *UsersRepositoryImpl) FindByEmail(ctx *fiber.Ctx, email string) (*model.User, error) {

	tx := repository.Db.Begin()
	defer helper.CommitOrRollback(tx)

	var result model.User

	query := tx.WithContext(ctx.Context()).Table(tableUser)
	query = query.Where("email = ?", strings.ToLower(email)).Find(&result)

	err := query.Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	return &result, nil
}

func (repository *UsersRepositoryImpl) FindFollowersByUserId(ctx *fiber.Ctx, userId uint, page int, pageSize int, searchQuery string) ([]model.UserWithProfile, int64, error) {

	var userWithProfile []model.UserWithProfile
	var totalCount int64

	tx := repository.Db.Begin()
	defer helper.CommitOrRollback(tx)

	// Offset
	offset := (page - 1) * pageSize

	// Query
	query := tx.WithContext(ctx.Context()).Table(tableFollowers)

	if searchQuery != "" {
		query = query.Where("username LIKE ? ", "%"+searchQuery+"%")
	}

	err := query.Count(&totalCount).Error
	if err != nil {
		return nil, 0, err
	}

	errResult := query.
		Offset(offset).
		Limit(pageSize).
		Find(&userWithProfile).
		Error

	if errResult != nil {
		return nil, 0, errResult
	}

	return userWithProfile, totalCount, nil
}

// func (repository *UsersRepositoryImpl) FindFollowingByUserId(userId int) ([]*model.UserWithProfile, error) {

// }

// func (repository *UsersRepositoryImpl) FindUserProfileById(userId int) (*model.Profile, error) {

// }

// func (repository *UsersRepositoryImpl) UpdateProfileById(userId int, req model.ProfileUpdateRequest) error {

// }

func (repository *UsersRepositoryImpl) CreateFollowUserById(ctx *fiber.Ctx, userId int, userTargetId int) error {

	tx := repository.Db.Begin()
	defer helper.CommitOrRollback(tx)

	queryData := model.FollowUsers{
		FollowerUserId:  userId,
		FollowingUserId: userTargetId,
	}

	err := tx.WithContext(ctx.Context()).Table(tableFollowers).Create(&queryData).Error

	if err != nil {
		return err
	}

	return nil
}

// func (repository *UsersRepositoryImpl) DeleteFollowUserById(followId int) error {

// }
