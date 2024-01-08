package myCronJob

// Iface ..
type Iface interface {
	Start()
	Stop()
	AddScheduleFunc(schedule string, f func()) error
}
