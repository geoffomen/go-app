package echoargsdm

import "ibingli.com/internal/pkg/myDatabase"

// EchoargsEntity ..
type EchoargsEntity struct {
	myDatabase.BaseEntity // 基础字段
}

// TableName 表名
func (e EchoargsEntity) TableName() string {
	return "echoargs"
}
