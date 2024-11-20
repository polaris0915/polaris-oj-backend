package router

import (
	"polaris-oj-backend/controller/file_controller"

	"github.com/gin-gonic/gin"
)

// 定义管理员所有的路由
func FileAdd(group *gin.RouterGroup) {
	group.Static("/download", "./upload/")
	group.POST("/upload", file_controller.FileUpload)
}
