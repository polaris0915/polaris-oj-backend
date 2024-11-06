package questionsubmit_controller

import (
	questionsubmitservice "polaris-oj-backend/service/questionsubmit_service"
)

// 实例化 GQuestionSubmitController提供给全局
var GQuestionSubmitController = NewQuestionSubmitController(questionsubmitservice.GQuestionSubmitService)

// QuestionSubmitController 负责处理用户请求
type QuestionSubmitController struct {
	questionSubmitService *questionsubmitservice.QuestionSubmitService
}

// NewQuestionSubmitController 初始化 QuestionSubmitController
func NewQuestionSubmitController(quesionService *questionsubmitservice.QuestionSubmitService) *QuestionSubmitController {
	return &QuestionSubmitController{questionSubmitService: quesionService}
}
