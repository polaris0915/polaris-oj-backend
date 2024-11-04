package router

import (
	questioncontroller "polaris-oj-backend/controller/question_controller"

	"github.com/gin-gonic/gin"
)

// 定义用户所有的路由
func QuestionAdd(group *gin.RouterGroup) {
	group.POST("/add", questioncontroller.GQuestionController.AddQuestion)
	group.POST("/update", questioncontroller.GQuestionController.UpdateQuestion)
	group.POST("/delete", questioncontroller.GQuestionController.DeleteQuestion)
	group.GET("/get", questioncontroller.GQuestionController.GetQuestionById)
}
