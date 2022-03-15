package viperimp

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	vp, err := New("config")

	if err != nil {
		panic(fmt.Sprintf("failed to initrialize config component, err: %v", err))
	}
	assert.Equal(t, 8000, vp.GetIntOrDefault("server.port", 0))
}
