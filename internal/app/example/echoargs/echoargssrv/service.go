package echoargssrv

import (
	"example.com/internal/app/common/base/vo"
	"example.com/internal/app/example/echoargs/echoargsdm"
)

func (srv *Service) Echo(ctx vo.SessionInfo, args echoargsdm.EhcoReqDto) (*echoargsdm.EchoRspDto, error) {
	return &echoargsdm.EchoRspDto{EhcoReqDto: args}, nil
}
