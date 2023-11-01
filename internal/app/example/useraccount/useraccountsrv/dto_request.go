package useraccountsrv

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

// LoginRequestDTO ...
type CreateRequestDto struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (s CreateRequestDto) Validate() ([]string, error) {
	rt := make([]string, 0)
	a := strings.TrimSpace(s.Account)
	if utf8.RuneCountInString(a) <= 0 || utf8.RuneCountInString(a) > 50 {
		rt = append(rt, "account: 长度不能小于0或大于50")
	}
	p := strings.TrimSpace(s.Password)
	if len(p) <= 0 || len(p) > 50 {
		rt = append(rt, "password: 长度不能小于0或大于50")
	}
	if len(rt) > 0 {
		return rt, fmt.Errorf("校验失败")
	}
	return rt, nil
}

// LoginRequestDTO ...
type LoginRequestDto struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (s LoginRequestDto) Validate() ([]string, error) {
	rt := make([]string, 0)
	a := strings.TrimSpace(s.Account)
	if utf8.RuneCountInString(a) <= 0 || utf8.RuneCountInString(a) > 50 {
		rt = append(rt, "account: 长度不能小于0或大于50")
	}
	p := strings.TrimSpace(s.Password)
	if len(p) <= 0 || len(p) > 50 {
		rt = append(rt, "password: 长度不能小于0或大于50")
	}
	if len(rt) > 0 {
		return rt, fmt.Errorf("校验失败")
	}
	return rt, nil
}
