package middleware

import (
	"polaris-oj-backend/common"

	"polaris-oj-backend/utils"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

// 日志中间件：当用户请求到后端，将用户的信息注入的gin和标准库的上下文中
func GetUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var userInfo *utils.Claims
		var err error
		if userInfo, err = common.GetLoginUser(session); err != nil {
			c.Next()
		}

		c.Set("user", userInfo)
		c.Next()
	}
}
