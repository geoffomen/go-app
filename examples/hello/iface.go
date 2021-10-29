package hello

import (
	"github.com/geoffomen/go-app/pkg/database"
	"github.com/geoffomen/go-app/pkg/vo"
)

type Iface interface {
	NewWithDb(db *database.Client) Iface
	SayHello(sessInfo vo.SessionInfo) (interface{}, error)
	Echo(param EchoReqDto) (EchoRspDto, error)
	Error() (string, error)
}
