package middlewares

import (
	res "LDDP/utils/response"
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"time"
)

// RateLimitMiddleware 限流中间件
func RateLimitMiddleware(fillInterval time.Duration, cap, quantum int64) gin.HandlerFunc {
	bucket := ratelimit.NewBucketWithQuantum(fillInterval, cap, quantum)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) < 1 {
			res.ResError(c, res.CodeServerBusy)
			c.Abort()
			return
		}
		c.Next()
	}
}
