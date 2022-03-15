package ginimp

import (
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Starting time
		start := time.Now()
		// process request
		c.Next()
		// End Time
		end := time.Now()
		//execution time
		cost := end.Sub(start)
		clientIP := c.ClientIP()
		path := c.Request.URL.Path
		method := c.Request.Method
		statusCode := c.Writer.Status()
		args, exists := c.Get("args")
		if exists {
			ags := args.([]reflect.Value)
			log.Infof("status: %3d, cost: %13v, clientIp: %15s, method: %s, path: %s, args: %s",
				statusCode,
				cost,
				clientIP,
				method,
				path,
				reflectValueToString(ags),
			)
		} else {
			log.Infof("status: %3d, cost: %13v, clientIp: %15s, method: %s, path: %s",
				statusCode,
				cost,
				clientIP,
				method,
				path,
			)
		}
	}
}
