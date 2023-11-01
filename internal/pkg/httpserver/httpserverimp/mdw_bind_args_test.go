package httpserverimp

import (
	"bytes"
	"encoding/json"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestBindJson(t *testing.T) {
	type ReqJsonStruct struct {
		Id    int
		F32   float32
		F64   float64
		Email string
		Si    []int
		Sf32  []float32
		Sf64  []float64
		Ss    []string
		Tm    time.Time
		User  struct {
			Uid int
		}
	}
	arg := ReqJsonStruct{
		Id:    1,
		F32:   3.2,
		F64:   6.4,
		Email: "email",
		Si:    []int{10, 100},
		Sf32:  []float32{3.21, 3.22},
		Sf64:  []float64{6.41, 6.42},
		Ss:    []string{"str1", "str2"},
		Tm:    time.UnixMilli(1234567890),
		User: struct{ Uid int }{
			Uid: 1000000,
		},
	}
	b, _ := json.Marshal(arg)
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1", bytes.NewReader(b))
	req.Header.Add("Content-Type", "json")
	handler := func(arg ReqJsonStruct) (interface{}, error) {
		assert.Equal(t, arg.Id, 1)
		assert.Equal(t, arg.F32, float32(3.2))
		assert.Equal(t, arg.F64, float64(6.4))
		assert.Equal(t, arg.Email, "email")
		assert.Equal(t, arg.Si, []int{10, 100})
		assert.Equal(t, arg.Sf32, []float32{float32(3.21), float32(3.22)})
		assert.Equal(t, arg.Sf64, []float64{float64(6.41), float64(6.42)})
		assert.Equal(t, arg.Ss, []string{"str1", "str2"})
		assert.Equal(t, arg.Tm, time.UnixMilli(1234567890))
		assert.Equal(t, arg.User.Uid, 1000000)
		return nil, nil
	}
	handlerType := reflect.TypeOf(handler)
	ctx := &Ctx{handlerReflectType: handlerType, request: req}
	_ = bindArgs()(ctx)
	args := ctx.args
	handlerValue := reflect.ValueOf(handler)
	handlerValue.Call(args)
}

func TestBindQuery(t *testing.T) {
	type ReqQueryStruct struct {
		Id    int
		F32   float32
		F64   float64
		Email string
		Si    []int
		Sf32  []float32
		Sf64  []float64
		Ss    []string
		Tm    time.Time
		// St struct{Uid int}
	}
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1/?id=1&f32=3.2&f64=6.4&email=email&si=100&si=1000&sf32=3.21&sf32=3.22&sf64=6.41&sf64=6.42&ss=str1&ss=str2&tm=1234567890123", bytes.NewReader(nil))
	req.Header.Add("Content-Type", "")
	handler := func(arg ReqQueryStruct) (interface{}, error) {
		assert.Equal(t, arg.Id, 1)
		assert.Equal(t, arg.F32, float32(3.2))
		assert.Equal(t, arg.F64, float64(6.4))
		assert.Equal(t, arg.Email, "email")
		assert.Equal(t, arg.Si, []int{100, 1000})
		assert.Equal(t, arg.Sf32, []float32{float32(3.21), float32(3.22)})
		assert.Equal(t, arg.Sf64, []float64{float64(6.41), float64(6.42)})
		assert.Equal(t, arg.Ss, []string{"str1", "str2"})
		assert.Equal(t, arg.Tm, time.UnixMilli(1234567890123))
		return nil, nil
	}
	handlerType := reflect.TypeOf(handler)
	ctx := &Ctx{handlerReflectType: handlerType, request: req}
	_ = bindArgs()(ctx)
	args := ctx.args
	handlerValue := reflect.ValueOf(handler)
	handlerValue.Call(args)
}

func TestBindFormUrlEncode(t *testing.T) {
	type ReqQueryStruct struct {
		Id    int
		F32   float32
		F64   float64
		Email string
		Si    []int
		Sf32  []float32
		Sf64  []float64
		Ss    []string
		Tm    time.Time
		// St struct{Uid int}
	}
	formData := url.Values{}
	formData.Add("id", "1")
	formData.Add("f32", "3.2")
	formData.Add("f64", "6.4")
	formData.Add("email", "email")
	formData.Add("si", "10000")
	formData.Add("si", "100000")
	formData.Add("sf32", "3.21")
	formData.Add("sf32", "3.22")
	formData.Add("sf64", "6.41")
	formData.Add("sf64", "6.42")
	formData.Add("ss", "str1")
	formData.Add("ss", "str2")
	formData.Add("tm", "1234567890123")
	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1/", strings.NewReader(formData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	handler := func(arg ReqQueryStruct) (interface{}, error) {
		assert.Equal(t, arg.Id, 1)
		assert.Equal(t, arg.F32, float32(3.2))
		assert.Equal(t, arg.F64, float64(6.4))
		assert.Equal(t, arg.Email, "email")
		assert.Equal(t, arg.Si, []int{10000, 100000})
		assert.Equal(t, arg.Sf32, []float32{float32(3.21), float32(3.22)})
		assert.Equal(t, arg.Sf64, []float64{float64(6.41), float64(6.42)})
		assert.Equal(t, arg.Ss, []string{"str1", "str2"})
		assert.Equal(t, arg.Tm, time.UnixMilli(1234567890123))
		return nil, nil
	}
	handlerType := reflect.TypeOf(handler)
	ctx := &Ctx{handlerReflectType: handlerType, request: req}
	_ = bindArgs()(ctx)
	args := ctx.args
	handlerValue := reflect.ValueOf(handler)
	handlerValue.Call(args)
}

func TestBindMultipartForm(t *testing.T) {
	type ReqQueryStruct struct {
		Id    int
		F32   float32
		F64   float64
		Email string
		Si    []int
		Sf32  []float32
		Sf64  []float64
		Ss    []string
		Tm    time.Time
		// St struct{Uid int}
	}
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	fw, _ := w.CreateFormField("id")
	fw.Write([]byte("1"))
	fw, _ = w.CreateFormField("f32")
	fw.Write([]byte("3.2"))
	fw, _ = w.CreateFormField("f64")
	fw.Write([]byte("6.4"))
	fw, _ = w.CreateFormField("email")
	fw.Write([]byte("email"))
	fw, _ = w.CreateFormField("si")
	fw.Write([]byte("1000000"))
	fw, _ = w.CreateFormField("si")
	fw.Write([]byte("10000000"))
	fw, _ = w.CreateFormField("sf32")
	fw.Write([]byte("3.21"))
	fw, _ = w.CreateFormField("sf32")
	fw.Write([]byte("3.22"))
	fw, _ = w.CreateFormField("sf64")
	fw.Write([]byte("6.41"))
	fw, _ = w.CreateFormField("sf64")
	fw.Write([]byte("6.42"))
	fw, _ = w.CreateFormField("ss")
	fw.Write([]byte("str1"))
	fw, _ = w.CreateFormField("ss")
	fw.Write([]byte("str2"))
	fw, _ = w.CreateFormField("tm")
	fw.Write([]byte("1234567890123"))
	w.Close()

	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1/", &body)
	req.Header.Add("Content-Type", w.FormDataContentType())
	handler := func(arg ReqQueryStruct) (interface{}, error) {
		assert.Equal(t, arg.Id, 1)
		assert.Equal(t, arg.F32, float32(3.2))
		assert.Equal(t, arg.F64, float64(6.4))
		assert.Equal(t, arg.Email, "email")
		assert.Equal(t, arg.Si, []int{1000000, 10000000})
		assert.Equal(t, arg.Sf32, []float32{float32(3.21), float32(3.22)})
		assert.Equal(t, arg.Sf64, []float64{float64(6.41), float64(6.42)})
		assert.Equal(t, arg.Ss, []string{"str1", "str2"})
		assert.Equal(t, arg.Tm, time.UnixMilli(1234567890123))
		return nil, nil
	}
	handlerType := reflect.TypeOf(handler)
	ctx := &Ctx{handlerReflectType: handlerType, request: req}
	_ = bindArgs()(ctx)
	args := ctx.args
	handlerValue := reflect.ValueOf(handler)
	handlerValue.Call(args)
}
