package controller

import (
	"errors"
	"net/http"
	"polaris-oj-backend/common"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/models/dto/user_dto"
	"polaris-oj-backend/models/vo"
	uservo "polaris-oj-backend/models/vo/user_vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"polaris-oj-backend/service"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 实例化 GUserController提供给全局
var GUserController = NewUserController(service.GUserService)

// UserController 负责处理用户请求
type UserController struct {
	userService *service.UserService
}

// NewUserController 初始化 UserController
func NewUserController(userService *service.UserService) *UserController {
	return &UserController{userService: userService}
}

// Register
// @Tags 公共方法, 用户
// @Summary 用户注册
// @Param login body user_dto.UserRegisterRequest true "user login infos"
// @Success 200 {object} vo.BaseResponse "{Code:"0",Data:{...}, Message:""}"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/user/register [post]
func (uc *UserController) Register(c *gin.Context) {
	// 00. 函数结束固定调用NewSubResponse
	// response对象初始化遵循Code以及Message都是未决的状态
	// 方便之后判断，那么data接口的数据应该也是为nil的状态
	response := vo.NewSubResponse(c, http.StatusOK, constant.UNDEFINED.Code, nil, constant.UNDEFINED.Message)
	defer func() { // 使用闭包
		response.Response()
	}()
	// 1. 现将请求中的数据转换解析道对应的请求模型中
	userRegisterRequest := user_dto.NewUserRegisterRequest()
	if response.Err = c.ShouldBindJSON(userRegisterRequest); response.Err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================
	if userRegisterRequest.CheckPassword != userRegisterRequest.UserPassword {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = errors.New("两次输入的密码不一致").Error()
		return
	}
	// ================controller特殊的业务需求===================
	// 2. 通过DTO层将请求模型转化到数据表模型中
	user := new(allModels.User)
	if response.Err = userRegisterRequest.DtoToModel(user); response.Err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// 3. 将DTO层处理好的数据表模型的数据传入service中进行具体的逻辑操作
	// service层要是没有任务问题，那么返回的error也是为空
	// TODO unfinished: token看看能不能优化成结构体对象
	response.Err = uc.userService.RegisterUser(user)
	if response.Err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================

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

// GetLoginUser
// @Tags 公共方法, 用户
// @Summary 获取当前登录用户
// @Success 200 {object} vo.BaseResponse "{Code:"0",Data:{...}, Message:""}"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/get/user/login [get]
func (uc *UserController) GetLoginUser(c *gin.Context) {
	// 00. 函数结束固定调用NewSubResponse
	// response对象初始化遵循Code以及Message都是未决的状态
	// 方便之后判断，那么data接口的数据应该也是为nil的状态
	response := vo.NewSubResponse(c, http.StatusOK, constant.UNDEFINED.Code, nil, constant.UNDEFINED.Message)
	defer func() { // 使用闭包
		response.Response()
	}()
	// 1. 现将请求中的数据转换解析道对应的请求模型中

	// ================controller特殊的业务需求===================
	// 先尝试获取请求头中的token，没有的话就是未登录
	_, response.Err = c.Cookie(constant.USER_LOGIN_STATE)
	if _, response.Err = c.Cookie(constant.USER_LOGIN_STATE); response.Err != nil {
		response.Code = constant.NOT_LOGIN_ERROR.Code
		response.Data = nil
		response.Message = errors.New("未登录").Error()
		return
	}
	// ================controller特殊的业务需求===================

	// 2. 通过DTO层将请求模型转化到数据表模型中

	// 3. 将DTO层处理好的数据表模型的数据传入service中进行具体的逻辑操作
	// service层要是没有任务问题，那么返回的error也是为空
	session := sessions.Default(c)
	user := new(allModels.User)
	if response.Err = uc.userService.GetLoginUser(session, user); response.Err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================

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

// UpdateMyUser
// @Tags 私有方法, 用户
// @Summary 用户更新自己的数据
// @Param updateInfo body user_dto.UserUpdateMyUserRequest true "user login infos"
// @Success 200 {object} vo.BaseResponse "{Code:"0",Data:{...}, Message:""}"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/user/my [post]
func (uc *UserController) UpdateMyUser(c *gin.Context) {
	// 00. 函数结束固定调用NewSubResponse
	// response对象初始化遵循Code以及Message都是未决的状态
	// 方便之后判断，那么data接口的数据应该也是为nil的状态
	response := vo.NewSubResponse(c, http.StatusOK, constant.UNDEFINED.Code, nil, constant.UNDEFINED.Message)
	defer func() { // 使用闭包
		response.Response()
	}()
	// 1. 现将请求中的数据转换解析道对应的请求模型中
	userUpdateMyUserRequest := user_dto.NewUserUpdateMyUserRequest()
	if response.Err = c.ShouldBindJSON(userUpdateMyUserRequest); response.Err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 2. 通过DTO层将请求模型转化到数据表模型中
	user := new(allModels.User)
	if response.Err = userUpdateMyUserRequest.DtoToModel(user); response.Err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// 3. 将DTO层处理好的数据表模型的数据传入service中进行具体的逻辑操作
	// service层要是没有任务问题，那么返回的error也是为空
	session := sessions.Default(c)
	if response.Err = uc.userService.UpdateMyUser(session, user); response.Err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================

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

	// test for git
	response.Code = constant.SUCCESS.Code
	response.Data = true
	response.Message = ""
}
