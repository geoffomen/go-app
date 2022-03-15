package user

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/geoffomen/go-app/pkg/database"
	"github.com/geoffomen/go-app/pkg/webfw"
)

type CreateUserRequestDto struct {
	Name     string `json:"name"`
	NickName string `json:"nickName"`
	Avatar   string `json:"avatar"`
	Phone    string `json:"phone"`
}

func (s CreateUserRequestDto) Validate() ([]string, error) {
	rt := make([]string, 0)
	n := strings.TrimSpace(s.Name)
	if utf8.RuneCountInString(n) <= 0 || utf8.RuneCountInString(n) > 50 {
		rt = append(rt, "name: 长度不能小于0或大于50")
	}

	if len(rt) > 0 {
		return rt, fmt.Errorf("校验失败")
	}
	return rt, nil
}

type GetUserInfoRequestDto struct {
	Id int `json:"id"`
}

func (s GetUserInfoRequestDto) Validate() ([]string, error) {
	rt := make([]string, 0)
	if s.Id == 0 {
		rt = append(rt, "id: 必填且不能为0")
	}

	if len(rt) > 0 {
		return rt, fmt.Errorf("校验失败")
	}
	return rt, nil
}

type PageRequestDto struct {
	*webfw.PageRequestDto
	K3      *int   `json:"k3"`
	Keyword string `json:"keyword"`
	S1      struct {
		K1 int  `json:"k1"`
		K2 *int `json:"k2"`
	} `json:"s1"`
	Tm database.Mytime `json:"tm"`
	Sl []int           `json:"sl"`
}

func (s PageRequestDto) Validate() ([]string, error) {
	rt := make([]string, 0)

	data, err := s.PageRequestDto.Validate()
	if err != nil {
		rt = append(rt, data...)
	}

	if len(rt) > 0 {
		return rt, fmt.Errorf("校验失败")
	}
	return rt, nil

}
