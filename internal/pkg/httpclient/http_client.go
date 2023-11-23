package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"
)

var cl http.Client = http.Client{
	Timeout: time.Second * 60,
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 30 * time.Second,
	},
}

func DoRequest(req *http.Request) ([]byte, error) {
	requestDump, _ := httputil.DumpRequest(req, true)

	rsp, err := cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("req: %s, err: %s", requestDump, err)
	}
	if rsp.StatusCode != http.StatusOK {
		body, err := io.ReadAll(rsp.Body)
		if err != nil {
			return nil, fmt.Errorf("status_code: %d, 读取响应信息失败: %s", rsp.StatusCode, err)
		}
		return nil, fmt.Errorf("status_code: %d, content: %s", rsp.StatusCode, body)
	}
	body, err := io.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应信息失败，req: %s, err: %s ", requestDump, err)
	}
	return body, nil
}

func Get(targetUrl string, params map[string]interface{}, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, targetUrl, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	q := req.URL.Query()
	for k, v := range params {
		s, ok := v.(string)
		if ok {
			q.Add(k, s)
		} else {
			b, _ := json.Marshal(v)
			q.Add(k, string(b))
		}
	}
	req.URL.RawQuery = q.Encode()
	return DoRequest(req)
}

func PostJson(targetUrl string, headers map[string]string, content []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, targetUrl, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	return DoRequest(req)
}

func PostForm(targetUrl string, headers map[string]string, content map[string]string) ([]byte, error) {
	var body url.Values
	for k, v := range content {
		body.Add(k, v)
	}
	req, err := http.NewRequest(http.MethodPost, targetUrl, bytes.NewBuffer([]byte(body.Encode())))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return DoRequest(req)
}

func PostMultiPart(targetUrl string, headers map[string]string, content map[string]io.Reader) ([]byte, error) {
	var body bytes.Buffer
	w := multipart.NewWriter(&body)
	for key, r := range content {
		var fw io.Writer
		var err error
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, err
			}
		} else {
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, err
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, err
		}

	}
	w.Close()
	req, err := http.NewRequest(http.MethodPost, targetUrl, &body)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return DoRequest(req)
}

func PostBinary(targetUrl string, headers map[string]string, content []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, targetUrl, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	return DoRequest(req)
}
