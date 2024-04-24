package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
	"work_gin/utils"
)

func RateLimit() gin.HandlerFunc {
	// 创建令牌桶
	var times = time.Second * time.Duration(utils.Rate)
	bucket := ratelimit.NewBucket(times, utils.Capacity)
	return func(c *gin.Context) {
		// 判断
		if bucket.TakeAvailable(1) < 1 {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"code": http.StatusTooManyRequests, "message": "too many requests"})
			return
		}
		c.Next()
	}
}
