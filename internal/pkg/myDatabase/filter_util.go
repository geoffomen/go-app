package myDatabase

import (
	"reflect"
	"strings"
)

type Filter struct {
	selectColumns []string
	args          []any
	queries       []string
	offset        int64
	limit         int64
	order         string
	isTotal       bool
}

func NewFilter() *Filter {
	return &Filter{
		selectColumns: make([]string, 0),
		args:          make([]any, 0),
		queries:       make([]string, 0),
		isTotal:       false,
	}
}

func (f *Filter) Select(col string) *Filter {
	s := strings.Split(col, ",")
	for _, item := range s {
		f.selectColumns = append(f.selectColumns, strings.TrimSpace(item))
	}

	return f
}

func (f *Filter) GetSelect() []string {
	return f.selectColumns
}

func (f *Filter) BuildSelectString() string {
	if len(f.selectColumns) == 0 {
		return "*"
	} else {
		return strings.Join(f.selectColumns, ", ")
	}
}

// Where 过虑条件，例如：
// Where("col = ?", 1);
// Where("key in (?, ?)", val1, val2);
// Where("key in ?", []int{1,2});
// Where("key = ?", val)
func (f *Filter) Where(query string, value ...any) *Filter {
	if value == nil { // Where("key=val")
		f.queries = append(f.queries, query)
		return f
	} else if cnt := len(value); cnt > 1 { // Where("key in (?, ?)", val1, val2)
		f.queries = append(f.queries, query)
		f.args = append(f.args, value...)
	} else { // Where("key in ?", []int{1,2}) 或 Where("key = ?", val)
		rv := reflect.Indirect(reflect.ValueOf(value[0]))
		switch rv.Kind() {
		case reflect.Slice, reflect.Array:
			nvs := make([]interface{}, rv.Len())
			s := make([]string, 0, cnt)
			for i := 0; i < rv.Len(); i++ {
				nvs[i] = rv.Index(i).Interface()
				s = append(s, "?")
			}
			f.queries = append(f.queries, strings.Replace(query, "?", strings.Join(s, ","), 1))
			f.args = append(f.args, nvs...)
		default:
			f.queries = append(f.queries, query)
			f.args = append(f.args, value...)
		}
	}

	return f
}

func (f *Filter) GetArgs() []interface{} {
	return f.args
}

func (f *Filter) BuildQueryString() string {
	return strings.Join(f.queries, " AND ")
}

func (f *Filter) Offset(offset int64) *Filter {
	f.offset = offset
	return f
}

func (f *Filter) GetOffset() int64 {
	return f.offset
}

func (f *Filter) Limit(limit int64) *Filter {
	f.limit = limit
	return f
}

func (f *Filter) GetLimit() int64 {
	if f.limit == 0 {
		return 10
	} else {
		return f.limit
	}
}

func (f *Filter) Order(order string) *Filter {
	f.order = order
	return f
}

func (f *Filter) GetOrder() string {
	if f.order == "" {
		return "id ASC"
	} else {
		return f.order
	}
}

func (f *Filter) IsTotal(b bool) *Filter {
	f.isTotal = b
	return f
}

func (f *Filter) GetIsTotal() bool {
	return f.isTotal
}
