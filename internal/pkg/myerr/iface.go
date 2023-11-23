package myerr

type MyErrorIface interface {
	Error() string
	Marshal() string
	New(err error, additionalMsg ...string) MyErrorIface
	Newf(format string, args ...interface{}) MyErrorIface
	NewfWithCode(code int, format string, args ...interface{}) MyErrorIface
	AddMsgf(format string, args ...interface{}) MyErrorIface
	SetCode(code int) MyErrorIface
	GetCode() int
	SetLogLevel(logLevel int) MyErrorIface
	GetLogLevel() int
}
