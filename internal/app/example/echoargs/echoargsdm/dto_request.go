package echoargsdm

import (
	"fmt"
	"strings"
	"time"
	"unicode/utf8"
)

type EhcoReqDto struct {
	Id    int
	F32   float32
	F64   float64
	Email string
	Si    []int
	Sf32  []float32
	Sf64  []float64
	Ss    []string
	Tm    time.Time
}

func (s EhcoReqDto) Validate() ([]string, error) {
	rt := make([]string, 0)

	if s.Id == 0 {
		rt = append(rt, "id: 必填且不能为0")
	}
	if s.F32 == float32(0) {
		rt = append(rt, "f32: 必填且不能为0")
	}
	if s.F64 == float64(0) {
		rt = append(rt, "f64: 必填且不能为0")
	}
	email := strings.TrimSpace(s.Email)
	if utf8.RuneCountInString(email) <= 0 || utf8.RuneCountInString(email) > 50 {
		rt = append(rt, "email: 长度不能小于1或大于50")
	}
	if len(s.Si) == 0 {
		rt = append(rt, "si: 必填且不能为空")
	}
	if len(s.Sf32) == 0 {
		rt = append(rt, "sf32: 必填且不能为空")
	}
	if len(s.Sf64) == 0 {
		rt = append(rt, "sf64: 必填且不能为空")
	}
	if len(s.Ss) == 0 {
		rt = append(rt, "ss: 必填且不能为空")
	}
	if s.Tm.IsZero() {
		rt = append(rt, "tm: 必填且不能为0")
	}
	if len(rt) > 0 {
		return rt, fmt.Errorf("校验失败")
	}
	return rt, nil
}
