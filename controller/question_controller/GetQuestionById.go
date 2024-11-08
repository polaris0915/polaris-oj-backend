package question_controller

import (
	"net/http"

	"polaris-oj-backend/constant"
	"polaris-oj-backend/models/dto"
	"polaris-oj-backend/models/dto/question_dto"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/question_vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/gin-contrib/sessions"
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
func (qc *QuestionController) GetQuestionById(c *gin.Context) {
	// TODO unfinished: 需要添加中间件，只有管理员或者允许修改的人员才可以更新问题的内容
	// 00. 函数结束固定调用BaseResponse中调用Response
	var response vo.BaseResponse[question_vo.QuestionVO]
	defer func() { // 使用闭包
		response.Response(c, http.StatusOK)
	}()
	// 1. 绑定请求数据到DTO层模型中
	var requestDto dto.RequestDto[*allModels.User] = new(question_dto.QuestionQueryRequest)
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
			2. 第二个参数应该是 请求接口，即DTO层的RequestDto泛型接口
			3. 第三个参数应该是 数据表模型的实体
	*/
	session := sessions.Default(c)
	question := new(allModels.Question)
	if err := qc.questionService.GetQuestionById(session, requestDto, question); err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Message = err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 3. 最终所有的业务逻辑也进行完毕之后，将返回的数据表模型数据交给VO层进行脱敏等操作
	var responVo vo.ResponVo[*allModels.Question] = new(question_vo.QuestionVO)
	if err := responVo.GetResponseVo(question); err != nil {

		response.Code = constant.SYSTEM_ERROR.Code
		response.Message = err.Error()
		return
	}

	// 4. 所有步骤都没有问题之后就可以将vo层处理好的数据返回了
	response.Code = constant.SUCCESS.Code
	response.Data = responVo.(*question_vo.QuestionVO)
}
