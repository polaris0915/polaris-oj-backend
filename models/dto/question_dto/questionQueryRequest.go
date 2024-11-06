package question_dto

import (
	"github.com/go-playground/validator/v10"
)

// 用户的查询请求有可能根据多种情况查询，例如根据id，标题，内容等等，甚至创建人的信息
// 所以保留这些比较关键性的字段
type QuestionQueryRequest struct {
	Answer   string   `json:"answer"`
	Content  string   `json:"content"`
	Identity string   `json:"identity"`
	Tags     []string `json:"tags"`
	Title    string   `json:"title"`
	UserId   string   `json:"userId"`
}

func (u *QuestionQueryRequest) GetValidator() *validator.Validate {
	return validator.New()
}
