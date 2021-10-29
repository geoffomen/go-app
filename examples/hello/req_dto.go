package hello

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

type EchoReqDto struct {
	StrVal    string `json:"strVal"`
	IntVal    int    `json:"intVal"`
	IntPtrVal *int   `json:"intPtrVal"`
	StructVal struct {
		Id int `json:"id"`
	} `json:"structVal"`
	SliceVal []int `json:"sliceVal"`
}

func (s EchoReqDto) Validate() ([]string, error) {
	rt := make([]string, 0)

	sv := strings.TrimSpace(s.StrVal)
	if utf8.RuneCountInString(sv) <= 0 || utf8.RuneCountInString(sv) > 50 {
		rt = append(rt, "strVal: 必填且长度不能大于50")
	}
	if s.IntVal == 0 {
		rt = append(rt, "intVal: 必填且大于0")
	}
	if s.IntPtrVal == nil {
		rt = append(rt, "intPtrVal: 必填且大于0")
	}
	if s.StructVal.Id == 0 {
		rt = append(rt, "structVal.id: 必填且大于0")
	}

	if len(rt) > 0 {
		return rt, fmt.Errorf("校验失败")
	}
	return rt, nil
}
