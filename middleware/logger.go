package middleware

import (
	"context"

	"polaris-oj-backend/common"
	"polaris-oj-backend/common/polaris_logger"
	"polaris-oj-backend/config"
	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 日志中间件：当用户请求到后端，将用户的信息注入的gin和标准库的上下文中
func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		var userInfo *utils.Claims
		var err error

		enrichedLogger := polaris_logger.Logger.With(
			zap.String("ip", c.ClientIP()),
		)

		if userInfo, err = common.GetLoginUser(session); err == nil {
			enrichedLogger = enrichedLogger.With(
				zap.String("userIdentity", userInfo.Identity),
				zap.String("userAccount", userInfo.UserAccount),
			)
		} else {
			// 如果无法获取用户信息，记录警告日志
			enrichedLogger.Warn("Failed to retrieve user info", zap.Error(err))
		}
		ctx := context.WithValue(c.Request.Context(), config.Log.LogContextKey, userInfo)
		c.Request = c.Request.WithContext(ctx)

		c.Set("logger", enrichedLogger)
		c.Next()
	}
}
