package mylog

import (
	"testing"

	"github.com/geoffomen/go-app/internal/pkg/config"
)

func Test_log(t *testing.T) {
	conf := config.NewEmpty()
	conf.Set("log.EnableConsole", true)
	conf.Set("log.ConsoleJSONFormat", true)
	conf.Set("log.ConsoleLevel", "debug")
	conf.Set("log.EnableFile", true)
	conf.Set("log.FileJSONFormat", true)
	conf.Set("log.FileLevel", "info")
	conf.Set("log.FileLocation", "/tmp/log/myapp/info.log")
	conf.Set("log.ErrFileLevel", "error")
	conf.Set("log.ErrFileLocation", "/tmp/log/myapp/error.log")
	New(conf)

	Debugf("debug")
	Infof("info")
	Warnf("warn")
	Errorf("error")
	// Panicf("panic")
	// Fatal("fatal")
}
