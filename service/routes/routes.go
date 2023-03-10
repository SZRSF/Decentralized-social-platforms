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
		// 关注用户
		v1.POST("/user/followings", controller.AddFollowHandler)
		// 取消关注用户
		v1.DELETE("/user/followings/:target", controller.DeleteFollowHandler)

		// 获取家信息
		v1.GET("/family", controller.FamilyHandler)
		// 根据家id获取家信息
		v1.GET("/family/:id", controller.FamilyDetailHandler)
		// 根据用户id获取用户家信息
		v1.GET("/family/myFamily/:user_id", controller.MyFamilyHandler)

		// 发布文章
		v1.POST("/post/image", controller.PostImageHandler)
		v1.POST("/post", controller.CreatePostHandler)
		// 根据id返回文章信息
		v1.GET("/post/:id", controller.GetPostDetailHandler)
		// 获取作品信息列表
		v1.GET("/posts/", controller.GetPostListHandler)
		// 收藏作品
		v1.POST("/works/collections", controller.AddCollectHandler)
		// 取消收藏作品
		v1.DELETE("/works/collections/:worksId", controller.DeleteCollectHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
