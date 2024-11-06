package questionsubmit_vo

import (
	"errors"
	"polaris-oj-backend/models/vo/question_vo"
	"polaris-oj-backend/models/vo/user_vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/jinzhu/copier"
)

type QuestionSubmitVO struct {
	Identity   string                 `json:"identity"`   // 唯一ID
	Language   string                 `json:"language"`   // 编程语言
	Conetnt    string                 `json:"conetnt"`    // 用户代码
	JudgeInfo  string                 `json:"judgeInfo"`  // 判题信息（json 对象）
	Status     int32                  `json:"status"`     // 判题状态（0 - 待判题、1 - 判题中、2 - 成功、3 - 失败）
	QuestionID string                 `json:"questionId"` // 题目 id
	UserID     string                 `json:"userId"`     // 题目创建用户唯一ID
	CreatedAt  time.Time              `json:"created_at"` // 创建时间
	UpdatedAt  time.Time              `json:"updated_at"` // 更新时间
	UserVo     user_vo.UserVO         `json:"user"`       // 提交用户信息
	QuestionVo question_vo.QuestionVO `json:"question"`   // 题目信息
}

func (u *QuestionSubmitVO) GetValidator() *validator.Validate {
	return validator.New()
}

func (u *QuestionSubmitVO) GetResponseVo(questionSubmit *allModels.QuestionSubmit) error {
	// ============能直接拷贝的字段=======================
	if err := copier.Copy(u, questionSubmit); err != nil {
		return err
	}
	// ============自定义字段转换为json字符串格式===============
	// 将questionSubmit中的User脱敏
	var userVo = new(user_vo.UserVO)
	if err := userVo.GetResponseVo(questionSubmit.User); err != nil {
		return errors.New("数据转换失败")
	}
	u.UserVo = *userVo

	// 将questionSubmit中的Question脱敏
	var questionVo = new(question_vo.QuestionVO)
	if err := questionVo.GetResponseVo(questionSubmit.Question); err != nil {
		return errors.New("数据转换失败")
	}
	u.QuestionVo = *questionVo

	// 校验
	if err := u.GetValidator().Struct(u); err != nil {
		return err
	}
	return nil
}
