package usercontroller

import (
	"net/http"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/models/vo"

	// "polaris-oj-backend/service/userservice"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// UserLogout
// @Tags 私有方法, 用户
// @Summary 用户退出登录
// @Success 200 {object} vo.BaseResponse "{Code:"0",Data:{...}, Message:""}"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/user/logout [post]
func (uc *UserController) UserLogout(c *gin.Context) {
	// 00. 函数结束固定调用NewSubResponse
	// response对象初始化遵循Code以及Message都是未决的状态
	// 方便之后判断，那么data接口的数据应该也是为nil的状态
	response := vo.NewSubResponse(c, http.StatusOK, constant.UNDEFINED.Code, nil, constant.UNDEFINED.Message)
	defer func() { // 使用闭包
		response.Response()
	}()

	session := sessions.Default(c)
	session.Clear()
	if err := session.Save(); err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = "退出失败"
	}

	response.Code = constant.SUCCESS.Code
	response.Data = true
	response.Message = ""
}
