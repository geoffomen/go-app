package useraccountrepo

type LoggerIface interface {
	Infof(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}