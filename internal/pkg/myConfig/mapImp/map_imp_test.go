package mapImp

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestXxx(t *testing.T) {
	c, _ := New()
	c.Set("k1", 1)
	c.Set("foo", "bar")
	c.Set("time", time.Now())

	assert.Equal(t, true, c.IsSet("k1"))
	assert.Equal(t, false, c.IsSet("k2"))

	v, err := c.Get("foo")
	assert.Equal(t, "bar", v.(string))
	assert.Nil(t, err)

	v, err = c.Get("notExist")
	assert.Equal(t, nil, v)
	assert.NotNil(t, err)

	fmt.Println(c.AllSettings())
}
