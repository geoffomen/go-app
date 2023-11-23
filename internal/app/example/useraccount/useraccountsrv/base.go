package useraccountsrv

import (
	"example.com/internal/pkg/myconfig"
	"example.com/internal/pkg/mylog"
)

// Service ...
type Service struct {
	config      myconfig.MyConfigIface
	logger      mylog.MyLogIface
	repo        useraccountRepo
	echoargssrv ecoArgsSrvIface
}

var instance *Service = &Service{}

// New ...
func New(
	config myconfig.MyConfigIface,
	logger mylog.MyLogIface,
	repo useraccountRepo,
	echoargssrv ecoArgsSrvIface,
) *Service {
	*instance = Service{
		config:      config,
		logger:      logger,
		repo:        repo,
		echoargssrv: echoargssrv,
	}
	return instance
}

// GetInstance ..
func GetInstance() *Service {
	return instance
}
