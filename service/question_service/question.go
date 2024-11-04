package questionservice

import (
	"gorm.io/gorm"
	"polaris-oj-backend/database/mysql"
)

// 实例化 QuestionService提供给全局
var GQuestionService = NewQuestionService(mysql.DB)

// QuestionService 定义用户相关的服务
type QuestionService struct {
	db *gorm.DB
}

// NewQuestion 初始化 QuestionService
func NewQuestionService(db *gorm.DB) *QuestionService {
	return &QuestionService{db: db}
}
