package users

import (
	"todolist-app/model"
	"todolist-app/types"

	"gorm.io/gorm"
)

type UsersRepository interface {
	Register(req types.RegisterRequest) error
	FindByUsernameOrEmail(req string, isEmail bool) (*model.User, error)
}

type UsersRepositoryImpl struct {
	Db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) UsersRepository {

	return &UsersRepositoryImpl{
		Db: db,
	}
}

func (repository *UsersRepositoryImpl) Register(req types.RegisterRequest) error {

	return nil
}
func (repository *UsersRepositoryImpl) FindByUsernameOrEmail(req string, isEmail bool) (*model.User, error) {

	return nil, nil
}
