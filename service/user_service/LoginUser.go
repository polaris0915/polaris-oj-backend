package userservice

import (
	"errors"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"gorm.io/gorm"
)

// 用户登录
func (s *UserService) LoginUser(user *allModels.User) (string, error) {
	if utils.IsAnyBlank(user.UserPassword, user.UserAccount) {
		return "", errors.New("账号或者密码不能为空")
	}
	password := utils.GetMd5(user.UserPassword)

	if err := s.db.Where("userAccount = ? AND userPassword = ?", user.UserAccount, password).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("账号或密码错误")
		}
		return "", errors.New("系统错误: 查询失败")
	}

	return utils.GetToken(user.Identity, user.UserAccount, user.UserRole)
}
