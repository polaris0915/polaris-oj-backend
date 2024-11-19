package user_service

import (
	"errors"

	"polaris-oj-backend/models/dto/user_dto"
	"polaris-oj-backend/models/enums/userrole_enum"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"
)

// 验证用户注册信息的输入
func validateRegistrationInput(userAccount, userPassword, checkPassword string) error {
	if utils.IsAnyBlank(userAccount, userPassword) {
		return errors.New("参数缺失")
	}
	if len(userAccount) < 4 {
		return errors.New("账号长度不足")
	}
	if len(userPassword) < 8 {
		return errors.New("密码长度不足")
	}
	if checkPassword != userPassword {
		return errors.New("两次密码不一致")
	}
	return nil
}

// 注册用户
func (s *Service) AddUser(request *user_dto.UserAddRequest) error {
	//===============基础字段的检查=========================
	user := new(allModels.User)
	if err := validateRegistrationInput(request.UserAccount, request.UserPassword, request.CheckPassword); err != nil {
		return err
	}
	var cnt int64
	if err := s.db.Model(user).Where("userAccount = ?", request.UserAccount).Count(&cnt).Error; err != nil || cnt > 0 {
		return errors.New("账号已存在")
	}
	//===============新字段的添加=========================
	user.Identity = utils.GetUUID()
	user.UserAccount = request.UserAccount
	user.UserPassword = utils.GetMd5(request.UserPassword)
	user.UserRole = userrole_enum.DEFAULT_ROLE

	if err := s.db.Create(user).Error; err != nil {
		return errors.New("创建用户失败")
	}
	return nil
}

/*
	一个可行的错误处理以及日志输出的构想：
		错误的输出等级
		错误信息
		错误发生在什么地方
		处理谁的业务的时候发生的错误
		myError.New(level, errInfo, )
*/
