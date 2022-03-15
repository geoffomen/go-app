package cronjob

// Iface ..
type Iface interface {
	Start()
	Stop()
	AddScheduleFunc(schedule string, f func()) error
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
