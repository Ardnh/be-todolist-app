package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username string
	Email    string
	Password string
}

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

type LoginRequest struct {
	Username string
	Password string
}

type UserUpdateRequest struct {
	Bio       string
	Role      string
	Facebook  string
	Instagram string
	LinkedIn  string
	Twitter   string
}

type FollowUserCreateRequest struct {
	UserId   string
	TargetId string
}
