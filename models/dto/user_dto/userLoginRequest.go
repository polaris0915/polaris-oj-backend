package user_dto

import (
	"github.com/go-playground/validator/v10"
)

type UserLoginRequest struct {
	// Identity     string `json:"identity,omitempty"`
	UserAccount  string `json:"userAccount" validate:"required"`
	UserPassword string `json:"userPassword" validate:"required"`
}

func (u *UserLoginRequest) GetValidator() *validator.Validate {
	return validator.New()
}
