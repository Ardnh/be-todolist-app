package model

import "gorm.io/gorm"

type User struct {
	*gorm.Model
	Username string
	Email    string
	Password string
}

type FollowUsers struct {
	*gorm.Model
	FollowingUserId int
	FollowerUserId  int
}

type RegisterRequest struct {
	Username string
	Email    string
	Password string
}

type LoginRequest struct {
	Email    string
	Password string
}

type Profile struct {
	UserId    int
	Bio       string
	Role      string
	Facebook  string
	Instagram string
	LinkedIn  string
	Twitter   string
}

type UserWithProfile struct {
	User
	Profile
}

type ProfileUpdateRequest struct {
	UserId    int
	Bio       string
	Role      string
	Facebook  string
	Instagram string
	LinkedIn  string
	Twitter   string
}

type ProfileCreateRequest struct {
	Bio       string
	Role      string
	Facebook  string
	Instagram string
	LinkedIn  string
	Twitter   string
}

type FollowUserCreateRequest struct {
	UserId   int
	TargetId int
}
