package simpleimp

import "log"

type mylogger struct{}

func New() *mylogger {
	return &mylogger{}
}

func (ml *mylogger) Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (ml *mylogger) Errorf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func (ml *mylogger) Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
