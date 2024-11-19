package user_service

import (
	"errors"
	"polaris-oj-backend/common"
	"polaris-oj-backend/models/dto/user_dto"
	"polaris-oj-backend/models/enums/userrole_enum"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/user_vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"
	"gorm.io/gorm"
)

// 用户登录
func (s *Service) LoginUser(request *user_dto.UserLoginRequest) (*user_vo.UserVO, error) {
	if utils.IsAnyBlank(request.UserPassword, request.UserAccount) {
		return nil, errors.New("账号或者密码不能为空")
	}
	password := utils.GetMd5(request.UserPassword)

	user := new(allModels.User)
	if err := s.db.Where("userAccount = ? AND userPassword = ?", request.UserAccount, password).First(user).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("账号或密码错误")
		}
		return nil, errors.New("系统错误: 查询失败")
	}

	var token string
	token, err := utils.GetToken(user.Identity, user.UserAccount, user.UserRole)
	if err != nil {
		return nil, errors.New("系统错误: 无法生成token")
	}

	session := sessions.Default(s.ctx)
	if err = common.SetCookies(session, userrole_enum.USER_LOGIN_STATE, token); err != nil {
		return nil, errors.New("系统错误: 无法设置cookie")
	}

	// 3. 最终所有的业务逻辑也进行完毕之后，将返回的数据表模型数据交给VO层进行脱敏等操作
	var responseVo vo.ResponVo[*allModels.User] = new(user_vo.UserVO)
	if err := responseVo.GetResponseVo(user); err != nil {
		return nil, err
	}
	return responseVo.(*user_vo.UserVO), nil
}
