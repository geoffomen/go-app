package zapImp

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
		FileLocation:      "/tmp/test_temp/info.log",
		ErrFileLevel:      "error",
		ErrFileLocation:   "/tmp/test_temp/back/err.log",
	})
	if err != nil {
		panic(fmt.Sprintf("failed to initrialize web framwork, err: %s", err))
	}
	zp.Infof("info")
	zp.Errorf("error")
}
