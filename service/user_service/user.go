package userservice

import (
	"polaris-oj-backend/database/mysql"

	"gorm.io/gorm"
)

// 掌管古希腊的神啊～～～～

// 实例化 UserService提供给全局
var GUserService = NewUserService(mysql.DB)

// UserService 定义用户相关的服务
type UserService struct {
	db *gorm.DB
}

// NewUserService 初始化 UserService
func NewUserService(db *gorm.DB) *UserService {
	return &UserService{db: db}
}
