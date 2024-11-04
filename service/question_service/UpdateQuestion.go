package questionservice

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"polaris-oj-backend/common"

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
func (s *QuestionService) UpdateQuestion(session sessions.Session, question *allModels.Question, user *allModels.User) error {
	// 获取题目创建人的用户信息
	var userInfo *utils.Claims
	var err error
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return err
	}
	// TODO: 在调用这个接口的时候就应该是已经通过中间件鉴权了，这边再次验证一下
	// 但是开发阶段可以先关闭
	// if userInfo.UserRole != constant.ADMIN_ROLE {
	// 	return errors.New("权限不足")
	// }

	// 查找修改题目的用户的信息
	if err = s.db.First(user, "identity = ?", userInfo.Identity).Error; err != nil {
		return err
	}
	// 更新问题业务
	dbQuestion := new(allModels.Question)
	// 根据问题的identity查找出问题
	if err := s.db.Model(dbQuestion).First(&dbQuestion, "identity = ?", question.Identity).Error; err != nil {
		return errors.New("没有该数据")
	}
	// TODO: 逻辑有问题
	// 更新dbQuestion即数据库中该问题需要更新的字段
	if err := utils.CopyModels(dbQuestion, question); err != nil {
		return err
	}

	return s.db.Save(dbQuestion).Error
}
