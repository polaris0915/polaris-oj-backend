package question_service

import (
	"errors"
	"polaris-oj-backend/common"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"

	"polaris-oj-backend/models/dto/question_dto"
	"polaris-oj-backend/models/enums/userrole_enum"

	"gorm.io/gorm"
)

// TODO unfinished: 通过id获取问题详情
func (s *QuestionService) GetQuestionById(session sessions.Session, requestDto any, question *allModels.Question) error {
	request, ok := requestDto.(*question_dto.QuestionQueryRequest)
	if !ok {
		return errors.New("类型断言失败")
	}
	// TODO unfinished: 需要编写具体逻辑
	// 获取当前登录用户
	var userInfo *utils.Claims
	var err error
	// 如果登录信息无效也就不能获取题目信息了
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return err
	}

	// 查询题目信息，如果没有问题，信息也就在question中了
	if err = s.db.Preload("User").First(question, "identity = ?", request.Identity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return gorm.ErrRecordNotFound
		}
		return err
	}
	// fmt.Printf("question: %+v\n", question)

	// 只有创建题目的作者以及管理员才可以查询到问题的详情
	if userInfo.UserRole != userrole_enum.ADMIN_ROLE && userInfo.Identity != question.User.Identity {
		return errors.New("权限不足")
	}
	return nil
}
