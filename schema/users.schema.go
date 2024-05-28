package schema

import "gorm.io/gorm"

type Users struct {
	*gorm.Model
	Username string
	Email    string
	Password string
}

type UserProfile struct {
	*gorm.Model
	UserID    int
	Bio       string
	Role      string
	Facebook  string
	Instagram string
	Linkedin  string
	twitter   string
}

type FollowUsers struct {
	*gorm.Model
	UserID       int
	FollowUserID int
}
