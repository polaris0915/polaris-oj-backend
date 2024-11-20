package router

import (
	usercontroller "polaris-oj-backend/controller/user_controller"

	"github.com/gin-gonic/gin"
)

// 定义用户所有的路由
func UserAdd(group *gin.RouterGroup) {
	group.POST("/login", usercontroller.Login)
	group.POST("/register", usercontroller.AddUser)
	group.POST("/my", usercontroller.UpdateMyUser)
	group.POST("/logout", usercontroller.UserLogout)
	group.POST("/delete", usercontroller.DeleteUser)
	group.POST("/update-avatar", usercontroller.UpdateAvatar)
}
