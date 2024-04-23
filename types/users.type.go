package types

type User struct {
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

type FollowUser struct {
	FollowingUserID int
	FollowedUserID  int
}
