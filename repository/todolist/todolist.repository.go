package todolist

import "gorm.io/gorm"

type TodolistRepository interface {
}

type TodolistRepositoryImpl struct {
	Db *gorm.DB
}

func NewTodolistRepository(db *gorm.DB) TodolistRepository {

	return &TodolistRepositoryImpl{
		Db: db,
	}
}
