package echoargssrv

import (
	"ibingli.com/internal/app/example/echoargs/echoargsdm"
	"ibingli.com/internal/app/example/useraccount/useraccountdm"
	"ibingli.com/internal/pkg/myConfig"
	"ibingli.com/internal/pkg/myDatabase"
	"ibingli.com/internal/pkg/myLog"
)

// Service ...
type Service struct {
	config         myConfig.Iface
	logger         myLog.Iface
	db             myDatabase.Iface
	repo           *EchoargsRepository
	useraccountSrv useraccountdm.SrvIface
}

var instance *Service = &Service{}

// New ...
func New(
	config myConfig.Iface,
	logger myLog.Iface,
	db myDatabase.Iface,
	uSrv useraccountdm.SrvIface,
) *Service {
	*instance = Service{
		config:         config,
		logger:         logger,
		db:             db,
		useraccountSrv: uSrv,
	}
	instance.repo = newRepo(logger)
	return instance
}

// GetInstance ..
func GetInstance() *Service {
	return instance
}

func (s *Service) NewWithDb(db myDatabase.Iface) echoargsdm.SrvIface {
	ns := *s
	s.db = db
	return &ns
}
