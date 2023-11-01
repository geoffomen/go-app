package entity

import (
	"time"
)

// BaseEntity ...
type BaseEntity struct {
	Id int
	// 创建时间
	CreatedTime time.Time
	CreatedBy   int
	// 最近修改时间
	UpdatedTime time.Time
	UpdatedBy   int
	// 删除时间
	DeletedTime time.Time
	DeletedBy   int
}
