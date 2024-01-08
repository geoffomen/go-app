package robfigImp

import "log"

type myLogger struct{}

func newLogger() *myLogger {
	return &myLogger{}
}

func (ml *myLogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (ml *myLogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (ml *myLogger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
