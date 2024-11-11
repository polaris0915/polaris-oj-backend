package question_service

import (
	"errors"
	"polaris-oj-backend/common"
	"polaris-oj-backend/models/dto/question_dto"

	"github.com/gin-contrib/sessions"

	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"
)

/*
	TODO Unfinished: 整个问题添加以及修改的逻辑是需要完善的
	因为如果是用户本身去修改题目，肯定是要规范用户修改题目的规则的
	因此这里第一个构想是，用户提交修改请求之后，通知管理员进行审核
	管理员审核无误之后，再将数据进行修改
	创建题目也是同样的逻辑
*/
// 目前只是开发测试基本的功能是否能够接通数据库
// 更新问题
func (s *Service) UpdateQuestion(request *question_dto.QuestionUpdateRequest) error {
	// 获取当前登录用户
	session := sessions.Default(s.ctx)
	// var userInfo *utils.Claims
	var err error
	if _, err = common.GetLoginUser(session); err != nil {
		return err
	}
	// TODO: 在调用这个接口的时候就应该是已经通过中间件鉴权了，这边再次验证一下
	// 但是开发阶段可以先关闭
	// if userInfo.UserRole != constant.ADMIN_ROLE {
	// 	return errors.New("权限不足")
	// }

	// 更新问题业务
	dbQuestion := new(allModels.Question)
	// 根据问题的identity查找出问题
	if err := s.db.Model(dbQuestion).First(&dbQuestion, "identity = ?", request.Identity).Error; err != nil {
		return errors.New("没有该数据")
	}
	// 更新dbQuestion即数据库中该问题需要更新的字段
	if err := utils.CopyModels(dbQuestion, request); err != nil {
		return err
	}

	return s.db.Save(dbQuestion).Error
}
