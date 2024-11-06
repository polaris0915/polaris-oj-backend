package questionsubmit_service

import (
	"gorm.io/gorm"
	"polaris-oj-backend/database/mysql"
)

// 实例化 QuestionSubmitService提供给全局
var GQuestionSubmitService = NewQuestionSubmitService(mysql.DB)

// QuestionSubmitService 定义用户相关的服务
type QuestionSubmitService struct {
	db *gorm.DB
}

// NewQuestion 初始化 QuestionSubmitService
func NewQuestionSubmitService(db *gorm.DB) *QuestionSubmitService {
	return &QuestionSubmitService{db: db}
}
