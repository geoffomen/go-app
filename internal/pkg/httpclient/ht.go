package httpclient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/geoffomen/go-app/internal/pkg/myerr"
	"github.com/geoffomen/go-app/internal/pkg/mylog"
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
	requestDump, err := httputil.DumpRequest(req, true)
	if err != nil {
		mylog.Debugf("dump request occur error: %s", err)
	}
	rsp, err := cl.Do(req)
	if err != nil {
		mylog.Debugf("req: %s, err: %v ", requestDump, err)
		return nil, myerr.New(err)
	}
	if rsp.StatusCode != 200 {
		b, err := ioutil.ReadAll(rsp.Body)
		if err != nil {
			mylog.Debugf("req: %s, status_code: %v ", requestDump, rsp.StatusCode)
			return nil, myerr.Newf("status_code: %d", rsp.StatusCode)
		}
		return nil, myerr.Newf("status_code: %d, content: %s", rsp.StatusCode, b)
	}
	b, err := ioutil.ReadAll(rsp.Body)
	if err != nil {
		mylog.Debugf("req: %s, err: %v ", requestDump, err)
		return nil, myerr.New(err)
	}
	return b, nil
}

func Get(host string, params map[string]string, headers map[string]string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, host, nil)
	if err != nil {
		return nil, myerr.New(err)
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
		return nil, myerr.New(err)
	}
	req, err := http.NewRequest(http.MethodPost, host, bytes.NewBuffer(b))
	if err != nil {
		return nil, myerr.New(err)
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
		return nil, myerr.New(err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	return doRequest(req)
}

func PostMultiPart(host string, headers map[string]string, content map[string]io.Reader) ([]byte, error) {
	// Prepare a form that you will submit to that URL.
	var b bytes.Buffer
	w := multipart.NewWriter(&b)
	for key, r := range content {
		var fw io.Writer
		var err error
		if x, ok := r.(io.Closer); ok {
			defer x.Close()
		}
		// Add an image file
		if x, ok := r.(*os.File); ok {
			if fw, err = w.CreateFormFile(key, x.Name()); err != nil {
				return nil, myerr.New(err)
			}
		} else {
			// Add other fields
			if fw, err = w.CreateFormField(key); err != nil {
				return nil, myerr.New(err)
			}
		}
		if _, err = io.Copy(fw, r); err != nil {
			return nil, myerr.New(err)
		}

	}
	// Don't forget to close the multipart writer.
	// If you don't close it, your request will be missing the terminating boundary.
	w.Close()
	// Now that you have a form, you can submit it to your handler.
	req, err := http.NewRequest(http.MethodPost, host, &b)
	if err != nil {
		return nil, myerr.New(err)
	}
	for k, v := range headers {
		req.Header.Set(k, v)
	}
	// Don't forget to set the content type, this will contain the boundary.
	req.Header.Set("Content-Type", w.FormDataContentType())

	return doRequest(req)
}

func PostBinary(host string, headers map[string]string, content []byte) ([]byte, error) {
	req, err := http.NewRequest(http.MethodPost, host, bytes.NewBuffer(content))
	if err != nil {
		return nil, myerr.New(err)
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}
	req.Header.Set("Content-Type", "application/octet-stream")

	return doRequest(req)
}
