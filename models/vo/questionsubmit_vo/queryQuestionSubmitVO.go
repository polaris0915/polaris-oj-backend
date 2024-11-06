package questionsubmit_vo

import (
	"errors"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/go-playground/validator/v10"
)

type QueryQuestionSubmitVO struct {
	vo.PageVo
	Records []QuestionSubmitVO `json:"records"`
}

func (u *QueryQuestionSubmitVO) GetValidator() *validator.Validate {
	return validator.New()
}

func (u *QueryQuestionSubmitVO) GetResponseVo(submitQuestions []*allModels.QuestionSubmit) error {
	// TODO: 可以用反射，但好像回影响性能，因为用户可以频繁分页查询
	for _, questionSubmit := range submitQuestions {
		questionSubmitVo := new(QuestionSubmitVO)
		if err := questionSubmitVo.GetResponseVo(questionSubmit); err != nil {
			return errors.New("数据转换失败")
		}
		u.Records = append(u.Records, *questionSubmitVo)
	}
	return nil
}
