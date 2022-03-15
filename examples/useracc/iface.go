package useracc

import (
	"github.com/storm-5/go-app/pkg/database"
	"github.com/storm-5/go-app/pkg/webfw"
)

// Iface ..
type Iface interface {
	NewWithDb(cl *database.Client) Iface
	Register(param CreateRequestDto) (int, error)
	Login(param LoginRequestDto) (*LoginResponseDto, error)
	Logout(param webfw.SessionInfo) (int, error)
	CreateToken(uid int) (string, error)
	ValidAndGetTokenData(tokenString string) (*webfw.SessionInfo, error)
}
