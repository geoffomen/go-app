package useraccctl

import (
	"github.com/storm-5/go-app/examples/user/userimp"
	"github.com/storm-5/go-app/examples/useracc"
	"github.com/storm-5/go-app/examples/useracc/useraccimp"
	"github.com/storm-5/go-app/pkg/config"
	"github.com/storm-5/go-app/pkg/database"
	"github.com/storm-5/go-app/pkg/webfw"
)

var (
	controllers map[string]interface{} = make(map[string]interface{})
	srv         *useraccimp.Service
)

// Controller ..
func Controller() map[string]interface{} {
	srv = useraccimp.New(
		database.GetClient(),
		config.GetInstance(),
		userimp.GetInstance(),
	)

	return controllers
}

func init() {
	m := map[string]interface{}{

		"POST /exam/v1/useracc/register": func(param useracc.CreateRequestDto) (interface{}, error) {
			return srv.Register(param)
		},

		"POST /exam/v1/useracc/login": func(param useracc.LoginRequestDto) (interface{}, error) {
			return srv.Login(param)
		},

		"/exam/v1/useracc/logout": func(sessData webfw.SessionInfo) (interface{}, error) {
			return srv.Logout(sessData)

		},
	}

	for p, h := range m {
		controllers[p] = h
	}
}
