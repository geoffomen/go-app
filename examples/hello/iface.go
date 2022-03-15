package hello

import (
	"github.com/storm-5/go-app/pkg/database"
	"github.com/storm-5/go-app/pkg/webfw"
)

type Iface interface {
	NewWithDb(db *database.Client) Iface
	SayHello(sessInfo webfw.SessionInfo) (interface{}, error)
	Echo(param EchoReqDto) (EchoRspDto, error)
	Error() (string, error)
}
