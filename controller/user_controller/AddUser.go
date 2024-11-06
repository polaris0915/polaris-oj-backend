package user_controller

import (
	"net/http"
	"polaris-oj-backend/constant"
	"polaris-oj-backend/models/dto"
	"polaris-oj-backend/models/dto/user_dto"
	"polaris-oj-backend/models/vo"
	"polaris-oj-backend/polaris_oj_backend/allModels"

	"github.com/gin-gonic/gin"
)

// AddUser
// @Tags 公共方法, 用户
// @Summary 用户注册
// @Param login body user_dto.UserAddRequest true "user Register infos"
// @Success 200 {object} vo.BaseResponse "{Code:"0",Data:{...}, Message:""}"
// @Failure 401 {string} string "Unauthorized"
// @Failure 403 {string} string "Forbidden"
// @Failure 404 {string} string "Not Found"
// @Router /api/user/register [post]
func (uc *UserController) AddUser(c *gin.Context) {
	// 00. 函数结束固定调用NewSubResponse
	// response对象初始化遵循Code以及Message都是未决的状态
	// 方便之后判断，那么data接口的数据应该也是为nil的状态
	response := vo.NewSubResponse(c, http.StatusOK, constant.UNDEFINED.Code, nil, constant.UNDEFINED.Message)
	defer func() { // 使用闭包
		response.Response()
	}()
	// 1. 绑定请求数据到DTO层模型中
	var requestDto dto.RequestDto[*allModels.User] = new(user_dto.UserAddRequest)
	/*
		添加新的接口后，要去dto.BindAndValidateRequest添加新的断言
		以便请求数据能够成功转换到对应的请求模型中
	*/
	if response.Err = dto.BindAndValidateRequest(c, requestDto); response.Err != nil {
		response.Code = constant.PARAMS_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
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
	// TODO unfinished: token看看能不能优化成结构体对象
	// 创建用户接口是不用session的
	// session := sessions.Default(c)
	user := new(allModels.User)
	response.Err = uc.userService.AddUser(nil, requestDto, user)
	if response.Err != nil {
		response.Code = constant.SYSTEM_ERROR.Code
		response.Data = nil
		response.Message = response.Err.Error()
		return
	}
	// ================controller特殊的业务需求===================

	// ================controller特殊的业务需求===================
	// 3. 最终所有的业务逻辑也进行完毕之后，将返回的数据表模型数据交给VO层进行脱敏等操作

	// 4. 所有步骤都没有问题之后就可以将vo层处理好的数据返回了
	response.Code = constant.SUCCESS.Code
	response.Data = true
	response.Message = ""
}
