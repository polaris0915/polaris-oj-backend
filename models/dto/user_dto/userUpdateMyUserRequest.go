package user_dto

import (
	// "polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/go-playground/validator/v10"
	// "github.com/jinzhu/copier"
)

type UserUpdateMyUserRequest struct {
	UserName    string `json:"userName"`
	UserProfile string `json:"userProfile"`
	UserEmail   string `json:"userEmail"`
}

func (u *UserUpdateMyUserRequest) GetValidator() *validator.Validate {
	return validator.New()
}
