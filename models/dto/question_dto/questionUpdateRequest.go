package question_dto

import (
	judgecase "polaris-oj-backend/models/dto/judgecase"
	judgeconfig "polaris-oj-backend/models/dto/judgeconfig"

	"github.com/go-playground/validator/v10"
)

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
