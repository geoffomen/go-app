package echoargssrv

import (
	"ibingli.com/internal/app/example/echoargs/echoargsdm"
	"ibingli.com/internal/pkg/myHttpServer"
)

func (srv *Service) Echo(ctx *myHttpServer.SessionInfo, args *echoargsdm.EhcoReqDto) (*echoargsdm.EchoRspDto, error) {
	return &echoargsdm.EchoRspDto{EhcoReqDto: *args}, nil
}