package question_dto

import (
	"github.com/go-playground/validator/v10"
)

type QuestionQueryByPageRequest struct {
	Current   int32  `json:"current"`
	PageSize  int32  `json:"pageSize"`
	SortField string `json:"sortField"`
	SortOrder string `json:"sortOrder"`
	// 可以由以下类别来查询
	Identity string `json:"identity"`
	UserID   string `json:"userId"` // 管理员查询的是否可以用UserID来查询
}

func (u *QuestionQueryByPageRequest) GetValidator() *validator.Validate {
	return validator.New()
}
