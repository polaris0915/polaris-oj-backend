package questionsubmit_controller

import (
	"net/http"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/models/dto/questionsubmit_dto"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/questionsubmit_vo"
	"polaris-oj-backend/service/questionsubmit_service"

	"github.com/gin-gonic/gin"
)

// ListQuestionSubmitByPage
// @Tags 问题提交
// @Summary 分页查询提交问题
// @Param queryInfo body questionsubmit_dto.QuestionSubmitQueryRequest true "query questionSubmit info"
// @Success 200 {object} vo.BaseResponse[questionsubmit_vo.QueryQuestionSubmitVO]
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/question/question_submit/list/page [post]
func ListQuestionSubmitByPage(c *gin.Context) {
	// TODO unfinished: 需要添加中间件，只有管理员或者允许修改的人员才可以更新问题的内容
	// 00. 函数结束固定调用BaseResponse中调用Response
	var response vo.BaseResponse[questionsubmit_vo.QueryQuestionSubmitVO]
	defer func() { // 使用闭包
		response.Response(c, http.StatusOK)
	}()
	// 1. 绑定请求数据到DTO层模型中
	requestDto := new(questionsubmit_dto.QuestionSubmitQueryRequest)
	if err := c.ShouldBindJSON(requestDto); err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Message = err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 2. 将DTO层处理好的数据表模型的数据传入service中进行具体的逻辑操作
	// service层要是没有任务问题，那么返回的error也是为空
	var err error
	if response.Data, err = questionsubmit_service.NewService(c).ListQuestionSubmitByPage(requestDto); err != nil {
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
