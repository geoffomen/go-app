package useraccountctl

import (
	"ibingli.com/internal/app/example/echoargs/echoargssrv"
	"ibingli.com/internal/app/example/useraccount/useraccountsrv"
	"ibingli.com/internal/pkg/myConfig"
	"ibingli.com/internal/pkg/myDatabase"
	"ibingli.com/internal/pkg/myLog"
)

var (
	controllers map[string]interface{} = make(map[string]interface{})
	srv         *useraccountsrv.Service
)

func New(config myConfig.Iface, logger myLog.Iface, db myDatabase.Iface) map[string]interface{} {
	srv = useraccountsrv.New(
		config,
		logger,
		db,
		echoargssrv.GetInstance(),
	)

	return controllers
}
