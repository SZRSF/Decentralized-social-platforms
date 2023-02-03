package routes

import (
	"net/http"
	"zengzhicheng/Decentralized-social-platforms/controller"
	"zengzhicheng/Decentralized-social-platforms/logger"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册业务路由
	r.POST("/signup", controller.SignUpHandler)
	r.POST("/login", controller.LoginHandler)

	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
	return r
}
