package digestutil

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// GenUuid ..
func GenUuid() string {
	u := uuid.NewV4().String()
	return strings.ReplaceAll(u, "-", "")
}

// Md5Encryption ..
func Md5Encryption(value string, salt string) string {
	m5 := md5.New()
	m5.Write([]byte(value))
	m5.Write([]byte(salt))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}

func Md5SumBinary(value []byte) string {
	m5 := md5.New()
	m5.Write(value)
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}

func Md5SumFile(filePath string) (string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("无法打开文件, file: %s, error: %s", filePath, err)
	}
	content, err := ioutil.ReadAll(f)
	if err != nil {
		return "", fmt.Errorf("无法读取文件, file: %s, error: %s", filePath, err)
	}
	m5 := md5.New()
	m5.Write(content)
	st := m5.Sum(nil)
	return hex.EncodeToString(st), nil
}
