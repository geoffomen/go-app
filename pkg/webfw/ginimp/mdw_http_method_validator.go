package ginimp

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func httpMethodValidator() gin.HandlerFunc {
	return func(c *gin.Context) {
		methods, ok := pathToMethod[c.Request.URL.Path]
		if !ok {
			c.Error(fmt.Errorf("%s", stack(1)))
			c.Error(fmt.Errorf("path not found"))
			c.Status(404)
			c.Abort()
			return
		}
		found := false
		for _, method := range methods {
			switch method {
			case "", c.Request.Method:
				found = true
			}
			if found {
				break
			}
		}
		if !found {
			c.Error(fmt.Errorf("%s", stack(1)))
			c.Error(fmt.Errorf("method not allowed"))
			c.Status(405)
			c.Abort()
			return
		}
		c.Next()
	}
}
