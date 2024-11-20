package user_vo

import (
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type UserVO struct {
	Identity    string    `json:"identity"`
	UserAvatar  string    `json:"userAvatar"`
	UserName    string    `json:"userName"`
	UserProfile string    `json:"userProfile"`
	UserRole    string    `json:"userRole"`
	CreatedAt   time.Time `json:"createTime"`
	UpdatedAt   time.Time `json:"updateTime"`
	UserEmail   string    `json:"userEmail"`
}

func (u *UserVO) GetValidator() *validator.Validate {
	return validator.New()
}

func (u *UserVO) GetResponseVo(user *allModels.User) error {
	if err := copier.Copy(u, user); err != nil {
		return err
	}

	// 校验，没做
	if err := u.GetValidator().Struct(u); err != nil {
		return err
	}
	return nil
}
