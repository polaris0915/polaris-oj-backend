package questioncontroller

import (
	"net/http"

	"polaris-oj-backend/common"
	"polaris-oj-backend/constant"

	"polaris-oj-backend/models/vo"
	questionvo "polaris-oj-backend/models/vo/question_vo"

	"polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/gin-gonic/gin"
)

// DeleteQuestion
// @Tags 私有方法,问题
// @Summary 问题删除
// @Param deleteInfo body common.DeleteRequest true "delete question info"
// @Success 200 {object} vo.BaseResponse "{Code:"0",Data:{...}, Message:""}"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/question/delete [post]
func (qc *QuestionController) DeleteQuestion(c *gin.Context) {
	// TODO unfinished: 需要添加中间件，只有管理员或者允许修改的人员才可以更新问题的内容
	// 00. 函数结束固定调用NewSubResponse
	// response对象初始化遵循Code以及Message都是未决的状态
	// 方便之后判断，那么data接口的数据应该也是为nil的状态
	response := vo.NewSubResponse(c, http.StatusOK, constant.UNDEFINED.Code, nil, constant.UNDEFINED.Message)
	defer func() { // 使用闭包
		response.Response()
	}()
	// 1. 现将请求中的数据转换解析道对应的请求模型中
	deleteRequest := common.NewDeleteRequest()
	if response.Err = c.ShouldBindJSON(deleteRequest); response.Err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 2. 通过DTO层将请求模型转化到数据表模型中
	question := new(allModels.Question)
	if response.Err = deleteRequest.DtoToModel(question); response.Err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// 3. 将DTO层处理好的数据表模型的数据传入service中进行具体的逻辑操作
	// service层要是没有任务问题，那么返回的error也是为空
	// session := sessions.Default(c)
	// user := new(allModels.User) // 获取题目的创建人的用户信息
	// TODO unfinished: 需要引入中间件，鉴权之后才能删除
	if response.Err = qc.questionService.DeleteQuestion(question); response.Err != nil {
		// 暂时统一将service导出的错误定义为系统错误
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 4. 最终所有的业务逻辑也进行完毕之后，将返回的数据表模型数据交给VO层进行脱敏等操作
	questionVo := questionvo.NewQuestionVO()
	if response.Err = questionVo.GetQuestionVO(question, nil); response.Err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}

	// 5. 所有步骤都没有问题之后就可以将vo层处理好的数据返回了
	response.Code = constant.SUCCESS.Code
	response.Data = questionVo
	response.Message = ""
}
