package useraccountctl

import (
	"database/sql"

	"example.com/internal/app/common/base/vo"
	"example.com/internal/app/example/useraccount/useraccountrepo"
	"example.com/internal/app/example/useraccount/useraccountsrv"
	"example.com/internal/pkg/myconfig"
	"example.com/internal/pkg/mylog"
)

func New(config myconfig.MyConfigIface, logger mylog.MyLogIface, db *sql.DB) map[string]interface{} {
	srv := useraccountsrv.New(
		config,
		logger,
		useraccountrepo.New(db, logger),
	)
	return map[string]interface{}{
		"POST /example/api/v1/useraccount/register": func(ctx vo.SessionInfo, args useraccountsrv.CreateRequestDto) (interface{}, error) {
			return srv.Register(ctx, args)
		},
	}
}
