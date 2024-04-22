package model

type Todolist struct {
	UserID     int
	CategoryID int
	Title      string
	IsPublic   bool
}

type TodolistItem struct {
	TodolistID int
	Todo       string
	IsDone     bool
}

type FollowsTodolist struct {
	FollowingUserID    int
	FollowedTodolistID int
}
