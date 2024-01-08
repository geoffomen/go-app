package useraccountctl

import (
	"ibingli.com/internal/app/example/useraccount/useraccountdm"
	"ibingli.com/internal/pkg/myHttpServer"
)

func init() {
	m := map[string]interface{}{

		"POST /example/api/v1/useraccount/register": func(ctx *myHttpServer.SessionInfo, args *useraccountdm.CreateRequestDto) (interface{}, error) {
			return srv.Register(ctx, args)
		},
	}

	for p, h := range m {
		controllers[p] = h
	}
}
