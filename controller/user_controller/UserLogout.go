package user_controller

import (
	"net/http"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLogout
// @Tags 用户
// @Summary 用户退出登录
// @Success 200 {object} vo.BaseResponse[bool]
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/user/logout [post]
func (uc *UserController) UserLogout(c *gin.Context) {
	// 00. 函数结束固定调用BaseResponse中调用Response
	var response vo.BaseResponse[bool]
	defer func() { // 使用闭包
		response.Response(c, http.StatusOK)
	}()

	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		response.Code = constant.SYSTEM_ERROR.Code

		response.Message = "退出失败"
	}

	response.Code = constant.SUCCESS.Code
	response.Data = utils.GetBoolPtr(true)

}
