package user_controller

import userservice "polaris-oj-backend/service/user_service"

// 实例化 GUserController提供给全局

var GUserController = NewUserController(userservice.GUserService)

// UserController 负责处理用户请求
type UserController struct {
	userService *userservice.UserService
}

// NewUserController 初始化 UserController
func NewUserController(userService *userservice.UserService) *UserController {
	return &UserController{userService: userService}
}
