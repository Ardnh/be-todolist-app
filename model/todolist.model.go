package model

import "gorm.io/gorm"

type Todolist struct {
	gorm.Model
	UserID     int
	CategoryID int
	Title      string
	IsPublic   bool
}

type TodolistItem struct {
	gorm.Model
	TodolistID int
	Todo       string
	IsDone     bool
}

type FollowsTodolist struct {
	gorm.Model
	FollowingUserID    int
	FollowedTodolistID int
}
