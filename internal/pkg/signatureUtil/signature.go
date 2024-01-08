package signatureUtil

import (
	"crypto/hmac"
	"crypto/sha1"
	"fmt"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"strings"
)

var NeedSignHeaders = map[string]bool{
	"range":        true,
	"content-type": true,
	"content-md5":  true,
}

func New(secretID string, secretKey string, req *http.Request, beginTmMili int64, endTmMili int64) (string, error) {
	keyTm := fmt.Sprintf("%d;%d", beginTmMili, endTmMili)
	signKey := calSignKey(secretKey, keyTm)

	queryParams, queryParamKeys := genQueryParamString(req)
	headerParams, headerParamKeys := genHeaderParamString(req)
	httpString := fmt.Sprintf("%s\n%s\n%s\n%s\n", strings.ToLower(req.Method), req.URL.Path, queryParams, headerParams)

	h := sha1.New()
	h.Write([]byte(httpString))
	stringToSign := fmt.Sprintf("%s\n%s\n%x\n", "sha1", keyTm, h.Sum(nil))

	h = hmac.New(sha1.New, []byte(signKey))
	h.Write([]byte(stringToSign))
	signature := fmt.Sprintf("%x", h.Sum(nil))

	return strings.Join([]string{
		"b-sign-algorithm=" + "sha1",
		"b-ak=" + secretID,
		"b-sign-time=" + keyTm,
		"b-header-list=" + strings.Join(headerParamKeys, ";"),
		"b-query-param-list=" + strings.Join(queryParamKeys, ";"),
		"b-signature=" + signature,
	}, "&"), nil
}

func IsSignatureValid(secretID string, secretKey string, req *http.Request, signature string) bool {
	parts := strings.Split(signature, "&")
	if len(parts) < 3 || !strings.HasPrefix(parts[2], "b-sign-time=") {
		return false
	}
	sgParts := strings.Split(parts[2], "=")
	if len(sgParts) < 2 {
		return false
	}
	tmParts := strings.Split(sgParts[1], ";")
	if len(tmParts) < 2 {
		return false
	}
	bgTm, err := strconv.ParseInt(tmParts[0], 10, 64)
	if err != nil {
		return false
	}
	endTm, err := strconv.ParseInt(tmParts[1], 10, 64)
	if err != nil {
		return false
	}
	rt, err := New(secretID, secretKey, req, bgTm, endTm)
	if err != nil {
		return false
	}
	return rt == signature
}

func GetSignatureExpireTimeMili(signature string) (int64, error) {
	parts := strings.Split(signature, "&")
	if len(parts) < 3 || !strings.HasPrefix(parts[2], "b-sign-time=") {
		return 0, fmt.Errorf("format error")
	}
	sgParts := strings.Split(parts[2], "=")
	if len(sgParts) < 2 {
		return 0, fmt.Errorf("format error")
	}
	tmParts := strings.Split(sgParts[1], ";")
	if len(tmParts) < 2 {
		return 0, fmt.Errorf("format error")
	}
	endTm, err := strconv.ParseInt(tmParts[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("format error")
	}
	if len(tmParts[1]) < 13 {
		return endTm * 1000, nil
	}
	return endTm, nil
}

func GetSignatureSecretId(signature string) (string, error) {
	parts := strings.Split(signature, "&")
	if len(parts) < 2 || !strings.HasPrefix(parts[1], "b-ak=") {
		return "", fmt.Errorf("format error")
	}
	akParts := strings.Split(parts[1], "=")
	if len(akParts) < 2 {
		return "", fmt.Errorf("format error")
	}

	return akParts[1], nil
}

func calSignKey(key string, keyTm string) string {
	h := hmac.New(sha1.New, []byte(key))
	h.Write([]byte(keyTm))
	return fmt.Sprintf("%x", h.Sum(nil))
}

func genQueryParamString(req *http.Request) (queryParams string, queryParamKeyList []string) {
	m := make(map[string][]string)
	ks := make([]string, 0)
	for k, v := range req.URL.Query() {
		ks = append(ks, k)
		m[k] = make([]string, 0, len(v))
		m[k] = append(m[k], v...)
	}
	sort.Strings(ks)

	var pairs []string
	for _, k := range ks {
		nk := strings.ToLower(url.QueryEscape(k))
		queryParamKeyList = append(queryParamKeyList, nk)
		item := m[k]
		sort.Strings(item)
		for _, v := range item {
			pairs = append(pairs, fmt.Sprintf("%s=%s", nk, url.QueryEscape(v)))
		}
	}
	queryParams = strings.Join(pairs, "&")

	return queryParams, queryParamKeyList
}

func genHeaderParamString(req *http.Request) (headerParams string, headerParamKeyList []string) {
	m := make(map[string][]string)
	ks := make([]string, 0)
	for k, v := range req.Header {
		ks = append(ks, k)
		m[k] = v
	}
	sort.Strings(ks)

	var pairs []string
	for _, k := range ks {
		nk := strings.ToLower(url.QueryEscape(k))
		if !NeedSignHeaders[nk] {
			continue
		}
		headerParamKeyList = append(headerParamKeyList, nk)
		item := m[k]
		sort.Strings(item)
		for _, v := range item {
			pairs = append(pairs, fmt.Sprintf("%s=%s", nk, url.QueryEscape(v)))
		}
	}
	headerParams = strings.Join(pairs, "&")
	return headerParams, headerParamKeyList
}
