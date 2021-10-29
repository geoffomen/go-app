package webfw

import (
	"sync"

	"github.com/geoffomen/go-app/internal/pkg/config"
	"github.com/geoffomen/go-app/internal/pkg/mylog"
	"github.com/geoffomen/go-app/internal/pkg/webfw/ginimp"
)

// Webfw ...
type Iface interface {
	RegisterHandler(map[string]interface{})
	Start() error
}

var (
	once     sync.Once
	ins      Iface
	isInited bool = false
)

// New ..
func New(conf config.Iface) Iface {
	once.Do(func() {
		if isInited {
			return
		}
		server, err := ginimp.New(ginimp.Configuration{
			Profile: conf.GetStringOrDefault("profile", "test"),
			Port:    conf.GetStringOrDefault("server.port", "8080"),
		})
		if err != nil {
			mylog.Panicf("failed to initrialize web framwork, err: %v", err)
		}
		ins = server
		isInited = true
	})

	return ins
}

// GetInstance ..
func GetInstance() Iface {
	return ins
}
