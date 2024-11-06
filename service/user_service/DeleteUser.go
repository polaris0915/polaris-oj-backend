package user_service

import (
	"errors"

	"github.com/gin-contrib/sessions"

	"polaris-oj-backend/common"
	"polaris-oj-backend/models/dto"
	"polaris-oj-backend/models/enums/userrole_enum"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"
)

// 删除问题
func (s *UserService) DeleteUser(session sessions.Session, requestDto any, user *allModels.User) error {
	request, ok := requestDto.(*dto.DeleteRequest)
	if !ok {
		return errors.New("类型断言失败")
	}
	// 校验当前登录用户信息
	var userInfo *utils.Claims
	var err error
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return err
	}
	if userInfo.UserRole != userrole_enum.ADMIN_ROLE {
		return errors.New("无权限")
	}
	// 校验删除Identity
	if utils.IsAnyBlank(request.Identity) {
		return errors.New("参数不正确")
	}

	// 已经删除就不要再删除了
	res := s.db.Delete(&allModels.User{}, "identity = ?", request.Identity)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected == 0 {
		return errors.New("已经删除")
	}
	return nil
}