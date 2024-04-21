package users

import (
	"todolist-app/model"

	"gorm.io/gorm"
)

type UsersRepository interface {
	Register(req model.RegisterRequest) error
	FindByUsernameOrEmail(req string, isEmail bool) (model.User, error)
}

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {

	return &UsersRepositoryImpl{
		Db: db,
	}
}

func (repository *UsersRepositoryImpl) Register(req model.RegisterRequest) error
func (repository *UsersRepositoryImpl) FindByUsernameOrEmail(req string, isEmail bool) (model.User, error)