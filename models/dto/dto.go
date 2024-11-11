package dto

import (
	"errors"
	"polaris-oj-backend/models/dto/question_dto"
	"polaris-oj-backend/models/dto/questionsubmit_dto"
	"polaris-oj-backend/models/dto/user_dto"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type RequestDto interface {
	GetValidator() *validator.Validate
}

// deprecated
func BindAndValidateRequest(c *gin.Context, requestDto any) error {
	if request, ok := requestDto.(*DeleteRequest); ok {
		return c.ShouldBindJSON(request)
	}
	// User
	if request, ok := requestDto.(*user_dto.UserAddRequest); ok {

		return c.ShouldBindJSON(request)
	}
	if request, ok := requestDto.(*user_dto.UserUpdateMyUserRequest); ok {
		return c.ShouldBindJSON(request)
	}
	if request, ok := requestDto.(*user_dto.UserLoginRequest); ok {
		return c.ShouldBindJSON(request)
	}
	// Question
	if request, ok := requestDto.(*question_dto.QuestionAddRequest); ok {
		return c.ShouldBindJSON(request)
	}

	if request, ok := requestDto.(*question_dto.QuestionQueryRequest); ok {
		return c.ShouldBindQuery(request)
	}
	if request, ok := requestDto.(*question_dto.QuestionUpdateRequest); ok {
		return c.ShouldBindJSON(request)
	}
	if request, ok := requestDto.(*question_dto.QuestionQueryByPageRequest); ok {
		return c.ShouldBindJSON(request)
	}
	// QuestionSubmit
	if request, ok := requestDto.(*questionsubmit_dto.QuestionSubmitAddRequest); ok {
		return c.ShouldBindJSON(request)
	}
	if request, ok := requestDto.(*questionsubmit_dto.QuestionSubmitQueryRequest); ok {
		return c.ShouldBindJSON(request)
	}

	return errors.New("类型断言错误")
}

type DeleteRequest struct {
	Identity string `json:"identity" validate:"required,uuid"`
}

func (u *DeleteRequest) GetValidator() *validator.Validate {
	return validator.New()
}
