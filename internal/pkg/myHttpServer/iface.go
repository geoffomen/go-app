package myHttpServer

import (
	"context"
	"database/sql/driver"
	"fmt"
	"strconv"
	"time"
)

type Iface interface {
	Listen() error
	Shutdown() error
	AddRouter(map[string]interface{}) error
}

// SessionInfo 记录会话信息
type SessionInfo struct {
	Ctx               context.Context
	Uid               int64  // 用户帐户ID
	TenantId          int64  // 租户ID
	Token             string // 访问令牌
	TokenExpireAtMili int64  // 访问令牌过期时间，毫秒
}

func (s *SessionInfo) SetContext(ctx context.Context) *SessionInfo {
	s.Ctx = ctx
	return s
}

type Mytime time.Time

func (mt *Mytime) UnmarshalJSON(data []byte) error {
	millis, err := strconv.ParseInt(string(data), 10, 64)
	if err != nil {
		*mt = Mytime(time.Unix(0, 0))
		return err
	}
	*mt = Mytime(time.UnixMilli(millis))
	return nil
}

func (mt Mytime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%d", time.Time(mt).UnixMilli())), nil
}

func (mt Mytime) String() string {
	if time.Time(mt).IsZero() {
		return ""
	} else {
		return time.Time(mt).Format("2006-01-02 15:04:05.000")
	}
}

// Scan convert from db
func (mt *Mytime) Scan(v interface{}) (err error) {
	d, _ := time.Parse(time.DateTime, "0001-01-01 00:00:00.000")
	switch b := v.(type) {
	case []byte:
		t, err := time.Parse("2006-01-02 15:04:05.000", string(b))
		if err != nil {
			*mt = Mytime(d)
		} else {
			*mt = Mytime(t)
		}
	default:
		*mt = Mytime(d)
	}
	return nil
}

// Value write to db
func (d Mytime) Value() (driver.Value, error) {
	if time.Time(d).IsZero() {
		return `0001-01-01 00:00:00`, nil
	}
	v := time.Time(d).Format("2006-01-02 15:04:05.000")
	return v, nil
}
