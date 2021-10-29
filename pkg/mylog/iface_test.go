package mylog

import (
	"testing"
)

func Test_log(t *testing.T) {
	New(Configuration{
		EnableConsole:     true,
		ConsoleJSONFormat: false,
		ConsoleLevel:      "debug",
		EnableFile:        true,
		FileJSONFormat:    true,
		FileLevel:         "info",
		FileLocation:      "/tmp/miis/back/info.log",
		ErrFileLevel:      "error",
		ErrFileLocation:   "/tmp/miis/back/err.log",
	})
	ins := GetInstance()
	ins.Debugf("debug")
	ins.Infof("info")
	ins.Warnf("warn")
	ins.Errorf("error")
	// Panicf("panic")
	// Fatal("fatal")
}
