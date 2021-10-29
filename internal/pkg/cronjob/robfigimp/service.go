package robfigimp

import (
	"log"

	"github.com/robfig/cron/v3"
)

// CornService ..
type CornService struct {
	cronIns *cron.Cron
}

func New() (*CornService, error) {
	c := cron.New(cron.WithSeconds())
	return &CornService{cronIns: c}, nil
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
				log.Printf("cron job occur error: %s, stack: %s", err, stack(1))
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
