package user_dto

import (
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
)

type UserRegisterRequest struct {
	UserAccount   string `json:"userAccount" validate:"required"`
	UserPassword  string `json:"userPassword" validate:"required"`
	CheckPassword string `json:"checkPassword" validate:"required"`
}

func (u *UserRegisterRequest) GetValidator() *validator.Validate {
	return validator.New()
}

func NewUserRegisterRequest() *UserRegisterRequest {
	u := new(UserRegisterRequest)
	return u
}

func (u *UserRegisterRequest) DtoToModel(user *allModels.User) error {
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
