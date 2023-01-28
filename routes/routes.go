package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yanzijie/webApp/logger"
	"net/http"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 注册路由
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "ok!ok!")
	})

	return r
}
