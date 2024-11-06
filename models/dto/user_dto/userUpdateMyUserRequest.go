package user_dto

import (
	// "polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/go-playground/validator/v10"
	// "github.com/jinzhu/copier"
)

type UserUpdateMyUserRequest struct {
	UserName    string `json:"userName"`
	UserAvatar  string `json:"userAvatar"`
	UserProfile string `json:"userProfile"`
}

func (u *UserUpdateMyUserRequest) GetValidator() *validator.Validate {
	return validator.New()
}
