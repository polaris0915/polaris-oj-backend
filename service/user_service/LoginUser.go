package user_service

import (
	"errors"

	"polaris-oj-backend/models/dto/user_dto"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"
	"gorm.io/gorm"
)

// 用户登录
func (s *UserService) LoginUser(session sessions.Session, requestDto any, user *allModels.User) (string, error) {
	request, ok := requestDto.(*user_dto.UserLoginRequest)
	if !ok {
		return "", errors.New("类型断言失败")
	}
	if utils.IsAnyBlank(request.UserPassword, request.UserAccount) {
		return "", errors.New("账号或者密码不能为空")
	}
	password := utils.GetMd5(request.UserPassword)

	if err := s.db.Where("userAccount = ? AND userPassword = ?", request.UserAccount, password).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", errors.New("账号或密码错误")
		}
		return "", errors.New("系统错误: 查询失败")
	}

	return utils.GetToken(user.Identity, user.UserAccount, user.UserRole)
}
