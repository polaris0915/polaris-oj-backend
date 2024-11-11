package user_service

import (
	// "polaris-oj-backend/database/mysql"

	"polaris-oj-backend/database/mysql"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Service struct {
	ctx *gin.Context
	db  *gorm.DB
}

func NewService(c *gin.Context) *Service {
	return &Service{
		ctx: c,
		db:  mysql.DB,
	}
}
