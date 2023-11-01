package echoargssrv

import (
	"example.com/internal/pkg/myconfig"
	"example.com/internal/pkg/mylog"
)

// Service ...
type Service struct {
	config myconfig.MyConfigIface
	logger mylog.MyLogIface
}

var instance *Service = &Service{}

// New ...
func New(
	config myconfig.MyConfigIface,
	logger mylog.MyLogIface,
) *Service {
	*instance = Service{
		config: config,
		logger: logger,
	}
	return instance
}

// GetInstance ..
func GetInstance() *Service {
	return instance
}
