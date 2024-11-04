package question_dto

import (
	judgecase "polaris-oj-backend/models/dto/judgecase"
	judgeconfig "polaris-oj-backend/models/dto/judgeconfig"
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
)

// type JudgeCase struct {
// 	Input  string `json:"input"`
// 	Output string `json:"output"`
// }

type QuestionUpdateRequest struct {
	Answer      string                  `json:"answer" validate:"required"`
	Content     string                  `json:"content" validate:"required"`
	Identity    string                  `json:"identity" validate:"required,uuid"`
	JudgeCase   []judgecase.JudgeCase   `json:"judgeCase" validate:"required"`
	JudgeConfig judgeconfig.JudgeConfig `json:"judgeConfig" validate:"required"`
	Tags        []string                `json:"tags" validate:"required"`
	Title       string                  `json:"title" validate:"required"`
}

func (u *QuestionUpdateRequest) GetValidator() *validator.Validate {
	return validator.New()
}

func NewQuestionUpdateRequest() *QuestionUpdateRequest {
	u := new(QuestionUpdateRequest)
	return u
}

func (u *QuestionUpdateRequest) DtoToModel(question *allModels.Question) error {
	// 校验
	if err := u.GetValidator().Struct(u); err != nil {
		return err
	}

	// 内置字段就可以直接copy
	bk := *question
	if err := copier.Copy(question, u); err != nil {
		*question = bk
		return err
	}
	// TODO unfinished: 能不能优化？
	// 如果存在很多字段都要转换，那就会出现很多这样的重复性代码
	// ============自定义字段转换为json字符串格式===============
	var err error
	if question.JudgeCase, err = utils.ModelToJson(u.JudgeCase); err != nil {
		return err
	}
	if question.JudgeConfig, err = utils.ModelToJson(u.JudgeConfig); err != nil {
		return err
	}
	if question.Tags, err = utils.ModelToJson(u.Tags); err != nil {
		return err
	}
	return nil
}
