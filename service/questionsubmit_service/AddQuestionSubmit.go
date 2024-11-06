package questionsubmit_service

import (
	"errors"
	"polaris-oj-backend/common"
	"polaris-oj-backend/models/dto/questionsubmit_dto"
	"polaris-oj-backend/models/enums/questionsubmitlanguage_enum"
	"polaris-oj-backend/models/enums/questionsubmitstatus_enum"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"
)

// 添加问题
func (s *QuestionSubmitService) AddQuestionSubmit(session sessions.Session, requestDto any, questionSubmit *allModels.QuestionSubmit) error {
	request, ok := requestDto.(questionsubmit_dto.QuestionSubmitAddRequest)
	if !ok {
		return errors.New("类型断言失败")
	}
	// 首先判断session是否有效
	var loginUserInfo *utils.Claims
	var err error
	if loginUserInfo, err = common.GetLoginUser(session); err != nil {
		return errors.New("登录信息过期，请重新登录")
	}

	//===============基础字段的检查=========================
	// 判断编程语言是否合法
	if _, ok := questionsubmitlanguage_enum.LANGUAGE[request.Language]; !ok {
		return errors.New("编程语言错误")
	}
	// 判断问题identity是否存在

	if err = s.db.First(&allModels.QuestionSubmit{}, "identity = ?", request.QuestionID).Error; err != nil {
		return errors.New("没有该问题")
	}
	//===============新字段的添加=========================
	// 设置questionSubmit的信息
	questionSubmit.Identity = utils.GetUUID()
	questionSubmit.UserID = loginUserInfo.Identity
	questionSubmit.QuestionID = request.QuestionID
	questionSubmit.Status = int32(questionsubmitstatus_enum.WAITING.Value)
	questionSubmit.JudgeInfo = "{}"

	// 保存到数据库
	if err = s.db.Save(questionSubmit).Error; err != nil {
		return errors.New("插入失败")
	}
	// TODO：执行判题任务
	return nil
}
