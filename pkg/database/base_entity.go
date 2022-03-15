package database

import (
	"database/sql/driver"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

// BaseEntity ...
type BaseEntity struct {
	Id int `json:"id"`
	// 创建时间
	CreatedTime Mytime `json:"createdTime"`
	CreatedBy   int    `json:"createdBy"`
	// 最近修改时间
	UpdatedTime Mytime `json:"updatedTime"`
	UpdatedBy   int    `josn:"updatedBy"`
	// 删除时间
	DeletedTime Mytime `json:"deletedTime"`
	DeletedBy   int    `json:"deletedBy"`
}

// Mytime used with orm
type Mytime struct {
	Time time.Time
}

// UnmarshalJSON ..
func (d *Mytime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	if s == "null" || s == "0" {
		d.Time, _ = time.Parse("2006-01-02 15:04:05.000", "0001-01-01 00:00:00")
		return nil
	}
	var sec int64 = 0
	var miliSec int64 = 0
	if len(s) > 13 {
		sc, err := strconv.ParseInt(s[0:10], 10, 64)
		if err != nil {
			return nil
		}
		sec = sc
		msc, err := strconv.ParseInt(s[10:13], 10, 64)
		if err != nil {
			return nil
		}
		miliSec = msc

	} else if len(s) > 10 {
		sc, err := strconv.ParseInt(s[0:10], 10, 64)
		if err != nil {
			return nil
		}
		sec = sc
		msc, err := strconv.ParseInt(s[10:], 10, 64)
		if err != nil {
			return nil
		}
		miliSec = msc
	} else {
		sc, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return nil
		}
		sec = sc
	}
	tm := time.Unix(sec, miliSec*1000000)
	d.Time = tm
	return nil
}

// MarshalJSON ..
func (d Mytime) MarshalJSON() ([]byte, error) {
	if d.Time.Year() < 1800 {
		return json.Marshal(nil)
	}
	un := d.Time.UnixNano() / 1000000
	return json.Marshal(un)
}

// Scan convert from db
func (d *Mytime) Scan(b interface{}) (err error) {

	switch x := b.(type) {
	case time.Time:
		d.Time = x
	case []byte:
		t, err := time.Parse("2006-01-02 15:04:05.000", string(b.([]byte)))
		if err != nil {
			d.Time = time.Time{}
		}
		d.Time = t
	default:
		d.Time = time.Time{}
	}
	return nil
}

// Value write to db
func (d Mytime) Value() (driver.Value, error) {
	if d.Time.IsZero() {
		return `0001-01-01 00:00:00`, nil
	}
	v := d.Time.Format("2006-01-02 15:04:05.000")
	return v, nil
}
