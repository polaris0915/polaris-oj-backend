package user_dto

import (
	"github.com/go-playground/validator/v10"
)

type UserAddRequest struct {
	UserAccount   string `json:"userAccount" validate:"required"`
	UserPassword  string `json:"userPassword" validate:"required"`
	CheckPassword string `json:"checkPassword" validate:"required"`
}

func (u *UserAddRequest) GetValidator() *validator.Validate {
	return validator.New()
}
