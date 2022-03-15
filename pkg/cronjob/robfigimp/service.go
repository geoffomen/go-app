package robfigimp

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

// CornService ..
type CornService struct {
	cronIns *cron.Cron
	logger  LogIface
}

type LogIface interface {
	Printf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
}

func New(logger LogIface) (*CornService, error) {
	if logger == nil {
		return nil, fmt.Errorf("logger不能等于nil")
	}
	c := cron.New(cron.WithSeconds())
	return &CornService{cronIns: c, logger: logger}, nil
}

// Start ..
func (srv *CornService) Start() {
	go srv.cronIns.Run()
}

// AddScheduleFunc ..
func (srv *CornService) AddScheduleFunc(schedule string, f func()) error {
	wrapper := func() {
		defer func() {
			if err := recover(); err != nil {
				srv.logger.Printf("cron job occur error: %s, stack: %s", err, stack(1))
			}
		}()
		f()
	}
	srv.cronIns.AddFunc(schedule, wrapper)
	return nil
}

// Stop ..
func (srv *CornService) Stop() {
	srv.cronIns.Stop()
}
