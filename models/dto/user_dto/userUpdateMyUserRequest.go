package user_dto

import (
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
)

type UserUpdateMyUserRequest struct {
	UserName    string `json:"userName"`
	UserAvatar  string `json:"userAvatar"`
	UserProfile string `json:"userProfile"`
}

func (u *UserUpdateMyUserRequest) GetValidator() *validator.Validate {
	return validator.New()
}

func NewUserUpdateMyUserRequest() *UserUpdateMyUserRequest {
	u := new(UserUpdateMyUserRequest)
	return u
}

func (u *UserUpdateMyUserRequest) DtoToModel(user *allModels.User) error {
	// 校验
	if err := u.GetValidator().Struct(u); err != nil {
		return err
	}

	// 内置字段就可以直接copy
	bk := *user
	if err := copier.Copy(user, u); err != nil {
		*user = bk
		return err
	}
	// ============自定义字段转换为json字符串格式===============

	return nil
}
