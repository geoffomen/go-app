package mylog

import (
	"fmt"
	"sync"

	"github.com/geoffomen/go-app/internal/pkg/config"
	"github.com/geoffomen/go-app/internal/pkg/mylog/zapimp"
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
		service, err := zapimp.New(zapimp.Configuration{
			EnableConsole:     conf.GetBoolOrDefault("log.enableConsole", true),
			ConsoleJSONFormat: conf.GetBoolOrDefault("log.consoleJSONFormat", true),
			ConsoleLevel:      conf.GetStringOrDefault("log.consoleLevel", "debug"),
			EnableFile:        conf.GetBoolOrDefault("log.enableFile", true),
			FileJSONFormat:    conf.GetBoolOrDefault("log.fileJSONFormat", true),
			FileLevel:         conf.GetStringOrDefault("log.fileLevel", "info"),
			FileLocation:      conf.GetStringOrDefault("log.fileLocation", "/tmp/miis/back/info.log"),
			ErrFileLevel:      conf.GetStringOrDefault("log.errFileLevel", "error"),
			ErrFileLocation:   conf.GetStringOrDefault("log.errFileLocation", "/tmp/miis/back/err.log"),
		})
		if err != nil {
			panic(fmt.Sprintf("failed to initrialize web framwork, err: %v", err))
		}
		ins = service
		isInited = true
	})

	return ins
}

func GetInstance() Iface {
	return ins
}

// Debugf ..
func Debugf(format string, args ...interface{}) {
	ins.Debugf(format, args...)
}

// Print ..
func Print(args ...interface{}) {
	ins.Infof("%v", args...)
}

// Println ..
func Println(args ...interface{}) {
	ins.Infof("%v\n", args...)
}

// Printf ..
func Printf(format string, args ...interface{}) {
	ins.Infof(format, args...)
}

// Infof ..
func Infof(format string, args ...interface{}) {
	ins.Infof(format, args...)
}

// Warnf ..
func Warnf(format string, args ...interface{}) {
	ins.Warnf(format, args...)
}

// Errorf ..
func Errorf(format string, args ...interface{}) {
	ins.Errorf(format, args...)
}

// Panicf ..
func Panicf(format string, args ...interface{}) {
	ins.Panicf(format, args...)
}

// Fatalf ..
func Fatalf(format string, args ...interface{}) {
	ins.Fatalf(format, args...)
}
