package hello

import (
	"github.com/geoffomen/go-app/pkg/database"
	"github.com/geoffomen/go-app/pkg/webfw"
)

type Iface interface {
	NewWithDb(db *database.Client) Iface
	SayHello(sessInfo webfw.SessionInfo) (interface{}, error)
	Echo(param EchoReqDto) (EchoRspDto, error)
	Error() (string, error)
}
