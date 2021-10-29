package account

import (
	"github.com/geoffomen/go-app/internal/pkg/database"
	"github.com/geoffomen/go-app/internal/pkg/vo"
)

// Iface ..
type Iface interface {
	NewWithDb(cl *database.Client) Iface
	Register(param CreateRequestDto) (int, error)
	Login(param LoginRequestDto) (*LoginResponseDto, error)
	Logout(param vo.SessionInfo) (int, error)
	CreateToken(uid int) (string, error)
	ValidAndGetTokenData(tokenString string) (*vo.SessionInfo, error)
}
