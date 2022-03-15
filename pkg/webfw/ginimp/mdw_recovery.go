package ginimp

import (
	"reflect"

	"github.com/gin-gonic/gin"
)

func recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				path := c.Request.URL.Path
				method := c.Request.Method
				args, exists := c.Get("args")
				if exists {
					ags := args.([]reflect.Value)
					log.Errorf("method: %s, path: %s, msg: %s, args: %s, stack: %s",
						method,
						path,
						err,
						reflectValueToString(ags),
						stack(3),
					)
				} else {
					log.Errorf("method: %s, path: %s, msg: %s, stack: %s",
						method,
						path,
						err,
						stack(3),
					)
				}
				c.AbortWithStatusJSON(500, "server error， retry later please")
			}
		}()
		c.Next() // execute all the handlers
	}
}
