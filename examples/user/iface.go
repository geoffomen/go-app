package user

import "github.com/geoffomen/go-app/internal/pkg/database"

// Iface 用户模块导出函数
type Iface interface {
	NewWithDb(cl *database.Client) Iface
	Create(param CreateUserRequestDto) (int, error)
	GetUserInfo(param GetUserInfoRequestDto) (UserInfoResponseDto, error)
}
