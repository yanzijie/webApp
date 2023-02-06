package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

// RateLimitMiddleware 限流 - 令牌桶
func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {
	/*
		// 创建指定填充速率和容量大小的令牌桶
		// 以每fillInterval填充一个令牌的速率填充, 直到给定的最大容量 capacity
		func NewBucket(fillInterval time.Duration, capacity int64) *Bucket
		// 创建指定填充速率、容量大小和每次填充的令牌数的令牌桶
		func NewBucketWithQuantum(fillInterval time.Duration, capacity, quantum int64) *Bucket
		// 创建填充速度为指定速率和容量大小的令牌桶
		// NewBucketWithRate(0.1, 200) 表示每秒填充20个令牌
		func NewBucketWithRate(rate float64, capacity int64) *Bucket
	*/

	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		// 如果取不到令牌就中断本次请求返回 rate limit...
		if bucket.TakeAvailable(1) < 1 {
			c.String(http.StatusOK, "rate limit...")
			c.Abort()
			return
		}
		// 取到令牌, 过
		c.Next()
	}
}
