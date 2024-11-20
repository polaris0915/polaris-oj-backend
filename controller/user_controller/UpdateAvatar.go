package user_controller

import (
	"net/http"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/service/user_service"

	"polaris-oj-backend/models/dto/user_dto"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/user_vo"

	"polaris-oj-backend/polaris_logger"

	"github.com/gin-gonic/gin"
)

// UpdateAvatar
// @Tags 用户
// @Summary 用户更新头像
// @Accept multipart/form-data
// @Param file formData file true "File to upload"
// @Success 200 {object} vo.BaseResponse[user_vo.UserVO]
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/user/update-avatar [post]
func UpdateAvatar(c *gin.Context) {
	// TODO unfinished: 需要添加中间件，只有管理员或者允许修改的人员才可以更新问题的内容
	// 00. 函数结束固定调用BaseResponse中调用Response
	var response vo.BaseResponse[user_vo.UserVO]
	defer func() { // 使用闭包
		response.Response(c, http.StatusOK)
	}()
	// 1. 绑定请求数据到DTO层模型中
	requestDto := new(user_dto.UserUpdateAvatarRequest)
	if err := c.ShouldBind(requestDto); err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Message = polaris_logger.Error(c, err.Error()).Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 2. 将DTO层处理好的数据表模型的数据传入service中进行具体的逻辑操作
	// service层要是没有任务问题，那么返回的error也是为空
	var err error
	if response.Data, err = user_service.NewService(c).UpdateAvatar(requestDto); err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Message = err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 3. 所有步骤都没有问题之后就可以将vo层处理好的数据返回了
	response.Code = constant.SUCCESS.Code
}
