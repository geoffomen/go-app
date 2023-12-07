package httpserverimp

import (
	"fmt"

	"example.com/internal/app/common/base/vo"
)

func authHandler(authHandler AuthIface) func(ctx *Ctx) error {
	return func(ctx *Ctx) error {
		if authHandler == nil {
			return fmt.Errorf("authHandler not set")
		}
		si, isValid := authHandler.Auth(ctx.request)
		if isValid {
			ctx.setSessionInfo(si)
		} else {
			ctx.setSessionInfo(&vo.SessionInfo{})
		}
		return ctx.Next()
	}
}
