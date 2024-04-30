package model

import "gorm.io/gorm"

type Category struct {
	*gorm.Model
	Name string
}

type CategoryCreateRequest struct {
	Name string `json:"name" validate:"required"`
}

type CategoryUpdateRequest struct {
	Id   int    `json:"id" validate:"required,number"`
	Name string `json:"name" validate:"required"`
}

type CategoryDeleteRequest struct {
	Id int `json:"id" validate:"required,number"`
}
