package router

import (
	usercontroller "polaris-oj-backend/controller/user_controller"

	"github.com/gin-gonic/gin"
)

// 定义用户所有的路由
func UserAdd(group *gin.RouterGroup) {
	group.POST("/login", usercontroller.GUserController.Login)
	group.POST("/register", usercontroller.GUserController.AddUser)
	group.POST("/my", usercontroller.GUserController.UpdateMyUser)
	group.POST("/logout", usercontroller.GUserController.UserLogout)
	group.POST("/delete", usercontroller.GUserController.DeleteUser)
}
