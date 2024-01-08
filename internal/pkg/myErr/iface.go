package myErr

type Iface interface {
	Error() string
	Marshal() string
	New(err error, additionalMsg ...string) Iface
	Newf(format string, args ...interface{}) Iface
	NewfWithCode(code int, format string, args ...interface{}) Iface
	AddMsgf(format string, args ...interface{}) Iface
	SetCode(code int) Iface
	GetCode() int
	SetLogLevel(logLevel int) Iface
	GetLogLevel() int
}
