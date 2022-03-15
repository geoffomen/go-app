package userctl

import (
	"github.com/geoffomen/go-app/examples/user"
	"github.com/geoffomen/go-app/examples/user/userimp"
	"github.com/geoffomen/go-app/pkg/database"
)

var (
	controllers map[string]interface{} = make(map[string]interface{})
	srv         *userimp.Service
)

// Controller ..
func Controller() map[string]interface{} {
	srv = userimp.New(
		database.GetClient(),
	)
	return controllers
}

func init() {
	m := map[string]interface{}{

		"GET /exam/v1/user/info": func(param user.GetUserInfoRequestDto) (interface{}, error) {
			return srv.GetUserInfo(param)
		},

		"GET /exam/v1/user/page": func(param user.PageRequestDto) (interface{}, error) {
			return srv.Page(param)
		},
	}

	for p, h := range m {
		controllers[p] = h
	}
}
