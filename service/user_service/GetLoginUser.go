package user_service

import (
	"errors"
	"polaris-oj-backend/common"
	"polaris-oj-backend/database/mysql"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"
)

// 获取当前登录用户信息
func (s *UserService) GetLoginUser(session sessions.Session, requestDto any, user *allModels.User) error {
	// 问题是如何获取用户信息呢？仅凭一个cookie
	var userInfo *utils.Claims
	var err error
	// 通过common.GetLoginUser解析用户信息，如果没有则是过期
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return err
	}
	// 查询用户信息
	// TODO unfinished: 可以引入redis，提升性能
	identity := userInfo.Identity

	if err := mysql.DB.Model(user).First(user, "identity = ?", identity).Error; err != nil {
		return errors.New("系统错误")
	}
	return nil
}
