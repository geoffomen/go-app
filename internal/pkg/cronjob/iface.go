package cronjob

import (
	"fmt"
	"sync"

	"github.com/geoffomen/go-app/internal/pkg/cronjob/robfigimp"
)

// Iface ..
type Iface interface {
	Start()
	Stop()
	AddScheduleFunc(schedule string, f func()) error
}

var (
	once     sync.Once
	ins      Iface
	isInited bool = false
)

// New
func New(profile string) Iface {
	once.Do(func() {
		if isInited {
			return
		}
		service, err := robfigimp.New()
		if err != nil {
			panic(fmt.Sprintf("failed to initrialize config component, err: %v", err))
		}
		ins = service
		isInited = true
	})
	return ins
}

// GetInstance ..
func GetInstance() Iface {
	return ins
}
