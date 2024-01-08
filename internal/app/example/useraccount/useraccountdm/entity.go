package useraccountdm

import "ibingli.com/internal/pkg/myDatabase"

// UseraccountEntity 登录帐号
type UseraccountEntity struct {
	myDatabase.BaseEntity
	Account  string
	Password string
	Salt     string
	Status   int
	Name     string
	Avatar   string
	Phone    string
}

// TableName 表名
func (e UseraccountEntity) TableName() string {
	return "user_account"
}

// LogintokenEntity 已登录信息
type LogintokenEntity struct {
	myDatabase.BaseEntity
	Uid        int
	Token      string
	ExpireAt   int64
	Persistent int
}

// TableName 表名
func (e LogintokenEntity) TableName() string {
	return "login_token"
}
