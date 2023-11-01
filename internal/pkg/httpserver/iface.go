package httpserver

type HttpServerIface interface {
	Listen() error
	Shutdown() error
	AddRouter(map[string]interface{}) error
}
