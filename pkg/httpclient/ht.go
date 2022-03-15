package httpclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
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

func doRequest(req *http.Request) ([]byte, error) {
	requestDump, _ := httputil.DumpRequest(req, true)

	rsp, err := cl.Do(req)
	if err != nil {
		return nil, fmt.Errorf("req: %s, err: %s", requestDump, err)
	}
	if rsp.StatusCode != 200 {
		b, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			return nil, fmt.Errorf("status_code: %d, 读取响应信息失败: %s", rsp.StatusCode, err)
		}
		return nil, fmt.Errorf("status_code: %d, content: %s", rsp.StatusCode, b)
	}
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应信息失败，req: %s, err: %s ", requestDump, err)
	}
	return b, nil
}

func Get(host string, params map[string]string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, host, nil)
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	var q url.Values
	for k, v := range params {
		q = req.URL.Query()
		q.Add(k, v)
	}
	req.URL.RawQuery = q.Encode()
	return doRequest(req)
}

func PostJson(host string, headers map[string]string, content map[string]string) ([]byte, error) {
	b, err := json.Marshal(content)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(http.MethodPost, host, bytes.NewBuffer(b))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/json")

	return doRequest(req)
}

func PostForm(host string, headers map[string]string, content map[string]string) ([]byte, error) {
	var q url.Values
	for k, v := range content {
		q.Add(k, v)
	}
	req, err := http.NewRequest(http.MethodPost, host, bytes.NewBuffer([]byte(q.Encode())))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return doRequest(req)
}

func PostMultiPart(host string, headers map[string]string, content map[string]io.Reader) ([]byte, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
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
	req, err := http.NewRequest(http.MethodPost, host, &b)
	if err != nil {
		return nil, err
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", w.FormDataContentType())

	return doRequest(req)
}

func PostBinary(host string, headers map[string]string, content []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, host, bytes.NewBuffer(content))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	return doRequest(req)
}
