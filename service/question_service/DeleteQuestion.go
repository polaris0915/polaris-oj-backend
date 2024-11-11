package question_service

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
func (s *Service) DeleteQuestion(request *dto.DeleteRequest) error {
	// 校验当前登录用户信息
	session := sessions.Default(s.ctx)
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

	res := s.db.Delete(&allModels.Question{}, "identity = ?", request.Identity)
	if res.Error != nil {
		return res.Error
	} else if res.RowsAffected == 0 {
		return errors.New("已经删除")
	}
	return nil
}
