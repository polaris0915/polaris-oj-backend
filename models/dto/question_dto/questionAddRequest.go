package question_dto

import (
	judgecase "polaris-oj-backend/models/dto/judgecase"
	judgeconfig "polaris-oj-backend/models/dto/judgeconfig"

	"github.com/go-playground/validator/v10"
)

type QuestionAddRequest struct {
	Answer      string                  `json:"answer" validate:"required"`
	Content     string                  `json:"content" validate:"required"`
	Identity    string                  `json:"identity"`
	JudgeCase   []judgecase.JudgeCase   `json:"judgeCase" validate:"required"`
	JudgeConfig judgeconfig.JudgeConfig `json:"judgeConfig" validate:"required"`
	Tags        []string                `json:"tags"`
	Title       string                  `json:"title" validate:"required"`
}

func (u *QuestionAddRequest) GetValidator() *validator.Validate {
	return validator.New()
}
