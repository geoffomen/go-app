package echoargssrv

import (
	"example.com/internal/pkg/myconfig"
	"example.com/internal/pkg/mylog"
)

// Service ...
type Service struct {
	config         myconfig.MyConfigIface
	logger         mylog.MyLogIface
	useraccountSrv useraccountSrvIface
}

var instance *Service = &Service{}

// New ...
func New(
	config myconfig.MyConfigIface,
	logger mylog.MyLogIface,
	uSrv useraccountSrvIface,
) *Service {
	*instance = Service{
		config:         config,
		logger:         logger,
		useraccountSrv: uSrv,
	}
	return instance
}

// GetInstance ..
func GetInstance() *Service {
	return instance
}
