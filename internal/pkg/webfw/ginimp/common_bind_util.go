package ginimp

import (
	"encoding/json"
	"fmt"
	"path"
	"reflect"
	"strings"

	"github.com/geoffomen/go-app/internal/pkg/vo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func bindArgs(c *gin.Context, h func(c *gin.Context, ft reflect.Type) (reflect.Value, error)) ([]reflect.Value, error) {
	handler, ok := pathToHandler[c.Request.Method+path.Join(c.Request.URL.Path)]
	if !ok {
		return nil, fmt.Errorf("no handler under the path: %s", c.Request.URL.Path)
	}
	handlerType := reflect.TypeOf(handler)

	args := make([]reflect.Value, handlerType.NumIn())
	for i := 0; i < handlerType.NumIn(); i++ {
		fType := handlerType.In(i)
		fts := fType.String()
		switch fts {
		case "http.ResponseWriter":
			args[i] = reflect.ValueOf(c.Writer)
		case "http.Request":
			args[i] = reflect.ValueOf(*c.Request)
		case "*http.Request":
			args[i] = reflect.ValueOf(c.Request)
		case "io.ReadCloser":
			r, err := c.Request.GetBody()
			if err != nil {
				return nil, err
			}
			args[i] = reflect.ValueOf(r)
		case "multipart.Form":
			r, err := c.MultipartForm()
			if err != nil {
				return nil, err
			}
			args[i] = reflect.ValueOf(*r)
		case "*multipart.Form":
			r, err := c.MultipartForm()
			if err != nil {
				return nil, err
			}
			args[i] = reflect.ValueOf(r)
		case reflect.TypeOf(vo.SessionInfo{}).String():
			sessionInfo, exist := c.Get("sessionInfo")
			if !exist {
				return nil, fmt.Errorf("no session info")
			}
			sessInfo := sessionInfo.(vo.SessionInfo)
			args[i] = reflect.ValueOf(sessInfo)
		default:
			fv, err := h(c, fType)
			if err != nil {
				return nil, err
			}
			nv, ok := fv.Interface().(vo.Validate)
			if ok {
				msg, err := nv.Validate()
				if err != nil {
					sb := strings.Builder{}
					for _, item := range msg {
						sb.WriteString(item)
						sb.WriteString("; ")
					}
					return []reflect.Value{}, fmt.Errorf("%s", sb.String())
				}
			}
			args[i] = fv
		}
	}
	return args, nil
}

func bindStruct(c *gin.Context, ft reflect.Type, bd binding.Binding) (reflect.Value, error) {
	fv := reflect.New(ft)
	m := make(map[int]reflect.Value)
	sfs := []reflect.StructField{}
	for i := 0; i < ft.NumField(); i++ {
		field := ft.Field(i)
		if field.Type.Kind() == reflect.Struct || field.Type.Kind() == reflect.Ptr {
			switch field.Type.String() {
			case "vo.Mytime":
				var p string
				switch bd {
				case binding.Query:
					p = c.Query(LowerFirst(field.Name))
				default:
					p = c.PostForm(LowerFirst(field.Name))
				}
				tv := vo.Mytime{}
				json.Unmarshal([]byte(p), &tv)
				m[i] = reflect.ValueOf(tv)
			default:
				if fv.Elem().Field(i).CanSet() {
					fv, err := bindField(c, field, field.Type, bd)
					if err != nil {
						return reflect.Value{}, fmt.Errorf("args bind error: %s", err)
					}
					m[i] = fv
				}
			}
		} else {
			field.Tag = reflect.StructTag(fmt.Sprintf(`%s form:"%s"`, field.Tag, LowerFirst(field.Name)))
			sfs = append(sfs, field)
		}
	}
	nft := reflect.StructOf(sfs)
	nfv := reflect.New(nft)
	v := nfv.Interface()

	err := c.ShouldBindWith(v, bd)
	if err != nil {
		return reflect.Value{}, fmt.Errorf("args bind error: %s", err)
	}
	for i := 0; i < nft.NumField(); i++ {
		field := nft.Field(i)
		fv.Elem().FieldByName(field.Name).Set(reflect.ValueOf(v).Elem().FieldByName(field.Name))
	}
	// fv = reflect.ValueOf(v).Elem().Convert(ft)
	for k, val := range m {
		fv.Elem().Field(k).Set(val)
	}
	return reflect.Indirect(fv), nil
}

func bindField(c *gin.Context, field reflect.StructField, ft reflect.Type, bd binding.Binding) (reflect.Value, error) {
	fv := reflect.New(ft)
	switch ft.Kind() {
	case reflect.Ptr:
		sfv, _ := bindField(c, field, ft.Elem(), bd)
		nfv := reflect.New(ft.Elem())
		nfv.Elem().Set(sfv)
		fv.Elem().Set(nfv)
		return reflect.Indirect(fv), nil
	case reflect.Struct:
		nfv, err := bindStruct(c, ft, binding.Query)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("binding failed: %s", err)
		}
		return reflect.Indirect(nfv), nil
	default:
		sfs := []reflect.StructField{}
		field.Tag = reflect.StructTag(fmt.Sprintf(`%s form:"%s"`, field.Tag, LowerFirst(field.Name)))
		sfs = append(sfs, field)
		nft := reflect.StructOf(sfs)
		nfv := reflect.New(nft)
		v := nfv.Interface()

		err := c.ShouldBindWith(v, bd)
		if err != nil {
			return reflect.Value{}, fmt.Errorf("args bind error: %s", err)
		}
		fv.Elem().Set(reflect.Indirect(nfv.Elem().Field(0)))
		return reflect.Indirect(fv), nil
	}
}
