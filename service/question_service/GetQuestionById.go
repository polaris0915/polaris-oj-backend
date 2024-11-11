package question_service

import (
	"errors"
	"polaris-oj-backend/common"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"

	"polaris-oj-backend/models/dto/question_dto"
	"polaris-oj-backend/models/enums/userrole_enum"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/question_vo"

	"gorm.io/gorm"
)

// TODO unfinished: 通过id获取问题详情
func (s *Service) GetQuestionById(request *question_dto.QuestionQueryRequest) (*question_vo.QuestionVO, error) {
	// TODO unfinished: 需要编写具体逻辑
	// 获取当前登录用户
	session := sessions.Default(s.ctx)
	var userInfo *utils.Claims
	var err error
	// 如果登录信息无效也就不能获取题目信息了
	if userInfo, err = common.GetLoginUser(session); err != nil {
		return nil, err
	}

	// 查询题目信息，如果没有问题，信息也就在question中了
	question := new(allModels.Question)
	if err = s.db.Preload("User").First(question, "identity = ?", request.Identity).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	// 只有创建题目的作者以及管理员才可以查询到问题的详情
	if userInfo.UserRole != userrole_enum.ADMIN_ROLE && userInfo.Identity != question.User.Identity {
		return nil, errors.New("权限不足")
	}

	// 3. 最终所有的业务逻辑也进行完毕之后，将返回的数据表模型数据交给VO层进行脱敏等操作
	var responVo vo.ResponVo[*allModels.Question] = new(question_vo.QuestionVO)
	if err := responVo.GetResponseVo(question); err != nil {
		return nil, err
	}
	return responVo.(*question_vo.QuestionVO), err
}
