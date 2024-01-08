package viperImp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	vp, err := New("test_config")

	if err != nil {
		panic(fmt.Sprintf("failed to initrialize config component, err: %v", err))
	}
	assert.Equal(t, 8000, vp.GetIntOrDefault("server.port", 0))
	assert.Equal(t, "default", vp.GetStringOrDefault("server.notExist", "default"))
	assert.Equal(t, "default", vp.GetStringOrDefault("not.exist", "default"))
}
