package user_dto

import (
	"mime/multipart"

	"github.com/go-playground/validator/v10"
)

type UserUpdateAvatarRequest struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

func (u *UserUpdateAvatarRequest) GetValidator() *validator.Validate {
	return validator.New()
}
