package signatureutil

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSignatureGen(t *testing.T) {
	secretID := "957f9493d"
	secretKey := "3242827310774e2e9ff6ab354666cf12"

	uri := "http://example.com/miis/openapi/v1/slice/page"
	now := time.Now()
	startTime := now.Unix()
	endTime := now.Add(time.Hour * 24*30*12*99).Unix()

	content := `{"idBegin": 0}`
	b, err := json.Marshal(content)
	assert.Nil(t, err)
	bl := len(b)

	req, _ := http.NewRequest("POST", uri, bytes.NewBuffer(b))
	assert.Equal(t, int64(bl), req.ContentLength)
	req.Header.Add("Content-Type", "application/json")

	signature, err := New(secretID, secretKey, req, startTime, endTime)
	fmt.Println(signature)
	assert.Nil(t, err)
}

var wg *sync.WaitGroup
var rt bool

func handler(w http.ResponseWriter, r *http.Request) {
	defer wg.Done()
	expectedSignature := r.Header.Get("Authorization")
	secretID := "957f9493d"
	secretKey := "3242827310774e2e9ff6ab354666cf12"
	rt = IsSignatureValid(secretID, secretKey, r, expectedSignature)
	fmt.Println(rt)
}

func TestSignatureValid(t *testing.T) {
	wg = new(sync.WaitGroup)
	wg.Add(1)
	http.HandleFunc("/api/v1/res", handler)
	go http.ListenAndServe("0.0.0.0:9090", nil)

	cmd := exec.Command("curl", "-X", "POST",
		"-H", "Content-Type: application/json",
		"-H", "content-md5: md5",
		"-H", "Content-Length: 12",
		"-H", "Range: bytes=0-",
		"-H", "origin: 127.0.0.1",
		"-H", "OtherHeader: OtherHeaderValue",
		"-H", "Authorization: b-sign-algorithm=sha1&b-ak=957f9493d&b-sign-time=1692170326;1692177526&b-header-list=content-md5;content-type;range&b-query-param-list=another-empty-key;k1;k2&b-signature=d3f3c2f7e2377223f8a4abcc7f0c84e8db52e677",
		"-d", "{\"key\": \"v\"}",
		"http://127.0.0.1:9090/api/v1/res?k1=1&k1=2&k2=v&another-empty-key",
	)
	cmd.Run()

	wg.Wait()
	assert.Equal(t, rt, true)
}
