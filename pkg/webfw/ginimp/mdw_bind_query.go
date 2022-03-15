package ginimp

import (
	"fmt"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func bindQuery() gin.HandlerFunc {
	return func(c *gin.Context) {
		if len(c.Request.URL.RawQuery) > 0 {
			args, err := bindArgs(c, bindWithQuery)
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

func bindWithQuery(c *gin.Context, ft reflect.Type) (reflect.Value, error) {
	switch ft.Kind() {
	case reflect.Int:
		return reflect.Value{}, fmt.Errorf("not supported type")
	case reflect.String:
		return reflect.Value{}, fmt.Errorf("not supported type")
	case reflect.Slice:
		return reflect.Value{}, fmt.Errorf("not supported type")
	case reflect.Struct:
		fv, err := bindStruct(c, ft, binding.Query)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("binding failed")
		}
		return fv, nil
	case reflect.Ptr:
		return reflect.Value{}, fmt.Errorf("not supported type")
	default:
		return reflect.Value{}, fmt.Errorf("not supported type")
	}
}
