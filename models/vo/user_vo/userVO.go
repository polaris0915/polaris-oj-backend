package uservo

import (
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"time"

	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
)

type UserVO struct {
	// models.ModelValidator
	Identity    string    `json:"identity"`
	UserAvatar  string    `json:"userAvatar"`
	UserName    string    `json:"userName"`
	UserProfile string    `json:"userProfile"`
	UserRole    string    `json:"userRole"`
	CreatedAt   time.Time `json:"createTime"`
	UpdatedAt   time.Time `json:"updateTime"`
}

func (u *UserVO) GetValidator() *validator.Validate {
	return validator.New()
}

func NewUserVo() *UserVO {
	u := new(UserVO)
	return u
}

func (u *UserVO) GetUserVo(user *allModels.User) error {
	if err := copier.Copy(u, user); err != nil {
		return err
	}

	// 校验，没做
	if err := u.GetValidator().Struct(u); err != nil {
		return err
	}
	return nil
}
