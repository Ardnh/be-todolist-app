package app

import (
	"os"
	"todolist-app/helper"
	"todolist-app/schema"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbConnection() *gorm.DB {

	postgresDsn := os.Getenv("APP_DSN")
	db, err := gorm.Open(postgres.Open(postgresDsn), &gorm.Config{})

	helper.PanicIfError(err)

	db.AutoMigrate(
		&schema.Category{},
		&schema.Todolist{},
		&schema.TodolistItem{},
		&schema.FollowsTodolist{},
		&schema.UserProfile{},
		&schema.FollowUser{},
		&schema.Users{},
	)

	return db
}
