package httpserverimp

import "example.com/internal/app/common/base/vo"

func authHandler() func(ctx *Ctx) error {
	return func(ctx *Ctx) error {
		ctx.setSessionInfo(&vo.SessionInfo{})
		return ctx.Next()
	}
}
