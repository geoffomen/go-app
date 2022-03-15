package useraccimp

import "github.com/geoffomen/go-app/pkg/database"

// AccountEntity 登录帐号
type AccountEntity struct {
	database.BaseEntity // 基础字段
	Account       string
	Password      string
	Salt          string
	Status        int
	Uid           int
}

// TableName 表名
func (e AccountEntity) TableName() string {
	return "user_account"
}

// LoginTokenEntity 已登录信息
type LoginTokenEntity struct {
	database.BaseEntity // 基础字段
	UID           int
	Token         string
	ExpireAt      int64
	Persistent    int
}

// TableName 表名
func (e LoginTokenEntity) TableName() string {
	return "login_token"
}
