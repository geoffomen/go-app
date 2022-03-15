package ginimp

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin"
)

func init() {
	middleWares = append(middleWares, bindJson())
}

func bindJson() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.GetHeader("Content-Type"), "application/json") {
			args, err := bindArgs(c, bindWithJson)
			if err != nil {
				c.Error(fmt.Errorf("%s", stack(1)))
				c.Error(err)
				c.Status(400)
				c.Abort()
				return
			}
			c.Set("args", args)
		}
		c.Next()
	}
}

func bindWithJson(c *gin.Context, ft reflect.Type) (reflect.Value, error) {
	switch ft.Kind() {
	case reflect.Int:
		return reflect.Value{}, fmt.Errorf("not supported type")
	case reflect.String:
		return reflect.Value{}, fmt.Errorf("not supported type")
	case reflect.Slice:
		return reflect.Value{}, fmt.Errorf("not supported type")
	case reflect.Struct:
		v := reflect.New(ft).Interface()
		err := c.ShouldBindJSON(v)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("args bind error: %s", err)
		}
		nfv := reflect.Indirect(reflect.ValueOf(v))
		return nfv, nil
	case reflect.Ptr:
		return reflect.Value{}, fmt.Errorf("not supported type")
	default:
		return reflect.Value{}, fmt.Errorf("not supported type")
	}
}
