package digestutil

import (
	"crypto/md5"
	"encoding/hex"
	"strings"

	uuid "github.com/satori/go.uuid"
)

// 生成盐
func GenUuid() string {
	salt := uuid.NewV4().String()
	return strings.ReplaceAll(salt, "-", "")
}

// 盐 + 明文 得到密文
func Md5Encryption(value string, salt string) string {
	m5 := md5.New()
	m5.Write([]byte(value))
	m5.Write([]byte(salt))
	st := m5.Sum(nil)
	return hex.EncodeToString(st)
}
