package vo

import (
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
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
				sb.WriteString(camelToUnderscore(f[0 : len(f)-1]))
				sb.WriteString(" ASC")
			} else if strings.HasSuffix(f, "-") {
				sb.WriteString(camelToUnderscore(f[0 : len(f)-1]))
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

// camelToUnderscore case to CamelToUnderscore style.
func camelToUnderscore(camelStr string) string {
	l := utf8.RuneCountInString(camelStr)
	ss := strings.Split(camelStr, "")

	// we just care about the key of idx map,
	// the value of map is meaningless
	idx := make(map[int]int, 1)

	var rs []rune
	for _, s := range camelStr {
		rs = append(rs, []rune(string(s))...)
	}

	for i := l - 1; i >= 0; {
		if unicode.IsUpper(rs[i]) {
			var start, end int
			end = i
			j := i - 1
			for ; j >= 0; j-- {
				if unicode.IsLower(rs[j]) {
					start = j + 1
					break
				}
			}
			// handle the case: "BBC" or "AaBBB" case
			if end == l-1 {
				idx[start] = 1
			} else {
				if start == end {
					// value=1 is meaningless
					idx[start] = 1
				} else {
					idx[start] = 1
					idx[end] = 1
				}
			}
			i = j - 1
		} else {
			i--
		}
	}

	for i := l - 1; i >= 0; i-- {
		ss[i] = strings.ToLower(ss[i])
		if _, ok := idx[i]; ok && i != 0 {
			ss = append(ss[0:i], append([]string{"_"}, ss[i:]...)...)
		}
	}

	return strings.Join(ss, "")
}
