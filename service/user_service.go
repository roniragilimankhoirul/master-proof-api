package service

import (
	"master-proof-api/dto"
)

type UserService interface {
	Create(request dto.UserCreateRequest) error
	FindById(email string, nim string) (dto.UserResponse, error)
	Login(request dto.UserLoginRequest) (dto.UserLoginResponse, error)
	ResetPassword(email string) error
	FindByRole(role string) ([]*dto.GetUserByRole, error)
	FindAllTeacher(role string) ([]*dto.GetUserByRole, error)
	UpdatePhotoProfile(request *dto.UpdateUserPhotoRequest) error
}
