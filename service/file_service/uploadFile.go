package file_service

import (
	"polaris-oj-backend/models/dto/file_dto"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/file_vo"
	"polaris-oj-backend/polaris_logger"
)

// func (s *Service) UploadFile(request *file_dto.FileUploadRequest) (*file_vo.FileVO, error) {
// 	// 校验文件
// 	if err, ok := FileValidator.ValidateFile(request.File); !ok {
// 		return nil, polaris_logger.Error(s.ctx, err.Error())
// 	}
// 	// 给文件生成随机名
// 	fileName := FileValidator.GetRandomFileName(request.FileName)
// 	/*
// 		已经登录的用户统一用用户的uuid作为文件夹的名字
// 	*/
// 	if user, exits := s.ctx.Get("user"); exits {
// 		// db_user := &allModels.User{
// 		// 	UserAvatar: fileName,
// 		// }
// 		userInfo, _ := user.(*utils.Claims)
// 		uploadPath := "./upload/" + userInfo.Identity + "/" + fileName
// 		// 保存都本地
// 		if err := s.ctx.SaveUploadedFile(request.File, uploadPath); err != nil {
// 			return nil, polaris_logger.Error(s.ctx, err.Error())
// 		}
// 		// // 更新数据库
// 		// if err := s.db.Model(db_user).Where("identity = ?", userInfo.Identity).Updates(db_user).Error; err != nil {
// 		// 	return nil, polaris_logger.Error(s.ctx, err.Error())
// 		// }
// 		// 组织响应数据
// 		var responseVo vo.ResponVo[string] = new(file_vo.FileVO)
// 		if err := responseVo.GetResponseVo("/api/file/download/" + userInfo.Identity + "/" + fileName); err != nil {
// 			return nil, polaris_logger.Error(s.ctx, "系统错误")
// 		}
// 		return responseVo.(*file_vo.FileVO), nil
// 	}
// 	return nil, polaris_logger.Error(s.ctx, "未登录")

// }

func (s *Service) UploadFile(request *file_dto.FileUploadRequest) (*file_vo.FileVO, error) {
	// 进行上传文件的操作

	fileUrl, err := UploadFile(s.ctx, &UploadFileInfo{
		File: request.File,
	})
	if err != nil {
		return nil, polaris_logger.Error(s.ctx, err.Error())
	}
	// 组织响应数据
	var responseVo vo.ResponVo[string] = new(file_vo.FileVO)
	if err := responseVo.GetResponseVo(fileUrl); err != nil {
		return nil, polaris_logger.Error(s.ctx, "系统错误")
	}
	return responseVo.(*file_vo.FileVO), nil
}
