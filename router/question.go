package router

import (
	"github.com/gin-gonic/gin"
	"polaris-oj-backend/controller"
)

// 定义用户所有的路由
func QuestionAdd(group *gin.RouterGroup) {
	group.POST("/add", controller.GQuestionController.AddQuestion)
	group.POST("/update", controller.GQuestionController.UpdateQuestion)
	group.POST("/delete", controller.GQuestionController.DeleteQuestion)
	group.GET("/get", controller.GQuestionController.GetQuestionById)
}
