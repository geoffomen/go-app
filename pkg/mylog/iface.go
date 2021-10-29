package mylog

import (
	"fmt"
	"sync"

	"github.com/geoffomen/go-app/pkg/mylog/zapimp"
)

//Logger is our contract for the logger
type Iface interface {
	Debugf(format string, args ...interface{})
	Print(args ...interface{})
	Println(args ...interface{})
	Printf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
}

type Configuration struct {
	EnableConsole     bool
	ConsoleJSONFormat bool
	ConsoleLevel      string
	EnableFile        bool
	FileJSONFormat    bool
	FileLevel         string
	FileLocation      string
	ErrFileLevel      string
	ErrFileLocation   string
}

var (
	once     sync.Once
	ins      Iface
	isInited bool = false
)

// New ..
func New(conf Configuration) Iface {
	once.Do(func() {
		if isInited {
			return
		}
		service, err := zapimp.New(zapimp.Configuration{
			EnableConsole:     conf.EnableConsole,
			ConsoleJSONFormat: conf.ConsoleJSONFormat,
			ConsoleLevel:      conf.ConsoleLevel,
			EnableFile:        conf.EnableFile,
			FileJSONFormat:    conf.FileJSONFormat,
			FileLevel:         conf.FileLevel,
			FileLocation:      conf.FileLocation,
			ErrFileLevel:      conf.ErrFileLevel,
			ErrFileLocation:   conf.ErrFileLocation,
		})
		if err != nil {
			panic(fmt.Sprintf("failed to initrialize web framwork, err: %s", err))
		}
		ins = service
		isInited = true
	})

	return ins
}

func GetInstance() Iface {
	return ins
}