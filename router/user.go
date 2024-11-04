package router

import (
	"github.com/gin-gonic/gin"
	"polaris-oj-backend/controller"
)

// 定义用户所有的路由
func UserAdd(group *gin.RouterGroup) {
	group.POST("/login", controller.GUserController.Login)
	group.POST("/register", controller.GUserController.Register)
	group.POST("/my", controller.GUserController.UpdateMyUser)
	group.POST("/logout", controller.GUserController.UserLogout)
}
