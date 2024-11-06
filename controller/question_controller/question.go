package question_controller

import (
	questionservice "polaris-oj-backend/service/question_service"
)

// 实例化 GQuestionController提供给全局
var GQuestionController = NewQuestionController(questionservice.GQuestionService)

// QuestionController 负责处理用户请求
type QuestionController struct {
	questionService *questionservice.QuestionService
}

// NewQuestionController 初始化 QuestionController
func NewQuestionController(quesionService *questionservice.QuestionService) *QuestionController {
	return &QuestionController{questionService: quesionService}
}
