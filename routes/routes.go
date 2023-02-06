package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yanzijie/webApp/logger"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	// 使用中间件
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	// 设置限流中间件,每2秒放一个令牌，桶的大小也只有1，所以每2秒只能请求一次
	// r.Use(middleware.RateLimitMiddleware(2*time.Second, 1))

	// 注册路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok!ok!")
	})

	return r
}
