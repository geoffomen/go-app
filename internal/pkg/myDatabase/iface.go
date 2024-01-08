package myDatabase

import (
	"context"
	"database/sql"
	"time"
)

type Iface interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
}

// BaseEntity ...
type BaseEntity struct {
	Id int64
	// 创建时间
	CreatedTime time.Time
	CreatedBy   int64
	// 最近修改时间
	UpdatedTime time.Time
	UpdatedBy   int64
	// 删除时间
	DeletedTime time.Time
	DeletedBy   int64
}
