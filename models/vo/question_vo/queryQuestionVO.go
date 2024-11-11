package question_vo

import (
	"errors"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/go-playground/validator/v10"
)

type QueryQuestionVO struct {
	vo.PageVo
	Records []QuestionVO `json:"records"`
}

func (u *QueryQuestionVO) GetValidator() *validator.Validate {
	return validator.New()
}

func (u *QueryQuestionVO) GetResponseVo(questions []*allModels.Question) error {
	// TODO: 可以用反射，但好像回影响性能，因为用户可以频繁分页查询
	for _, question := range questions {
		questionVo := new(QuestionVO)
		if err := questionVo.GetResponseVo(question); err != nil {
			return errors.New("数据转换失败")
		}
		u.Records = append(u.Records, *questionVo)
	}
	return nil
}
