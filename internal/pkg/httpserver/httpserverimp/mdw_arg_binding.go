package httpserverimp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func argBindingHandler() func(ctx *Ctx) error {
	return func(ctx *Ctx) error {

		handlerType := ctx.getHandlerReflectType()

		args := make([]reflect.Value, handlerType.NumIn())
		for i := 0; i < handlerType.NumIn(); i++ {
			fType := handlerType.In(i)

			switch fts := fType.String(); fts {
			case "http.ResponseWriter":
				args[i] = reflect.ValueOf(ctx.responseWriter)
			case "http.Request":
				args[i] = reflect.ValueOf(*ctx.request)
			case "*http.Request":
				args[i] = reflect.ValueOf(ctx.request)
			case "multipart.Form":
				err := ctx.request.ParseMultipartForm(1 << 10)
				if err != nil {
					return err
				}
				args[i] = reflect.ValueOf(*ctx.request.MultipartForm)
			case "*multipart.Form":
				err := ctx.request.ParseMultipartForm(1 << 10)
				if err != nil {
					return err
				}
				args[i] = reflect.ValueOf(ctx.request.MultipartForm)
			case "vo.SessionInfo":
				si := ctx.getSessionInfo()
				args[i] = reflect.ValueOf(*si)
			case "*vo.SessionInfo":
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
	rfv := reflect.New(rft)
	for i := 0; i < rft.NumField(); i++ {
		field := rft.Field(i)
		fieldName := ""
		if field.Tag.Get("json") != "" {
			fieldName = field.Tag.Get("json")
		} else {
			fieldName = lowerFirst(field.Name)
		}
		strs := formData[fieldName]
		switch k := field.Type.Kind(); k {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			val, _ := strconv.ParseInt(strs[0], 10, 64)
			rfv.Elem().Field(i).SetInt(val)
		case reflect.Float32, reflect.Float64:
			val, _ := strconv.ParseFloat(strs[0], 64)
			rfv.Elem().Field(i).SetFloat(val)
		case reflect.String:
			rfv.Elem().Field(i).SetString(strs[0])
		case reflect.Slice:
			switch et := field.Type.Elem().Kind(); et {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				v := reflect.MakeSlice(field.Type, 0, len(strs))
				for _, item := range strs {
					it, _ := strconv.ParseInt(item, 10, 64)
					v = reflect.Append(v, reflect.ValueOf(it).Convert(field.Type.Elem()))
				}
				rfv.Elem().Field(i).Set(v)
			case reflect.Float32, reflect.Float64:
				v := reflect.MakeSlice(field.Type, 0, len(strs))
				for _, item := range strs {
					it, _ := strconv.ParseFloat(item, 64)
					v = reflect.Append(v, reflect.ValueOf(it).Convert(field.Type.Elem()))
				}
				rfv.Elem().Field(i).Set(v)
			case reflect.String:
				rfv.Elem().Field(i).Set(reflect.ValueOf(strs))
			}
		case reflect.Struct:
			if strings.EqualFold(field.Type.String(), "time.Time") {
				t, err := time.Parse(time.RFC3339, strs[0])
				if err != nil {
					return reflect.Value{}, err
				}
				rfv.Elem().Field(i).Set(reflect.ValueOf(t))
			} else {
				return reflect.Value{}, fmt.Errorf("form binding not support nested struct")
			}
		default:
			return reflect.Value{}, fmt.Errorf("not support type. Form binding only support [interger float string slice], otherwise use json binding")
		}
	}
	return reflect.Indirect(rfv), nil
}

func lowerFirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}
