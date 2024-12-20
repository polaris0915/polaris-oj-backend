package questionsubmit_service

import (
	"errors"
	"polaris-oj-backend/common"
	"polaris-oj-backend/models/dto/questionsubmit_dto"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/questionsubmit_vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/gin-contrib/sessions"
)

// 分页查询用户提交的问题
func (s *Service) ListQuestionSubmitByPage(request *questionsubmit_dto.QuestionSubmitQueryRequest) (*questionsubmit_vo.QueryQuestionSubmitVO, error) {
	// 首先判断session是否有效
	session := sessions.Default(s.ctx)
	// var loginUserInfo *utils.Claims
	var err error
	if _, err = common.GetLoginUser(session); err != nil {
		return nil, errors.New("登录信息过期，请重新登录")
	}
	// 分页查询业务
	var pageSize int
	if request.PageSize == 0 {
		pageSize = 0
	} else {
		pageSize = int(request.PageSize)
	}
	var currentPage int
	if request.Current == 0 {
		currentPage = 1
	} else {
		currentPage = (int(request.Current) - 1) * pageSize
	}
	// TODO EMERGENCY: DTO需要重构
	var allResults []*allModels.QuestionSubmit
	query := s.db.Model(&allModels.QuestionSubmit{}).
		Joins("User"). // 加载User关联时指定identity
		Preload("Question").
		Preload("Question.User").
		Or("status = ?", request.Status)
	if request.Language != "" {
		query.Or("language = ?", request.Language)
	}
	if request.QuestionID != "" {
		query.Or("questionId = ?", request.QuestionID)
	}
	if request.UserID != "" {
		query.Or("userId = ?", request.UserID)
	}
	if request.SortField != "" {
		order := request.SortField
		if request.SortOrder != "" {
			order += (" " + request.SortOrder)
		}
		query.Order(order)
	}

	if pageSize > 0 {
		offset := (currentPage - 1) * pageSize
		query.Limit(pageSize).Offset(offset)
	}
	var count int64
	if err = query.Count(&count).Find(&allResults).Error; err != nil {
		return nil, err
	}

	var responseVo vo.ResponVo[[]*allModels.QuestionSubmit] = new(questionsubmit_vo.QueryQuestionSubmitVO)
	if err = responseVo.GetResponseVo(allResults); err != nil {
		return nil, err
	}
	allQueries, _ := responseVo.(*questionsubmit_vo.QueryQuestionSubmitVO)
	allQueries.Total = count

	return responseVo.(*questionsubmit_vo.QueryQuestionSubmitVO), nil
}
