package model

import "gorm.io/gorm"

type FollowUsers struct {
	*gorm.Model
	UserId          int
	FollowingUserId int
}

type FollowUserCreateRequest struct {
	UserId    int
	Following int
}

type UnfollowUserRequest struct {
	UserId   int
	Unfollow int
}
