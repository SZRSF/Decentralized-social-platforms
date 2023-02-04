package routes

import (
	"net/http"
	"zengzhicheng/Decentralized-social-platforms/controller"
	"zengzhicheng/Decentralized-social-platforms/logger"
	"zengzhicheng/Decentralized-social-platforms/middlewares"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	r.GET("/ping", middlewares.JWTAuthMiddleware(), func(c *gin.Context) {
		// 如果是登录的用户, 判断请求头中是否有 有效的JWT
		isLogin := true
		c.Request.Header.Get("Authorization")
		if isLogin {
			c.String(http.StatusOK, "pong")
		} else {
			// 否则直接返回登录
			c.String(http.StatusOK, "请登录")
		}

	})
	return r
}
