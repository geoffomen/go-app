package zapimp

import (
	"fmt"
	"testing"
)

func Test_log(t *testing.T) {
	zp, err := New(Configuration{
		EnableConsole:     true,
		ConsoleJSONFormat: true,
		ConsoleLevel:      "debug",
		EnableFile:        true,
		FileJSONFormat:    true,
		FileLevel:         "info",
		FileLocation:      "/tmp/miis/back/info.log",
		ErrFileLevel:      "error",
		ErrFileLocation:   "/tmp/miis/back/err.log",
	})
	if err != nil {
		panic(fmt.Sprintf("failed to initrialize web framwork, err: %s", err))
	}
	zp.Debugf("debug")
	zp.Infof("info")
	zp.Warnf("warn")
	zp.Errorf("error")
}
