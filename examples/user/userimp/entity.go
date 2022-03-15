package userimp

import (
	"github.com/storm-5/go-app/pkg/database"
)

type UserEntity struct {
	database.BaseEntity        // 包含相关属性
	Name                string `json:"name"`
	NickName            string `json:"nickName"`
	Avatar              string `json:"avatar"`
	Phone               string `json:"phone"`
	Status              int    `json:"status"`
}

func (u UserEntity) TableName() string {
	return "user"
}
