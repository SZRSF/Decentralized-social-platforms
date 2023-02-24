package routes

import (
	"net/http"
	"zengzhicheng/Decentralized-social-platforms/controller"
	"zengzhicheng/Decentralized-social-platforms/logger"
	"zengzhicheng/Decentralized-social-platforms/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	//if mode == gin.ReleaseMode {
	//	gin.SetMode(gin.ReleaseMode)	// gin设置发布模式
	//}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	v1 := r.Group("/api/v1")

	// 注册
	v1.POST("/signup", controller.SignUpHandler)
	// 登录
	v1.POST("/login", controller.LoginHandler)

	v1.Use(middlewares.JWTAuthMiddleware()) // 应用JWT认证中间件

	{
		// 获取用户信息
		v1.GET("/user/:user_id", controller.UserDetailHandler)

		// 获取家信息
		v1.GET("/family", controller.FamilyHandler)
		// 根据id获取家信息
		v1.GET("/family/:id", controller.FamilyDetailHandler)

		// 发布文章
		v1.POST("/post", controller.CreatePostHandler)
		// 根据id返回文章信息
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		// 获取作品详细信息
		v1.GET("/posts", controller.GetPostListHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
