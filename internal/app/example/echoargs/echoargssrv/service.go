package echoargssrv

import "example.com/internal/app/common/base/vo"

func (srv *Service) Echo(ctx vo.SessionInfo, args EhcoReqDto) (*EchoRspDto, error) {
	return &EchoRspDto{args}, nil
}
