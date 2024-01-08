package robfigImp

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

// CornService ..
type CornService struct {
	cronIns *cron.Cron
	logger  loggerIface
}

func New(logger loggerIface) (*CornService, error) {
	c := cron.New(cron.WithSeconds())
	return &CornService{cronIns: c, logger: logger}, nil
}

// Start ..
func (srv *CornService) Start() {
	srv.cronIns.Run()
}

// AddScheduleFunc ..
func (srv *CornService) AddScheduleFunc(schedule string, f func()) error {
	wrapper := func() {
		defer func() {
			if err := recover(); err != nil {
				srv.logger.Errorf("cron job occur error: %s, stack: %s", err, stack(2))
			}
		}()
		f()
	}
	jobId, err := srv.cronIns.AddFunc(schedule, wrapper)
	if err != nil {
		panic(fmt.Sprintf("定时器注册失败, err: %s", err))
	} else {
		srv.logger.Infof("定时器注册成功, jobId: %d", jobId)
	}
	return err
}

// Stop ..
func (srv *CornService) Stop() {
	srv.cronIns.Stop()
}
