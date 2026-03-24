package users

import (
	"todolist-app/helper"
	"todolist-app/model"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type FollowRepository interface {
	CreateFollowUserById(ctx *fiber.Ctx, userId int, userTargetId int) error
	FindFollowersByUserId(ctx *fiber.Ctx, userId uint, page int, pageSize int, username string) (*[]model.UserWithProfile, int64, error)
	FindFollowingByUserId(ctx *fiber.Ctx, userId uint, page int, pageSize int, username string) (*[]model.UserWithProfile, int64, error)
	DeleteFollowUserById(ctx *fiber.Ctx, followId int) error
}

type FollowRepositoryImpl struct {
	Db *gorm.DB
}

func NewFollowRepository(db *gorm.DB) FollowRepository {

	return &FollowRepositoryImpl{
		Db: db,
	}
}

var tableUser = "users"
var tableProfile = "user_profiles"
var tableFollowers = "follow_users"

func (repository *FollowRepositoryImpl) FindFollowingByUserId(ctx *fiber.Ctx, userId uint, page int, pageSize int, username string) (*[]model.UserWithProfile, int64, error) {

	// Cari account yang diikuti oleh userId
	tx := repository.Db.Begin()
	defer helper.CommitOrRollback(tx)
	var totalCount int64

	var results *[]model.UserWithProfile

	// Offset
	offset := (page - 1) * pageSize

	query := tx.WithContext(ctx.Context()).Table("follow_users fu").
		Select("fu.user_id, fu.following_user_id, up.role, u.username").
		Joins("join user_profiles up on up.user_id = fu.following_user_id").
		Joins("join users u on u.id = fu.following_user_id").
		Where("fu.user_id = ? and u.username like ?", userId, "%"+username+"%")

	errCount := query.Count(&totalCount).Error
	if errCount != nil {
		return nil, 0, errCount
	}

	errResult := query.
		Offset(offset).
		Limit(pageSize).
		Scan(&results).Error

	if errResult != nil {
		return nil, 0, errResult
	}

	return results, totalCount, nil
}

func (repository *FollowRepositoryImpl) FindFollowersByUserId(ctx *fiber.Ctx, userId uint, page int, pageSize int, username string) (*[]model.UserWithProfile, int64, error) {

	// Cari account yang mengikuti userId
	var results *[]model.UserWithProfile
	var totalCount int64

	tx := repository.Db.Begin()
	defer helper.CommitOrRollback(tx)

	// Offset
	offset := (page - 1) * pageSize

	query := tx.WithContext(ctx.Context()).Table("follow_users fu").
		Select("fu.following_user_id as user_id, fu.user_id as followed_user_id, up.role, u.username").
		Joins("join user_profiles up on up.user_id = fu.user_id").
		Joins("join users u on u.id = fu.user_id").
		Where("fu.following_user_id = ? and u.username like ?", userId, "%"+username+"%")

	errCount := query.Count(&totalCount).Error
	if errCount != nil {
		return nil, 0, errCount
	}

	errResult := query.
		Offset(offset).
		Limit(pageSize).
		Scan(&results).Error

	if errResult != nil {
		return nil, 0, errResult
	}

	return results, totalCount, nil
}

func (repository *FollowRepositoryImpl) CreateFollowUserById(ctx *fiber.Ctx, userId int, userTargetId int) error {

	tx := repository.Db.Begin()
	defer helper.CommitOrRollback(tx)

	queryData := model.FollowUsers{
		UserId:          userId,
		FollowingUserId: userTargetId,
	}

	err := tx.WithContext(ctx.Context()).Table(tableFollowers).Create(&queryData).Error

	if err != nil {
		return err
	}

	return nil
}

func (repository *FollowRepositoryImpl) DeleteFollowUserById(ctx *fiber.Ctx, followId int) error {

	tx := repository.Db.Begin()
	defer helper.CommitOrRollback(tx)

	var model model.FollowUsers

	tx.WithContext(ctx.Context()).
		Table(tableFollowers).
		Where("id = ?", followId).
		Delete(&model)

	return nil
}
