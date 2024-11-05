package userservice

import (
	"errors"
	"polaris-oj-backend/models/enums/userrole_enum"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"
)

// 验证用户注册信息的输入
func validateRegistrationInput(userAccount, userPassword string) error {
	if utils.IsAnyBlank(userAccount, userPassword) {
		return errors.New("参数缺失")
	}
	if len(userAccount) < 4 {
		return errors.New("账号长度不足")
	}
	if len(userPassword) < 8 {
		return errors.New("密码长度不足")
	}
	return nil
}

// 注册用户
func (s *UserService) RegisterUser(user *allModels.User) error {
	if err := validateRegistrationInput(user.UserAccount, user.UserPassword); err != nil {
		return err
	}
	var cnt int64
	if err := s.db.Model(user).Where("userAccount = ?", user.UserAccount).Count(&cnt).Error; err != nil || cnt > 0 {
		return errors.New("账号已存在")
	}
	user.Identity = utils.GetUUID()
	user.UserPassword = utils.GetMd5(user.UserPassword)
	user.UserRole = userrole_enum.DEFAULT_ROLE

	if err := s.db.Create(user).Error; err != nil {
		return errors.New("创建用户失败")
	}
	return nil
}
