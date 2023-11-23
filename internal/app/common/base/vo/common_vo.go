package vo

import (
	"context"
	"strconv"
	"time"
)

// SessionInfo 记录会话信息
type SessionInfo struct {
	Ctx               context.Context
	Uid               int    // 用户帐户ID
	TenantId          int    // 租户ID
	Token             string // 访问令牌
	TokenExpireAtMili int64  // 访问令牌过期时间，毫秒
}

func (s *SessionInfo) SetContext(ctx context.Context) *SessionInfo {
	s.Ctx = ctx
	return s
}

const CookieName = "example-app-session-id"

// 表示状态
const (
	FALSE int = iota
	TRUE
)

const (
	DayTimeYYYY_MM_DD_HH_mm_ssSSSFormat = "2006-01-02 15:04:05.000"
	DayTimeYYYY_MM_DD_HH_mm_ssFormat    = "2006-01-02 15:04:05"
	DayTimeYYYY_MM_DD_HH_mmFormat       = "2006-01-02 15:04"
	DayTimeYYYYMMDDHHmmssFormat         = "20060102150405"
	DayYYYY_MM_DDFormat                 = "2006-01-02"
	DayTimeYYYYMMDDFormat               = "20060102"
	DayTimeIsoFormat                    = "2006-01-02T15:04:05Z07:00"
)

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
	return time.Time(mt).MarshalJSON()
}
