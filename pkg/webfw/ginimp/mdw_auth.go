package ginimp

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/geoffomen/go-app/pkg/webfw"
	"github.com/gin-gonic/gin"
)

func authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		needAuth := true
		for _, p := range conf.GetStringSliceOrDefault("server.noNeedAuthPath", []string{}) {
			if strings.HasPrefix(c.Request.URL.Path, p) {
				needAuth = false
				break
			}
		}

		sessionInfo := &webfw.SessionInfo{}
		if needAuth {
			// do auth check
			authStr := c.GetHeader("Authorization")
			if authStr == "" {
				c.Error(fmt.Errorf("%s", stack(1)))
				c.Error(fmt.Errorf("请登录"))
				c.Status(http.StatusUnauthorized)
				c.Abort()
				return
			}
			s := strings.Split(authStr, " ")
			if len(s) < 2 {
				c.Error(fmt.Errorf("%s", stack(1)))
				c.Error(fmt.Errorf("请求头无效，格式: Bearer <token>"))
				c.Status(http.StatusUnauthorized)
				c.Abort()
				return
			}
			if s[0] != "Bearer" {
				c.Error(fmt.Errorf("%s", stack(1)))
				c.Error(fmt.Errorf("令牌类型错误"))
				c.Status(http.StatusUnauthorized)
				c.Abort()
				return
			}
		}

		c.Set("sessionInfo", *sessionInfo)
		c.Next()
	}
}
