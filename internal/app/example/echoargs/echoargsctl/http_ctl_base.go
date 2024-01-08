package echoargsctl

import (
	"ibingli.com/internal/app/example/echoargs/echoargssrv"
	"ibingli.com/internal/app/example/useraccount/useraccountsrv"
	"ibingli.com/internal/pkg/myConfig"
	"ibingli.com/internal/pkg/myDatabase"
	"ibingli.com/internal/pkg/myLog"
)

var (
	controllers map[string]interface{} = make(map[string]interface{})
	srv         *echoargssrv.Service
)

func New(config myConfig.Iface, logger myLog.Iface, db myDatabase.Iface) map[string]interface{} {
	// 初始化服务，并注入依赖
	srv = echoargssrv.New(
		config,
		logger,
		db,
		useraccountsrv.GetInstance(),
	)

	return controllers
}
