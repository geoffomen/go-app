package userimp

import (
	"github.com/geoffomen/go-app/pkg/vo"
)

type UserEntity struct {
	vo.BaseEntity        // 包含相关属性
	Name          string `json:"name"`
	NickName      string `json:"nickName"`
	Avatar        string `json:"avatar"`
	Phone         string `json:"phone"`
	Status        int    `json:"status"`
}

func (u UserEntity) TableName() string {
	return "user"
}
