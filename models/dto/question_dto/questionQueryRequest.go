package question_dto

import (
	"polaris-oj-backend/polaris_oj_backend/allModels"
	"polaris-oj-backend/utils"

	"github.com/go-playground/validator"
	"github.com/jinzhu/copier"
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

func NewQuestionQueryRequest() *QuestionQueryRequest {
	u := new(QuestionQueryRequest)
	return u
}

func (u *QuestionQueryRequest) DtoToModel(question *allModels.Question) error {
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
	if u.Tags != nil {
		if question.Tags, err = utils.ModelToJson(u.Tags); err != nil {
			return err
		}
	}
	return nil
}
