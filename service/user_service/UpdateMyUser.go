package user_service

import (
	"errors"
	"polaris-oj-backend/common"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/models/dto/user_dto"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"
)

// TODO: 逻辑有问题
func (s *UserService) UpdateMyUser(session sessions.Session, requestDto any, user *allModels.User) error {
	request, ok := requestDto.(*user_dto.UserUpdateMyUserRequest)
	if !ok {
		return errors.New("类型断言失败")
	}
	// 首先判断用户是否自己已经登录，如果没登录则返回
	userInfo := new(utils.Claims)
	var err error
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return errors.New("登录信息过期，请重新登录")
	}

	dbUser := new(allModels.User)
	if err = s.db.First(dbUser, "identity = ?", userInfo.Identity).Error; err != nil {
		return errors.New(constant.SYSTEM_ERROR.Message)
	}

	// TODO: 这个函数有问题，要修改
	if err = utils.CopyModels(dbUser, request); err != nil {
		// return errors.New(constant.PARAMS_ERROR.Message)
		return err
	}
	if err = s.db.Save(dbUser).Error; err != nil {
		return errors.New(constant.PARAMS_ERROR.Message)
	}
	// *user = *dbUser
	return nil

}
