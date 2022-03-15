package webfw

// Webfw ...
type Iface interface {
	RegisterHandler(map[string]interface{})
	Start() error
}

var (
	ins Iface
)

func SetInstance(i Iface) {
	ins = i
}

// GetInstance ..
func GetInstance() Iface {
	return ins
}
