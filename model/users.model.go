package model

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
