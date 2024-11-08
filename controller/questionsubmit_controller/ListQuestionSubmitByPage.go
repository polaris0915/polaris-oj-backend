package questionsubmit_controller

import (
	"net/http"

	"polaris-oj-backend/constant"
	"polaris-oj-backend/models/dto"
	"polaris-oj-backend/models/dto/questionsubmit_dto"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/questionsubmit_vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// ListQuestionSubmitByPage
// @Tags 问题提交
// @Summary 分页查询提交问题
// @Param queryInfo body questionsubmit_dto.QuestionSubmitQueryRequest true "query questionSubmit info"
// @Success 200 {object} vo.BaseResponse[[]questionsubmit_vo.QuestionSubmitVO]
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/question/question_submit/list/page [post]
func (qc *QuestionSubmitController) ListQuestionSubmitByPage(c *gin.Context) {
	// TODO unfinished: 需要添加中间件，只有管理员或者允许修改的人员才可以更新问题的内容
	// 00. 函数结束固定调用BaseResponse中调用Response
	var response vo.BaseResponse[questionsubmit_vo.QueryQuestionSubmitVO]
	defer func() { // 使用闭包
		response.Response(c, http.StatusOK)
	}()
	// 1. 绑定请求数据到DTO层模型中
	var requestDto dto.RequestDto[*allModels.QuestionSubmit] = new(questionsubmit_dto.QuestionSubmitQueryRequest)
	/*
		添加新的接口后，要去dto.BindAndValidateRequest添加新的断言
		以便请求数据能够成功转换到对应的请求模型中
	*/
	if err := dto.BindAndValidateRequest(c, requestDto); err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Message = err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 2. 将DTO层处理好的数据表模型的数据传入service中进行具体的逻辑操作
	// service层要是没有任务问题，那么返回的error也是为空
	/*
		传入service层的参数遵循一个原则：
			1. 第一个参数应该是 当前请求的session
			2. 第二个参数应该是 请求结构体，即每个api在DTO层的结构体
			3. 第三个参数应该是 数据表模型的实体
	*/
	session := sessions.Default(c)
	var data map[string]any
	var err error
	if data, err = qc.questionSubmitService.ListQuestionSubmitByPage(session, requestDto, nil); err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Message = err.Error()
		return
	}
	// ================controller特殊的业务需求===================
	all_questionSubmits, ok := data["data"].([]*allModels.QuestionSubmit)
	if !ok {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Message = "数据转换错误"
	}
	// ================controller特殊的业务需求===================
	// 4. 最终所有的业务逻辑也进行完毕之后，将返回的数据表模型数据交给VO层进行脱敏等操作
	// TODO question：这里是否有问题
	var responseVo vo.ResponVo[[]*allModels.QuestionSubmit] = new(questionsubmit_vo.QueryQuestionSubmitVO)
	if err = responseVo.GetResponseVo(all_questionSubmits); err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Message = err.Error()
		return
	}
	// ================controller特殊的业务需求===================
	// TODO continue: 完善responseVo中的pageVo
	// allQueries, ok := responseVo.(*questionsubmit_vo.QueryQuestionSubmitVO)
	// ================controller特殊的业务需求===================
	// 5. 所有步骤都没有问题之后就可以将vo层处理好的数据返回了
	response.Code = constant.SUCCESS.Code
	response.Data = responseVo.(*questionsubmit_vo.QueryQuestionSubmitVO)
}
