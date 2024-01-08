package useraccountsrv

import (
	"ibingli.com/internal/app/example/echoargs/echoargsdm"
	"ibingli.com/internal/app/example/useraccount/useraccountdm"
	"ibingli.com/internal/pkg/myConfig"
	"ibingli.com/internal/pkg/myDatabase"
	"ibingli.com/internal/pkg/myLog"
)

// Service ...
type Service struct {
	config      myConfig.Iface
	logger      myLog.Iface
	db          myDatabase.Iface
	repo        *Repository
	echoargssrv echoargsdm.SrvIface
}

var srv *Service = &Service{}

// New ...
func New(
	config myConfig.Iface,
	logger myLog.Iface,
	db myDatabase.Iface,
	echoargssrv echoargsdm.SrvIface,
) *Service {
	*srv = Service{
		config:      config,
		logger:      logger,
		db:          db,
		echoargssrv: echoargssrv,
	}
	srv.repo = newRepo(logger)
	return srv
}

// GetInstance ..
func GetInstance() *Service {
	return srv
}

func (s *Service) NewWithDb(db myDatabase.Iface) useraccountdm.SrvIface {
	ns := *s
	s.db = db
	return &ns
}
