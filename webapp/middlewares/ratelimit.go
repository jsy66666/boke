package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
	"net/http"
	"time"
)

func RateLimitMiddleware(fillInterval time.Duration, cap int64) func(c *gin.Context) {

	bucket := ratelimit.NewBucket(fillInterval, cap)
	return func(c *gin.Context) {
		if bucket.TakeAvailable(1) == 0 {
			c.String(http.StatusOK, "rate limit failed")
			c.Abort()
			return
		}
		c.Next()
	}
}
