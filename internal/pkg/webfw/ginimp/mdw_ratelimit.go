package ginimp

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juju/ratelimit"
)

func rateLimit() gin.HandlerFunc {
	bk := ratelimit.NewBucket(100, 500)
	return func(c *gin.Context) {
		rt := bk.TakeAvailable(1)
		if rt == 0 {
			c.Error(fmt.Errorf("%s", stack(1)))
			c.Status(http.StatusTooManyRequests)
			c.Abort()
			return
		}
		c.Next()
	}
}
