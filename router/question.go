package router

import (
	questioncontroller "polaris-oj-backend/controller/question_controller"
	"polaris-oj-backend/controller/questionsubmit_controller"

	"github.com/gin-gonic/gin"
)

// 定义用户所有的路由
func QuestionAdd(group *gin.RouterGroup) {
	group.POST("/add", questioncontroller.AddQuestion)
	group.POST("/update", questioncontroller.UpdateQuestion)
	group.POST("/delete", questioncontroller.DeleteQuestion)
	group.GET("/get", questioncontroller.GetQuestionById)
	group.POST("/list/page", questioncontroller.ListQuestionByPage)

	question_submit := group.Group("/question_submit")
	QuestionSubmitAdd(question_submit)
}

func QuestionSubmitAdd(group *gin.RouterGroup) {
	group.POST("/do", questionsubmit_controller.AddQuestionSubmit)
	group.POST("/list/page", questionsubmit_controller.ListQuestionSubmitByPage)
}
