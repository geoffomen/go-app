package accountctl

import (
	"github.com/geoffomen/go-app/examples/account"
	"github.com/geoffomen/go-app/examples/account/accountimp"
	"github.com/geoffomen/go-app/examples/user/userimp"
	"github.com/geoffomen/go-app/pkg/config"
	"github.com/geoffomen/go-app/pkg/database"
	"github.com/geoffomen/go-app/pkg/vo"
)

// Controller ..
func Controller() map[string]interface{} {
	srv := accountimp.New(
		database.GetClient(),
		config.GetInstance(),
		userimp.GetInstance(),
	)

	return map[string]interface{}{

		"POST /exam/v1/account/register": func(param account.CreateRequestDto) (interface{}, error) {
			return srv.Register(param)
		},

		"POST /exam/v1/account/login": func(param account.LoginRequestDto) (interface{}, error) {
			return srv.Login(param)
		},

		"/exam/v1/account/logout": func(sessData vo.SessionInfo) (interface{}, error) {
			return srv.Logout(sessData)

		},
	}
}
