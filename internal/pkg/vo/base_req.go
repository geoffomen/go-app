package vo

import (
	"fmt"
	"strings"

	"github.com/geoffomen/go-app/internal/pkg/stringutil"
)

type PageRequestDto struct {
	Page     int    `json:"page"`
	Offset   int    `json:"offset"`
	PageSize int    `json:"pageSize"`
	HasTotal int    `json:"hasTotal"`
	Sort     string `json:"sort"`
}

func (s *PageRequestDto) Validate() ([]string, error) {
	rt := make([]string, 0)

	if s.Page == 0 {
		rt = append(rt, "page: 必填且不能等于0")
	}
	if s.PageSize == 0 {
		rt = append(rt, "pageSize: 必填且不能等于0")
	}
	s.Offset = (s.Page - 1) * s.PageSize
	if s.Sort != "" {
		var sb strings.Builder
		fs := strings.Split(s.Sort, ",")
		for _, item := range fs {
			f := strings.Trim(item, " ")
			if strings.HasSuffix(f, "+") {
				sb.WriteString(stringutil.CamelToUnderscore(f[0 : len(f)-1]))
				sb.WriteString(" ASC")
			} else if strings.HasSuffix(f, "-") {
				sb.WriteString(stringutil.CamelToUnderscore(f[0 : len(f)-1]))
				sb.WriteString(" DESC")
			}
			sb.WriteString(", ")
		}
		s.Sort = strings.TrimSuffix(sb.String(), ", ")
	}

	if len(rt) > 0 {
		return rt, fmt.Errorf("校验失败")
	}
	return rt, nil
}
