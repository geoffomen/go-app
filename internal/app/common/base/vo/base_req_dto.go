package vo

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Validate 用于参数校验
type Validate interface {
	Validate() ([]string, error)
}

type PageRequestDto struct {
	Page     int    `json:"page"`
	PageSize int    `json:"pageSize"`
	HasTotal int    `json:"hasTotal"`
	Sort     string `json:"sort"`
	Offset   int    `json:"offset"` // 兼容旧模式
	Size     int    `jsoN:"size"`   // 兼容旧模式
}

// UnmarshalJSON ..
func (p *PageRequestDto) UnmarshalJSON(b []byte) error {
	// fmt.Printf("%s", b)
	var objmap map[string]json.RawMessage
	err := json.Unmarshal(b, &objmap)
	if err != nil {
		return fmt.Errorf("原因：%s, 参数: %s", err, string(b))
	}

	errMsgs := make([]string, 0)
	if v, ok := objmap["offset"]; ok {
		err = json.Unmarshal(v, &(p.Offset))
		if err != nil {
			errMsgs = append(errMsgs, fmt.Sprintf("offset: %s", err))
		}
	}
	if v, ok := objmap["page"]; ok {
		err = json.Unmarshal(v, &(p.Page))
		if err != nil {
			errMsgs = append(errMsgs, fmt.Sprintf("page: %s", err))
		}
	}
	if v, ok := objmap["pageSize"]; ok {
		err = json.Unmarshal(v, &(p.PageSize))
		if err != nil {
			errMsgs = append(errMsgs, fmt.Sprintf("pageSize: %s", err))
		}
	}
	if v, ok := objmap["hasTotal"]; ok {
		err = json.Unmarshal(v, &(p.HasTotal))
		if err != nil {
			errMsgs = append(errMsgs, fmt.Sprintf("hasTotal: %s", err))
		}
	}

	if p.PageSize == 0 {
		p.PageSize = 10
	}
	if p.PageSize > 1000 {
		p.PageSize = 1000
	}
	if p.Page > 0 {
		of := (p.Page - 1) * p.PageSize
		p.Offset = of
	}
	if p.Offset < 0 {
		p.Offset = 0
	}

	if v, ok := objmap["sort"]; ok {
		var sort string
		err = json.Unmarshal(v, &sort)
		if err != nil {
			errMsgs = append(errMsgs, fmt.Sprintf("sort: %s", err))
		}
		var sb strings.Builder
		fs := strings.Split(sort, ",")
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
		p.Sort = strings.TrimSuffix(sb.String(), ", ")
	}

	if len(errMsgs) > 0 {
		return fmt.Errorf(strings.Join(errMsgs, ";"))
	} else {
		return nil
	}
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
