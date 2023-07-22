package router

import (
	"bluebell/controller"
	"bluebell/logger"
	middlewares "bluebell/middwares"
	"bluebell/setting"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SetUp() *gin.Engine {
	// 设置 gin 框架日志输出模式
	gin.SetMode(setting.Conf.GinConfig.Mode)

	//创建一个路由引擎
	r := gin.Default()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	v1 := r.Group("/api/v1")
	v1.POST("/signup", controller.SignUpHandler)
	v1.POST("/login", controller.LoginHandler)
	{
		v1.GET("/community", controller.CommunityHandler)
	}

	// 应用 JWT 认证中间件以及令牌桶限流中间件
	v1.Use(middlewares.JWTAuthMiddleware(), middlewares.RateLimitMiddleware(2*time.Second, 1))
	{
		v1.POST("/post", controller.CreatePostHandler)
		v1.GET("/post/:id", controller.GetPostByIDHandler)
		v1.GET("/posts/", controller.GetPostListHandler)
		v1.POST("/vote", controller.PostVoteHandler)
	}
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})

	return r
}
