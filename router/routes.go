package router

import (
	"polaris-oj-backend/config"
	usercontroller "polaris-oj-backend/controller/user_controller"
	"polaris-oj-backend/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"

	_ "polaris-oj-backend/docs"

	"polaris-oj-backend/models/enums/userrole_enum"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	// swagger配置
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// session配置
	// 设置 session 存储引擎，这里使用 cookie 存储
	// TODO config: 通过下面代码设置session失效期，如果不指定 MaxAge，默认行为通常是 session 仅在会话期间有效，即关闭浏览器后 session 会失效。
	store := cookie.NewStore([]byte(config.Session.SESSION_PAIRKEY))
	// store.Options(sessions.Options{
	// 	MaxAge: int(constant.ValidTime) * 60, // 设置为 3600 秒（即 1 小时）
	// })
	r.Use(sessions.Sessions(userrole_enum.USER_LOGIN_STATE, store))

	r.Use(middleware.GetUserMiddleware())
	// 跨域中间件配置以及日志中间件
	r.Use(middleware.Cors(), middleware.LogMiddleware())

	// 路由规则
	//  -----------------------公共组的api----------------------
	api := r.Group("/api")
	// get组
	get := api.Group("/get")
	get.GET("/user/login", usercontroller.GetLoginUser)
	// -----------------------私有组的api----------------------
	// 用户组
	user := api.Group("/user")
	UserAdd(user)

	// 问题组
	question := api.Group("/question")
	QuestionAdd(question)

	// 管理员组
	admin := r.Group("/admin")
	AdminAdd(admin)

	// 管理员组
	file := api.Group("/file")
	FileAdd(file)

	return r
}
