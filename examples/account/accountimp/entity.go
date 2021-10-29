package accountimp

import "github.com/geoffomen/go-app/internal/pkg/vo"

// AccountEntity 登录帐号
type AccountEntity struct {
	vo.BaseEntity // 基础字段
	Account       string
	Password      string
	Salt          string
	Status        int
	Uid           int
}

// TableName 表名
func (e AccountEntity) TableName() string {
	return "account_info"
}

// LoginTokenEntity 已登录信息
type LoginTokenEntity struct {
	vo.BaseEntity // 基础字段
	UID           int
	Token         string
	ExpireAt      int64
	Persistent    int
}

// TableName 表名
func (e LoginTokenEntity) TableName() string {
	return "login_token"
}
