package myHttpServerImp

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"reflect"
	"strings"
	"time"
)

func logHandler() func(ctx *Ctx) error {
	return func(ctx *Ctx) error {
		// Starting time
		start := time.Now()
		// process request
		err := ctx.Next()
		// End Time
		end := time.Now()
		//execution time
		cost := end.Sub(start)
		clientIP := getIP(ctx.request)
		path := ctx.request.URL.Path
		method := ctx.request.Method
		if err != nil {
			var statusCode int
			var msg string
			logLevel := LogLevelError
			if e, ok := err.(myErrorIface); ok {
				statusCode = e.GetCode()
				msg = e.Marshal()
				logLevel = Loglevel(e.GetLogLevel())
			} else {
				statusCode = http.StatusInternalServerError
				msg = err.Error()
			}

			var logMsg string
			args := ctx.getArgs()
			logMsg = fmt.Sprintf("status: %3d, cost: %6v, clientIp: %15s, method: %s, path: %s, args: %s, error: %s",
				statusCode,
				cost,
				clientIP,
				method,
				path,
				reflectValueToString(args),
				msg,
			)
			switch logLevel {
			case LogLevelDebug, LogLevelInfo:
				ctx.logger.Infof("%s", logMsg)
			default:
				ctx.logger.Errorf("%s", logMsg)
			}
		} else {
			ctx.logger.Infof("status: %3d, cost: %6v, clientIp: %15s, method: %s, path: %s, args: %s",
				http.StatusOK,
				cost,
				clientIP,
				method,
				path,
				reflectValueToString(ctx.getArgs()),
			)
		}

		return nil
	}
}

// getIP returns request real ip.
func getIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if net.ParseIP(ip) != nil {
		return ip
	}

	ip = r.Header.Get("X-Forward-For")
	for _, i := range strings.Split(ip, ",") {
		if net.ParseIP(i) != nil {
			return i
		}
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "unknow client ip"
	}
	if net.ParseIP(ip) != nil {
		return ip
	}

	return "unknow client ip"
}

func reflectValueToString(args []reflect.Value) string {
	sb := strings.Builder{}
	sb.WriteString("[  ")
	for _, item := range args {
		if item.Type().Kind() == reflect.Ptr {
			sb.WriteString(item.Elem().Type().Name())
			sb.WriteString(": ")
			b, _ := json.Marshal(item.Elem().Interface())
			sb.WriteString(string(b))
			sb.WriteString("; ")
		} else {
			sb.WriteString(item.Type().Name())
			sb.WriteString(": ")
			b, _ := json.Marshal(item.Interface())
			sb.WriteString(string(b))
			sb.WriteString("; ")
		}
	}
	sb.WriteString("  ]")
	return sb.String()
}
