package dto

import (
	"master-proof-api/model"
	"mime/multipart"
)

type UserCreateRequest struct {
	Nim      string         `validate:"required,numeric" json:"nim"`
	Name     string         `validate:"required,min=1" json:"name"`
	Role     model.UserRole `validate:"" json:"role"`
	Email    string         `validate:"required,email" json:"email"`
	Password string         `validate:"required,min=8,max=32" json:"password"`
}
type UserLoginRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required,min=8,max=32" json:"password"`
}
type UserPasswordResetRequest struct {
	Email string `validate:"required,email" json:"email"`
}
type UserLoginResponse struct {
	Token string         `json:"token"`
	Role  model.UserRole `json:"role"`
}

type UserResponse struct {
	Nim      string `json:"nim"`
	Role     string `json:"role"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhotoUrl string `json:"photo_url"`
}

type GetUserByRole struct {
	Id       string `json:"id"`
	Nim      string `json:"nim"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	PhotoUrl string `json:"photo_url"`
}

type UpdateUserPhotoRequest struct {
	Id    string                `json:"id" validate:"required"`
	Photo *multipart.FileHeader `json:"photo" validate:"required"`
}
