package myHttpServerImp

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/url"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"ibingli.com/internal/pkg/myHttpServer"
)

type ReqStruct struct {
	Int       int
	IntPtr    *int
	F32       float32
	F32Ptr    *float32
	F64       float64
	F64Ptr    *float64
	String    string
	StringPtr *string
	Si        []int
	SiPtr     *[]int
	Sf32      []float32
	Sf32Ptr   *[]float32
	Sf64      []float64
	Sf64Ptr   *[]float64
	Ss        []string
	SsPtr     *[]string
	Tm        time.Time
	TmPtr     *time.Time
	Mt        myHttpServer.Mytime
	MtPtr     *myHttpServer.Mytime
	Struct    struct {
		StruInt int
	}
	StructPtr *struct {
		StruPtrInt int
	}
}

func TestBindJson(t *testing.T) {
	b := []byte(`{
		"Int": 1,
		"IntPtr": 2,
		"F32": 3.2,
		"F32Ptr": 3.29,
		"F64": 6.3,
		"F64Ptr": 6.33,
		"String": "string",
		"StringPtr": "pointString",
		"Si": [10,100],
		"SiPtr": [1000,10000],
		"Sf32": [3.21,3.22],
		"Sf32Ptr": [3.219,3.229],
		"Sf64": [6.11,6.12],
		"Sf64Ptr": [6.119,6.129],
		"Ss": ["str1","str2"],
		"SsPtr": ["str3","str5"],
		"Tm":"2000-01-01T00:00:00Z",
		"TmPtr":"1000-01-01T00:00:00Z",
		"Mt": 1234567890000,
		"MtPtr": 9876543210000,
		"Struct": {"StruInt": 1000000},
		"StructPtr": {"StruPtrInt": 9000000}}`)
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1", bytes.NewReader(b))
	req.Header.Add("Content-Type", "json")
	handler := func(arg ReqStruct) (interface{}, error) {
		assert.Equal(t, arg.Int, 1)
		assert.Equal(t, *arg.IntPtr, 2)
		assert.Equal(t, arg.F32, float32(3.2))
		assert.Equal(t, *arg.F32Ptr, float32(3.29))
		assert.Equal(t, arg.F64, float64(6.3))
		assert.Equal(t, *arg.F64Ptr, float64(6.33))
		assert.Equal(t, arg.String, "string")
		assert.Equal(t, *arg.StringPtr, "pointString")
		assert.Equal(t, arg.Si, []int{10, 100})
		assert.Equal(t, *arg.SiPtr, []int{1000, 10000})
		assert.Equal(t, arg.Sf32, []float32{float32(3.21), float32(3.22)})
		assert.Equal(t, *arg.Sf32Ptr, []float32{float32(3.219), float32(3.229)})
		assert.Equal(t, arg.Sf64, []float64{float64(6.11), float64(6.12)})
		assert.Equal(t, *arg.Sf64Ptr, []float64{float64(6.119), float64(6.129)})
		assert.Equal(t, arg.Ss, []string{"str1", "str2"})
		assert.Equal(t, *arg.SsPtr, []string{"str3", "str5"})
		tm, _ := time.Parse(time.DateTime, "2000-01-01 00:00:00")
		assert.Equal(t, arg.Tm, tm)
		tm2, _ := time.Parse(time.DateTime, "1000-01-01 00:00:00")
		assert.Equal(t, *arg.TmPtr, tm2)
		assert.Equal(t, arg.Mt, myHttpServer.Mytime(time.UnixMilli(1234567890000)))
		assert.Equal(t, *arg.MtPtr, myHttpServer.Mytime(time.UnixMilli(9876543210000)))
		assert.Equal(t, arg.Struct.StruInt, 1000000)
		assert.Equal(t, (*arg.StructPtr).StruPtrInt, 9000000)
		return nil, nil
	}
	handlerType := reflect.TypeOf(handler)
	ctx := &Ctx{handlerReflectType: handlerType, request: req}
	err := argBindingHandler()(ctx)
	assert.Nil(t, err)
	args := ctx.args
	handlerValue := reflect.ValueOf(handler)
	handlerValue.Call(args)
}

func TestBindQuery(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://127.0.0.1/?int=1&intPtr=2&f32=3.2&f32Ptr=3.29&f64=6.3&f64Ptr=6.33&string=string&stringPtr=pointString&si=10&si=100&siPtr=1000&siPtr=10000&sf32=3.21&sf32=3.22&sf32Ptr=3.219&sf32Ptr=3.229&sf64=6.11&sf64=6.12&sf64Ptr=6.119&sf64Ptr=6.129&ss=str1&ss=str2&ssPtr=str3&ssPtr=str5&tm=2000-01-01T00:00:00Z&tmPtr=1000-01-01T00:00:00Z&mt=1234567890000&mtPtr=9876543210000&struInt=1000000&struPtrInt=9000000", bytes.NewReader(nil))
	req.Header.Add("Content-Type", "")
	handler := func(arg ReqStruct) (interface{}, error) {
		assert.Equal(t, arg.Int, 1)
		assert.Equal(t, *arg.IntPtr, 2)
		assert.Equal(t, arg.F32, float32(3.2))
		assert.Equal(t, *arg.F32Ptr, float32(3.29))
		assert.Equal(t, arg.F64, float64(6.3))
		assert.Equal(t, *arg.F64Ptr, float64(6.33))
		assert.Equal(t, arg.String, "string")
		assert.Equal(t, *arg.StringPtr, "pointString")
		assert.Equal(t, arg.Si, []int{10, 100})
		assert.Equal(t, *arg.SiPtr, []int{1000, 10000})
		assert.Equal(t, arg.Sf32, []float32{float32(3.21), float32(3.22)})
		assert.Equal(t, *arg.Sf32Ptr, []float32{float32(3.219), float32(3.229)})
		assert.Equal(t, arg.Sf64, []float64{float64(6.11), float64(6.12)})
		assert.Equal(t, *arg.Sf64Ptr, []float64{float64(6.119), float64(6.129)})
		assert.Equal(t, arg.Ss, []string{"str1", "str2"})
		assert.Equal(t, *arg.SsPtr, []string{"str3", "str5"})
		tm, _ := time.Parse(time.DateTime, "2000-01-01 00:00:00")
		assert.Equal(t, arg.Tm, tm)
		tm2, _ := time.Parse(time.DateTime, "1000-01-01 00:00:00")
		assert.Equal(t, *arg.TmPtr, tm2)
		assert.Equal(t, arg.Mt, myHttpServer.Mytime(time.UnixMilli(1234567890000)))
		assert.Equal(t, *arg.MtPtr, myHttpServer.Mytime(time.UnixMilli(9876543210000)))
		assert.Equal(t, arg.Struct.StruInt, 1000000)
		assert.Equal(t, (*arg.StructPtr).StruPtrInt, 9000000)
		return nil, nil
	}
	handlerType := reflect.TypeOf(handler)
	ctx := &Ctx{handlerReflectType: handlerType, request: req}
	err := argBindingHandler()(ctx)
	assert.Nil(t, err)
	args := ctx.args
	handlerValue := reflect.ValueOf(handler)
	handlerValue.Call(args)
}

func TestBindFormUrlEncode(t *testing.T) {
	formData := url.Values{}
	formData.Add("int", "1")
	formData.Add("intPtr", "2")
	formData.Add("f32", "3.2")
	formData.Add("f32Ptr", "3.29")
	formData.Add("f64", "6.3")
	formData.Add("f64Ptr", "6.33")
	formData.Add("string", "string")
	formData.Add("stringPtr", "pointString")
	formData.Add("si", "10")
	formData.Add("si", "100")
	formData.Add("siPtr", "1000")
	formData.Add("siPtr", "10000")
	formData.Add("sf32", "3.21")
	formData.Add("sf32", "3.22")
	formData.Add("sf32Ptr", "3.219")
	formData.Add("sf32Ptr", "3.229")
	formData.Add("sf64", "6.11")
	formData.Add("sf64", "6.12")
	formData.Add("sf64Ptr", "6.119")
	formData.Add("sf64Ptr", "6.129")
	formData.Add("ss", "str1")
	formData.Add("ss", "str2")
	formData.Add("ssPtr", "str3")
	formData.Add("ssPtr", "str5")
	formData.Add("tm", "2000-01-01T00:00:00Z")
	formData.Add("tmPtr", "1000-01-01T00:00:00Z")
	formData.Add("mt", "1234567890000")
	formData.Add("mtPtr", "9876543210000")
	formData.Add("struInt", "1000000")
	formData.Add("struPtrInt", "9000000")
	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1/", strings.NewReader(formData.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	handler := func(arg ReqStruct) (interface{}, error) {
		assert.Equal(t, arg.Int, 1)
		assert.Equal(t, *arg.IntPtr, 2)
		assert.Equal(t, arg.F32, float32(3.2))
		assert.Equal(t, *arg.F32Ptr, float32(3.29))
		assert.Equal(t, arg.F64, float64(6.3))
		assert.Equal(t, *arg.F64Ptr, float64(6.33))
		assert.Equal(t, arg.String, "string")
		assert.Equal(t, *arg.StringPtr, "pointString")
		assert.Equal(t, arg.Si, []int{10, 100})
		assert.Equal(t, *arg.SiPtr, []int{1000, 10000})
		assert.Equal(t, arg.Sf32, []float32{float32(3.21), float32(3.22)})
		assert.Equal(t, *arg.Sf32Ptr, []float32{float32(3.219), float32(3.229)})
		assert.Equal(t, arg.Sf64, []float64{float64(6.11), float64(6.12)})
		assert.Equal(t, *arg.Sf64Ptr, []float64{float64(6.119), float64(6.129)})
		assert.Equal(t, arg.Ss, []string{"str1", "str2"})
		assert.Equal(t, *arg.SsPtr, []string{"str3", "str5"})
		tm, _ := time.Parse(time.DateTime, "2000-01-01 00:00:00")
		assert.Equal(t, arg.Tm, tm)
		tm2, _ := time.Parse(time.DateTime, "1000-01-01 00:00:00")
		assert.Equal(t, *arg.TmPtr, tm2)
		assert.Equal(t, arg.Mt, myHttpServer.Mytime(time.UnixMilli(1234567890000)))
		assert.Equal(t, *arg.MtPtr, myHttpServer.Mytime(time.UnixMilli(9876543210000)))
		assert.Equal(t, arg.Struct.StruInt, 1000000)
		assert.Equal(t, (*arg.StructPtr).StruPtrInt, 9000000)
		return nil, nil
	}
	handlerType := reflect.TypeOf(handler)
	ctx := &Ctx{handlerReflectType: handlerType, request: req}
	_ = argBindingHandler()(ctx)
	args := ctx.args
	handlerValue := reflect.ValueOf(handler)
	handlerValue.Call(args)
}

func TestBindMultipartForm(t *testing.T) {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	fw, _ := w.CreateFormField("int")
	fw.Write([]byte("1"))
	fw, _ = w.CreateFormField("intPtr")
	fw.Write([]byte("2"))
	fw, _ = w.CreateFormField("f32")
	fw.Write([]byte("3.2"))
	fw, _ = w.CreateFormField("f32Ptr")
	fw.Write([]byte("3.29"))
	fw, _ = w.CreateFormField("f64")
	fw.Write([]byte("6.3"))
	fw, _ = w.CreateFormField("f64Ptr")
	fw.Write([]byte("6.33"))
	fw, _ = w.CreateFormField("string")
	fw.Write([]byte("string"))
	fw, _ = w.CreateFormField("stringPtr")
	fw.Write([]byte("pointString"))
	fw, _ = w.CreateFormField("si")
	fw.Write([]byte("10"))
	fw, _ = w.CreateFormField("si")
	fw.Write([]byte("100"))
	fw, _ = w.CreateFormField("siPtr")
	fw.Write([]byte("1000"))
	fw, _ = w.CreateFormField("siPtr")
	fw.Write([]byte("10000"))
	fw, _ = w.CreateFormField("sf32")
	fw.Write([]byte("3.21"))
	fw, _ = w.CreateFormField("sf32")
	fw.Write([]byte("3.22"))
	fw, _ = w.CreateFormField("sf32Ptr")
	fw.Write([]byte("3.219"))
	fw, _ = w.CreateFormField("sf32Ptr")
	fw.Write([]byte("3.229"))
	fw, _ = w.CreateFormField("sf64")
	fw.Write([]byte("6.11"))
	fw, _ = w.CreateFormField("sf64")
	fw.Write([]byte("6.12"))
	fw, _ = w.CreateFormField("sf64Ptr")
	fw.Write([]byte("6.119"))
	fw, _ = w.CreateFormField("sf64Ptr")
	fw.Write([]byte("6.129"))
	fw, _ = w.CreateFormField("ss")
	fw.Write([]byte("str1"))
	fw, _ = w.CreateFormField("ss")
	fw.Write([]byte("str2"))
	fw, _ = w.CreateFormField("ssPtr")
	fw.Write([]byte("str3"))
	fw, _ = w.CreateFormField("ssPtr")
	fw.Write([]byte("str5"))
	fw, _ = w.CreateFormField("tm")
	fw.Write([]byte("2000-01-01T00:00:00Z"))
	fw, _ = w.CreateFormField("tmPtr")
	fw.Write([]byte("1000-01-01T00:00:00Z"))
	fw, _ = w.CreateFormField("mt")
	fw.Write([]byte("1234567890000"))
	fw, _ = w.CreateFormField("mtPtr")
	fw.Write([]byte("9876543210000"))
	fw, _ = w.CreateFormField("struInt")
	fw.Write([]byte("1000000"))
	fw, _ = w.CreateFormField("struPtrInt")
	fw.Write([]byte("9000000"))
	w.Close()

	req, _ := http.NewRequest(http.MethodPost, "http://127.0.0.1/", &body)
	req.Header.Add("Content-Type", w.FormDataContentType())
	handler := func(arg ReqStruct) (interface{}, error) {
		assert.Equal(t, arg.Int, 1)
		assert.Equal(t, *arg.IntPtr, 2)
		assert.Equal(t, arg.F32, float32(3.2))
		assert.Equal(t, *arg.F32Ptr, float32(3.29))
		assert.Equal(t, arg.F64, float64(6.3))
		assert.Equal(t, *arg.F64Ptr, float64(6.33))
		assert.Equal(t, arg.String, "string")
		assert.Equal(t, *arg.StringPtr, "pointString")
		assert.Equal(t, arg.Si, []int{10, 100})
		assert.Equal(t, *arg.SiPtr, []int{1000, 10000})
		assert.Equal(t, arg.Sf32, []float32{float32(3.21), float32(3.22)})
		assert.Equal(t, *arg.Sf32Ptr, []float32{float32(3.219), float32(3.229)})
		assert.Equal(t, arg.Sf64, []float64{float64(6.11), float64(6.12)})
		assert.Equal(t, *arg.Sf64Ptr, []float64{float64(6.119), float64(6.129)})
		assert.Equal(t, arg.Ss, []string{"str1", "str2"})
		assert.Equal(t, *arg.SsPtr, []string{"str3", "str5"})
		tm, _ := time.Parse(time.DateTime, "2000-01-01 00:00:00")
		assert.Equal(t, arg.Tm, tm)
		tm2, _ := time.Parse(time.DateTime, "1000-01-01 00:00:00")
		assert.Equal(t, *arg.TmPtr, tm2)
		assert.Equal(t, arg.Mt, myHttpServer.Mytime(time.UnixMilli(1234567890000)))
		assert.Equal(t, *arg.MtPtr, myHttpServer.Mytime(time.UnixMilli(9876543210000)))
		assert.Equal(t, arg.Struct.StruInt, 1000000)
		assert.Equal(t, (*arg.StructPtr).StruPtrInt, 9000000)
		return nil, nil
	}
	handlerType := reflect.TypeOf(handler)
	ctx := &Ctx{handlerReflectType: handlerType, request: req}
	_ = argBindingHandler()(ctx)
	args := ctx.args
	handlerValue := reflect.ValueOf(handler)
	handlerValue.Call(args)
}
