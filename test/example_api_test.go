package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestApi(t *testing.T) {
	type EhcoReqDto struct {
		Id    int
		F32   float32
		F64   float64
		Email string
		Si    []int
		Sf32  []float32
		Sf64  []float64
		Ss    []string
		Tm    time.Time
	}
	arg := EhcoReqDto{
		Id:    1,
		F32:   3.2,
		F64:   6.4,
		Email: "email",
		Si:    []int{10, 100},
		Sf32:  []float32{3.21, 3.22},
		Sf64:  []float64{6.41, 6.42},
		Ss:    []string{"str1", "str2"},
		Tm:    time.UnixMilli(1234567890),
	}
	jb, _ := json.Marshal(arg)

	qs := `id=1&f32=3.2&f64=6.4&email=email&si=100&si=1000&sf32=3.21&sf32=3.22&sf64=6.41&sf64=6.42&ss=str1&ss=str2&tm=2009-02-14T07:31:30.123Z&v=1000000`

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
	formData.Add("tm", "2009-02-14T07:31:30.123Z")

	var mpf bytes.Buffer
	w := multipart.NewWriter(&mpf)
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
	fw.Write([]byte("2009-02-14T07:31:30.123Z"))
	w.Close()

	// 创建测试集
	type TestStruct struct {
		Methon  string
		Url     string
		Headers map[string]string
		Content interface{}
	}
	tests := []TestStruct{
		{"GET", "http://localhost:8000/example/api/v1/echoargs/echo_query", map[string]string{}, qs},
		{"POST", "http://localhost:8000/example/api/v1/echoargs/echo_form", map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, formData.Encode()},
		{"POST", "http://localhost:8000/example/api/v1/echoargs/echo_multipart_form", map[string]string{"Content-Type": w.FormDataContentType()}, mpf.String()},
		{"POST", "http://localhost:8000/example/api/v1/echoargs/echo_json", map[string]string{"Content-Type": "application/json"}, string(jb)},
		{"GET", "http://localhost:8000/example/api/v2/echoargs/echo_query", map[string]string{}, qs},
		{"POST", "http://localhost:8000/example/api/v2/echoargs/echo_form", map[string]string{"Content-Type": "application/x-www-form-urlencoded"}, formData.Encode()},
		{"POST", "http://localhost:8000/example/api/v2/echoargs/echo_multipart_form", map[string]string{"Content-Type": w.FormDataContentType()}, mpf.String()},
		{"POST", "http://localhost:8000/example/api/v2/echoargs/echo_json", map[string]string{"Content-Type": "application/json"}, string(jb)},
	}

	// 执行测试
	for _, testCase := range tests {
		switch testCase.Methon {
		case "GET":
			req, _ := http.NewRequest(testCase.Methon, fmt.Sprintf("%s?%s", testCase.Url, testCase.Content), nil)
			cl := http.Client{}
			rsp, err := cl.Do(req)
			assert.Nil(t, err)
			fmt.Printf("%s", rsp.Status)
		case "POST":
			req, _ := http.NewRequest(testCase.Methon, testCase.Url, strings.NewReader(testCase.Content.(string)))
			for k, v := range testCase.Headers {
				req.Header.Set(k, v)
			}
			cl := http.Client{}
			rsp, err := cl.Do(req)
			assert.Nil(t, err)
			fmt.Printf("%s", rsp.Status)
		}
	}
}
