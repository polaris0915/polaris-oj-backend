package user_controller

import (
	"errors"
	"net/http"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/models/vo/user_vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"polaris-oj-backend/models/enums/userrole_enum"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// GetLoginUser
// @Tags 用户
// @Summary 获取当前登录用户
// @Success 200 {object} user_vo.UserVO
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/get/user/login [get]
func (uc *UserController) GetLoginUser(c *gin.Context) {
	// 00. 函数结束固定调用BaseResponse中调用Response
	var response vo.BaseResponse[user_vo.UserVO]
	defer func() { // 使用闭包
		response.Response(c, http.StatusOK)
	}()
	// 1. 绑定请求数据到DTO层模型中

	// ================controller特殊的业务需求===================
	// 先尝试获取请求头中的token，没有的话就是未登录
	if _, err := c.Cookie(userrole_enum.USER_LOGIN_STATE); err != nil {
		response.Code = constant.NOT_LOGIN_ERROR.Code
		response.Message = errors.New("未登录").Error()
		return
	}
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
	user := new(allModels.User)
	if err := uc.userService.GetLoginUser(session, nil, user); err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Message = err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 3. 最终所有的业务逻辑也进行完毕之后，将返回的数据表模型数据交给VO层进行脱敏等操作
	var responseVo vo.ResponVo[*allModels.User] = new(user_vo.UserVO)
	if err := responseVo.GetResponseVo(user); err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = err.Error()
		return
	}

	// 4. 所有步骤都没有问题之后就可以将vo层处理好的数据返回了
	response.Code = constant.SUCCESS.Code
	response.Data = responseVo.(*user_vo.UserVO)
}
