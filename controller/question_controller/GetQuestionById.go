package question_controller

import (
	"net/http"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/models/dto/question_dto"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/question_vo"
	"polaris-oj-backend/service/question_service"

	"github.com/gin-gonic/gin"
)

// GetQuestionById
// @Tags 问题
// @Summary 通过问题的Identity获取问题详情
// @Param Identity query string true "identity"
// @Success 200 {object} vo.BaseResponse[question_vo.QuestionVO]
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/question/get [get]
func GetQuestionById(c *gin.Context) {
	// TODO unfinished: 需要添加中间件，只有管理员或者允许修改的人员才可以更新问题的内容
	// 00. 函数结束固定调用BaseResponse中调用Response
	var response vo.BaseResponse[question_vo.QuestionVO]
	defer func() { // 使用闭包
		response.Response(c, http.StatusOK)
	}()
	// 1. 绑定请求数据到DTO层模型中
	requestDto := new(question_dto.QuestionQueryRequest)
	if err := c.ShouldBindQuery(requestDto); err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Message = err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 2. 将DTO层处理好的数据表模型的数据传入service中进行具体的逻辑操作
	// service层要是没有任务问题，那么返回的error也是为空
	var err error
	if response.Data, err = question_service.NewService(c).GetQuestionById(requestDto); err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 3. 所有步骤都没有问题之后就可以将vo层处理好的数据返回了
	response.Code = constant.SUCCESS.Code
}
