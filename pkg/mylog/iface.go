package mylog

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
	ins Iface
)

// SetInstance
func SetInstance(i Iface) {
	ins = i
}

// GetInstance ..
func GetInstance() Iface {
	return ins
}
