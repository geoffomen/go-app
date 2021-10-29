package userctl

import (
	"github.com/geoffomen/go-app/examples/user"
	"github.com/geoffomen/go-app/examples/user/userimp"
	"github.com/geoffomen/go-app/pkg/database"
)

// Controller ..
func Controller() map[string]interface{} {
	srv := userimp.New(
		database.GetClient(),
	)

	return map[string]interface{}{

		"GET /exam/v1/user/info": func(param user.GetUserInfoRequestDto) (interface{}, error) {
			return srv.GetUserInfo(param)
		},

		"GET /exam/v1/user/page": func(param user.PageRequestDto) (interface{}, error) {
			return srv.Page(param)
		},
	}
}
