package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/lvdbing/bgo/internal/helper/limiter"
	"github.com/lvdbing/bgo/internal/model"
	"github.com/lvdbing/bgo/pkg/errcode"
)

func RateLimiter(l limiter.LimiterInterface) gin.HandlerFunc {
	return func(c *gin.Context) {
		key := l.Key(c)
		if bucket, ok := l.GetBucket(key); ok {
			count := bucket.TakeAvailable(1)
			if count == 0 {
				resp := model.NewResponse(c)
				resp.SendError(errcode.TooManyRequests)
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
