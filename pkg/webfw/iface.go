package webfw

import (
	"fmt"
	"sync"

	"github.com/geoffomen/go-app/pkg/webfw/ginimp"
)

// Webfw ...
type Iface interface {
	RegisterHandler(map[string]interface{})
	Start() error
}

type Configuration struct {
	Profile string
	Port    string
}

var (
	once     sync.Once
	ins      Iface
	isInited bool = false
)

// New ..
func New(conf Configuration, logger ginimp.GinLogger) Iface {
	once.Do(func() {
		if isInited {
			return
		}
		server, err := ginimp.New(ginimp.Configuration{
			Profile: conf.Profile,
			Port:    conf.Port,
		}, logger)
		if err != nil {
			panic(fmt.Sprintf("failed to initrialize web framwork, err: %v", err))
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
