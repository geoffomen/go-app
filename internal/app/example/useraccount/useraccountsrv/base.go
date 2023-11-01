package useraccountsrv

import (
	"example.com/internal/pkg/myconfig"
	"example.com/internal/pkg/mylog"
)

// Service ...
type Service struct {
	config myconfig.MyConfigIface
	logger mylog.MyLogIface
	repo   useraccountRepo
}

var instance *Service = &Service{}

// New ...
func New(
	config myconfig.MyConfigIface,
	logger mylog.MyLogIface,
	repo useraccountRepo,
) *Service {
	*instance = Service{
		config: config,
		logger: logger,
		repo:   repo,
	}
	return instance
}

// GetInstance ..
func GetInstance() *Service {
	return instance
}
