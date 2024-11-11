package question_service

import (
	"errors"
	"polaris-oj-backend/common"
	"polaris-oj-backend/models/dto/question_dto"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/question_vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/gin-contrib/sessions"
)

// 分页查询用户提交的问题
func (s *Service) ListQuestionByPage(request *question_dto.QuestionQueryByPageRequest) (*question_vo.QueryQuestionVO, error) {
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
	var allResults []*allModels.Question
	query := s.db.Model(&allModels.Question{}).
		Preload("User")
	if request.Identity != "" {
		query.Or("identity = ?", request.Identity)
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

	var responseVo vo.ResponVo[[]*allModels.Question] = new(question_vo.QueryQuestionVO)
	if err = responseVo.GetResponseVo(allResults); err != nil {
		return nil, err
	}
	allQueries, _ := responseVo.(*question_vo.QueryQuestionVO)
	allQueries.Total = count

	return responseVo.(*question_vo.QueryQuestionVO), nil
}
