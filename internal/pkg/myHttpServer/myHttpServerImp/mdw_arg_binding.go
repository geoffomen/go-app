package myHttpServerImp

import (
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"

	"ibingli.com/internal/pkg/myHttpServer"
)

func argBindingHandler() func(ctx *Ctx) error {
	return func(ctx *Ctx) error {

		handlerType := ctx.getHandlerReflectType()

		args := make([]reflect.Value, handlerType.NumIn())
		for i := 0; i < handlerType.NumIn(); i++ {
			fType := handlerType.In(i)

			switch fType {
			case reflect.TypeOf((*http.ResponseWriter)(nil)).Elem():
				args[i] = reflect.ValueOf(ctx.responseWriter)
			case reflect.TypeOf((*http.Request)(nil)).Elem():
				args[i] = reflect.ValueOf(*ctx.request)
			case reflect.TypeOf((**http.Request)(nil)).Elem():
				args[i] = reflect.ValueOf(ctx.request)
			case reflect.TypeOf((*multipart.Form)(nil)).Elem():
				err := ctx.request.ParseMultipartForm(1 << 10)
				if err != nil {
					return err
				}
				args[i] = reflect.ValueOf(*ctx.request.MultipartForm)
			case reflect.TypeOf((**multipart.Form)(nil)).Elem():
				err := ctx.request.ParseMultipartForm(1 << 10)
				if err != nil {
					return err
				}
				args[i] = reflect.ValueOf(ctx.request.MultipartForm)
			case reflect.TypeOf((*myHttpServer.SessionInfo)(nil)).Elem():
				si := ctx.getSessionInfo()
				args[i] = reflect.ValueOf(*si)
			case reflect.TypeOf((**myHttpServer.SessionInfo)(nil)).Elem():
				si := ctx.getSessionInfo()
				args[i] = reflect.ValueOf(si)
			default:
				var rfv reflect.Value
				var err error
				contentType := ctx.request.Header.Get("Content-Type")
				if strings.Contains(contentType, "json") {
					rfv, err = jsonBinder(ctx.request, fType)
					if err != nil {
						return err
					}
				} else if strings.Contains(contentType, "form-urlencoded") {
					ctx.request.ParseForm()
					rfv, err = formBinder(ctx.request.Form, fType)
					if err != nil {
						return err
					}
				} else if strings.Contains(contentType, "multipart/form-data") {
					ctx.request.ParseMultipartForm(1 << 10)
					rfv, err = formBinder(ctx.request.Form, fType)
					if err != nil {
						return err
					}
				} else if contentType == "" {
					if len(ctx.request.URL.Query()) > 0 {
						ctx.request.ParseForm()
						rfv, err = formBinder(ctx.request.Form, fType)
						if err != nil {
							return err
						}
					} else {
						return fmt.Errorf("nothing to bind! Content-Type is empty and query string is empty")
					}
				} else {
					return fmt.Errorf("unsupported Content-Type: %s", contentType)
				}
				args[i] = rfv
			}
		}
		ctx.setArgs(args)
		return ctx.Next()
	}
}

func jsonBinder(r *http.Request, rft reflect.Type) (reflect.Value, error) {
	v := reflect.New(rft).Interface()
	if r.Body != nil {
		defer r.Body.Close()
		err := json.NewDecoder(r.Body).Decode(v)
		if err != nil && err != io.EOF {
			return reflect.Value{}, err
		}
	} else {
		return reflect.Value{}, fmt.Errorf("empty request body")
	}
	return reflect.Indirect(reflect.ValueOf(v)), nil
}

func formBinder(formData url.Values, rft reflect.Type) (reflect.Value, error) {
	var rfv reflect.Value
	var rftStru reflect.Type
	if rft.Kind() == reflect.Pointer {
		rftStru = rft.Elem()
		rfv = reflect.New(rft.Elem())
	} else {
		rftStru = rft
		rfv = reflect.New(rft)
	}

	err := bindStruct(rftStru, rfv, formData)

	if rft.Kind() == reflect.Pointer {
		return rfv, err
	} else {
		return rfv.Elem(), err
	}
}

// bindStruct 绑定结构体。嵌套的结构体有同名字段时，绑定行为是未定义的。
func bindStruct(rft reflect.Type, rfv reflect.Value, formData url.Values) error {

	for i := 0; i < rft.NumField(); i++ {
		field := rft.Field(i)
		fieldName := ""
		if field.Tag.Get("json") != "" {
			fieldName = field.Tag.Get("json")
		} else {
			fieldName = lowerFirst(field.Name)
		}
		err := bindField(fieldName, field.Type, rfv.Elem().Field(i), formData)
		if err != nil {
			return err
		}
	}

	return nil
}

func bindField(fieldName string, fieldType reflect.Type, destField reflect.Value, formData url.Values) error {
	strs := formData[fieldName]
	switch k := fieldType.Kind(); k {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if len(strs) > 0 {
			val, _ := strconv.ParseInt(strs[0], 10, 64)
			destField.SetInt(val)
		} else {
			destField.SetInt(0)
		}
	case reflect.Float32, reflect.Float64:
		if len(strs) > 0 {
			val, _ := strconv.ParseFloat(strs[0], 64)
			destField.SetFloat(val)
		} else {
			destField.SetFloat(0.0)
		}
	case reflect.String:
		if len(strs) > 0 {
			destField.SetString(strs[0])
		} else {
			destField.SetString("")
		}
	case reflect.Slice:
		switch et := fieldType.Elem().Kind(); et {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			v := reflect.MakeSlice(fieldType, 0, len(strs))
			for _, item := range strs {
				it, _ := strconv.ParseInt(item, 10, 64)
				v = reflect.Append(v, reflect.ValueOf(it).Convert(fieldType.Elem()))
			}
			destField.Set(v)
		case reflect.Float32, reflect.Float64:
			v := reflect.MakeSlice(fieldType, 0, len(strs))
			for _, item := range strs {
				it, _ := strconv.ParseFloat(item, 64)
				v = reflect.Append(v, reflect.ValueOf(it).Convert(fieldType.Elem()))
			}
			destField.Set(v)
		case reflect.String:
			destField.Set(reflect.ValueOf(strs))
		}
	case reflect.Struct: // 嵌套结构体只支持时间
		if fieldType == reflect.TypeOf((*time.Time)(nil)).Elem() {
			if len(strs) > 0 {
				t, err := time.Parse(time.RFC3339, strs[0])
				if err != nil {
					return err
				}
				destField.Set(reflect.ValueOf(t))
			} else {
				// pass
			}
		} else if fieldType == reflect.TypeOf((*myHttpServer.Mytime)(nil)).Elem() {
			if len(strs) > 0 {
				miliSec, err := strconv.ParseInt(strs[0], 10, 64)
				if err != nil {
					return err
				}
				t := myHttpServer.Mytime(time.UnixMilli(miliSec))
				destField.Set(reflect.ValueOf(t))
			} else {
				// pass
			}
		} else {
			pt := reflect.New(fieldType)
			bindStruct(fieldType, pt, formData)
			destField.Set(pt.Elem())
		}
	case reflect.Pointer:
		pt := reflect.New(fieldType.Elem())
		bindField(fieldName, fieldType.Elem(), pt.Elem(), formData)
		destField.Set(pt)
	default:
		return fmt.Errorf("not support type. Form binding only support [interger float string slice], otherwise use json binding")
	}
	return nil
}

func lowerFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
