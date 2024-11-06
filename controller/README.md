#### controller模板

```go
// Login【接口函数名字】
// @Tags 公共方法, 用户【分类tags】
// @Summary 用户登陆【描述】
// @Param login body user_dto.UserLoginRequest true "user login infos"【讲user_dto.UserLoginRequest更换成新接口的dto视图】
// @Success 200 {object} vo.BaseResponse "{Code:"0",Data:{...}, Message:""}"【可以不变】
// @Failure 401 {string} string "Unauthorized"【可以不变】
// @Failure 403 {string} string "Forbidden"【可以不变】
// @Failure 404 {string} string "Not Found"【可以不变】
// @Router /api/user/login [post]【api路径以及api提交方法】
func (uc *UserController) Login(c *gin.Context) {
  【UserController更换所属结构体，Login更换成新api名字】
	// 00. 函数结束固定调用NewSubResponse
	// response对象初始化遵循Code以及Message都是未决的状态
	// 方便之后判断，那么data接口的数据应该也是为nil的状态
	response := vo.NewSubResponse(c, http.StatusOK, constant.UNDEFINED.Code, nil, constant.UNDEFINED.Message)
	defer func() { // 使用闭包
		response.Response()
	}()
  【00基本不用变】
	// 1. 绑定请求数据到DTO层模型中
  【更换dto层视图模型】
	userLoginRequest := user_dto.NewUserLoginRequest()
	if response.Err = c.ShouldBindJSON(userLoginRequest); response.Err != nil {
		// 一般要是发生错误基本上都是将response的Err更新成返回的error
		// Code根据具体错误的情况再更新
		response.Code = constant.PARAMS_ERROR.Code
		// 发生错误就不要返回数据
		response.Data = nil
		// repsonse中的Message更新成response.Err.Error()
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================
	【如果有需要写进入dto层之前的一些简单逻辑】
	// ================controller特殊的业务需求===================
	// 2. 将DTO层数据与数据表响应字段的数据绑定到数据表模型中
  【更换数据库模型】
	user := new(allModels.User)
  【更换dto层的dto视图转数据库模型的方法】
	if response.Err = userLoginRequest.DtoToModel(user); response.Err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// 3. 将DTO层处理好的数据表模型的数据传入service中进行具体的逻辑操作
	// service层要是没有任务问题，那么返回的error也是为空
	// TODO: token看看能不能优化成结构体对象
	var token string【传入service的数据一般固定就是数据库模型】
  								【如果需要带出来数据，就通过返回值的办法操作】
	token, response.Err = uc.userService.LoginUser(user)
	if response.Err != nil {
		// 暂时统一将service导出的错误定义为系统错误
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================
	【如果有需要在出service层写一些简单逻辑，例如这里的设置cookie】
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
	userVo := uservo.NewUserVo()【更换VO模型】
	if response.Err = userVo.GetUserVo(user); response.Err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	【最终返回基本可以不变】
	// 5. 所有步骤都没有问题之后就可以将vo层处理好的数据返回了
	response.Code = constant.SUCCESS.Code
	response.Data = userVo
	response.Message = ""
}
```

```go
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
	// 1. 绑定请求数据到DTO层模型中
	userLoginRequest := user_dto.NewUserLoginRequest()
	if response.Err = c.ShouldBindJSON(userLoginRequest); response.Err != nil {
		// 一般要是发生错误基本上都是将response的Err更新成返回的error
		// Code根据具体错误的情况再更新
		response.Code = constant.PARAMS_ERROR.Code
		// 发生错误就不要返回数据
		response.Data = nil
		// repsonse中的Message更新成response.Err.Error()
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================
	// ================controller特殊的业务需求===================
	// 2. 将DTO层数据与数据表响应字段的数据绑定到数据表模型中
	user := new(allModels.User)
	if response.Err = userLoginRequest.DtoToModel(user); response.Err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// 3. 将DTO层处理好的数据表模型的数据传入service中进行具体的逻辑操作
	// service层要是没有任务问题，那么返回的error也是为空
	// TODO: token看看能不能优化成结构体对象
	var token string
	token, response.Err = uc.userService.LoginUser(user)
	if response.Err != nil {
		// 暂时统一将service导出的错误定义为系统错误
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
```

