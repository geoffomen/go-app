package httpClientUtil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet(t *testing.T) {
	hc := New(newLogger(), LogLevelDebug)
	bs, err := hc.Get("https://baidu.com", nil, nil)
	assert.Nil(t, err)
	fmt.Println(string(bs))
}
