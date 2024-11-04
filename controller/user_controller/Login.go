package usercontroller

import (
	"net/http"
	"polaris-oj-backend/common"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/models/dto/user_dto"
	"polaris-oj-backend/models/vo"
	uservo "polaris-oj-backend/models/vo/user_vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"

	// "polaris-oj-backend/service/userservice"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// Login
// @Tags 公共方法, 用户
// @Summary 用户登陆
// @Param login body user_dto.UserLoginRequest true "user login infos"
// @Success 200 {object} vo.BaseResponse "{Code:"0",Data:{...}, Message:""}"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/user/login [post]
func (uc *UserController) Login(c *gin.Context) {
	// 00. 函数结束固定调用NewSubResponse
	// response对象初始化遵循Code以及Message都是未决的状态
	// 方便之后判断，那么data接口的数据应该也是为nil的状态
	response := vo.NewSubResponse(c, http.StatusOK, constant.UNDEFINED.Code, nil, constant.UNDEFINED.Message)
	defer func() { // 使用闭包
		response.Response()
	}()
	// 1. 现将请求中的数据转换解析道对应的请求模型中
	userLoginRequest := user_dto.NewUserLoginRequest()
	if response.Err = c.ShouldBindJSON(userLoginRequest); response.Err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 2. 通过DTO层将请求模型转化到数据表模型中
	user := new(allModels.User)
	if response.Err = userLoginRequest.DtoToModel(user); response.Err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// 3. 将DTO层处理好的数据表模型的数据传入service中进行具体的逻辑操作
	// service层要是没有任务问题，那么返回的error也是为空
	// TODO unfinished: token看看能不能优化成结构体对象
	var token string
	token, response.Err = uc.userService.LoginUser(user)
	if response.Err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================
	// 设置cookies
	session := sessions.Default(c)
	response.Err = common.SetCookies(session, constant.USER_LOGIN_STATE, token)
	if response.Err = common.SetCookies(session, constant.USER_LOGIN_STATE, token); response.Err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================
	// 4. 最终所有的业务逻辑也进行完毕之后，将返回的数据表模型数据交给VO层进行脱敏等操作
	userVo := uservo.NewUserVo()
	if response.Err = userVo.GetUserVo(user); response.Err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}

	// 5. 所有步骤都没有问题之后就可以将vo层处理好的数据返回了
	response.Code = constant.SUCCESS.Code
	response.Data = userVo
	response.Message = ""
}
