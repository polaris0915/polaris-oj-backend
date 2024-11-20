package user_service

import (
	"polaris-oj-backend/models/dto/user_dto"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/user_vo"

	"polaris-oj-backend/polaris_logger"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/service/file_service"
	"polaris-oj-backend/utils"
)

func (s *Service) UpdateAvatar(request *user_dto.UserUpdateAvatarRequest) (*user_vo.UserVO, error) {
	user, exits := s.ctx.Get("user")
	if !exits {
		polaris_logger.Error(s.ctx, "未登录")
	}

	var url string
	var err error
	if url, err = file_service.UploadFile(s.ctx, &file_service.UploadFileInfo{
		File: request.File,
	}); err != nil {
		return nil, polaris_logger.Error(s.ctx, err.Error())
	}
	userInfo := user.(*utils.Claims)
	dbUser := &allModels.User{
		UserAvatar: url,
	}
	if err := s.db.Model(&allModels.User{}).Where("identity = ?", userInfo.Identity).Updates(dbUser).Error; err != nil {
		// TODO：需要删除刚才保存的文件
		return nil, polaris_logger.Error(s.ctx, err.Error())
	}

	// 3. 最终所有的业务逻辑也进行完毕之后，将返回的数据表模型数据交给VO层进行脱敏等操作
	var responseVo vo.ResponVo[*allModels.User] = new(user_vo.UserVO)
	if err := responseVo.GetResponseVo(dbUser); err != nil {
		return nil, polaris_logger.Error(s.ctx, err.Error())
	}
	return responseVo.(*user_vo.UserVO), nil

}
