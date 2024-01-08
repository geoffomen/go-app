package myHttpServerImp

import (
	"context"
	"fmt"

	"ibingli.com/internal/pkg/myHttpServer"
)

func authHandler(authHandler authIface) func(ctx *Ctx) error {
	return func(ctx *Ctx) error {
		if authHandler == nil {
			return fmt.Errorf("authHandler not set")
		}
		si, isValid := authHandler.Auth(ctx.request)
		if isValid {
			ctx.setSessionInfo(si)
		} else {
			ctx.setSessionInfo(&myHttpServer.SessionInfo{Ctx: context.Background()})
		}
		return ctx.Next()
	}
}
