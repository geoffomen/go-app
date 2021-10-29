package robfigimp

import (
	"bytes"
	"fmt"
	"runtime"
	"strings"
)

// stack returns a nicely formatted stack frame, skipping skip frames.
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
