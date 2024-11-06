package questionsubmit_dto

import (
	"github.com/go-playground/validator/v10"
)

type QuestionSubmitAddRequest struct {
	QuestionID string `json:"questionId" validate:"required,uuid"` // 问题的identity
	Language   string `json:"language" validate:"required"`        // 编程语言
	// Status     int32  `json:"status"`                              // 提交状态
	Conetnt string `json:"conetnt"`
}

func (u *QuestionSubmitAddRequest) GetValidator() *validator.Validate {
	return validator.New()
}
