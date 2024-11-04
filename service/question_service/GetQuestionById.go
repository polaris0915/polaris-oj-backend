package questionservice

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"polaris-oj-backend/common"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"gorm.io/gorm"
)

// TODO unfinished: 通过id获取问题详情
func (s *QuestionService) GetQuestionById(session sessions.Session, question *allModels.Question) error {
	// TODO unfinished: 需要编写具体逻辑
	// 获取当前登录用户
	var userInfo *utils.Claims
	var err error
	// 如果登录信息无效也就不能获取题目信息了
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return err
	}

	// 查询题目信息，如果没有问题，信息也就在question中了
	if err = s.db.Preload("User").First(question, "identity = ?", question.Identity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	// fmt.Printf("question: %+v\n", question)

	// 只有创建题目的作者以及管理员才可以查询到问题的详情
	if userInfo.UserRole != constant.ADMIN_ROLE && userInfo.Identity != question.User.Identity {
		return errors.New("权限不足")
	}
	return nil
}
