package ginimp

import (
	"bytes"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"unicode"
)

func stack(skip int) string {
	buf := new(bytes.Buffer)
	for i := skip; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s:%d (0x%x)\n", file, line, pc)

		funcName := runtime.FuncForPC(pc).Name()
		arr := strings.Split(funcName, "/")
		name := ""
		if len(arr) > 0 {
			name = arr[len(arr)-1]
		} else {
			name = funcName
		}
		fmt.Fprintf(buf, "\t%s: %d\n", name, line)
	}
	return buf.String()
}

func reflectValueToString(args []reflect.Value) string {
	sb := strings.Builder{}
	sb.WriteString("[  ")
	for _, item := range args {
		if item.Type().Kind() == reflect.Ptr {
			sb.WriteString(item.Elem().Type().Name())
			sb.WriteString(": ")
			sb.WriteString(fmt.Sprintf("%s", item.Elem().Interface()))
			sb.WriteString("; ")
		} else {
			sb.WriteString(item.Type().Name())
			sb.WriteString(": ")
			sb.WriteString(fmt.Sprintf("%s", item.Interface()))
			sb.WriteString("; ")
		}
	}
	sb.WriteString("  ]")
	return sb.String()
}

func UpperFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func LowerFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
