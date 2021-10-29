package config

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	New("config_iface_test")
	assert.Equal(t, 8080, GetInstance().GetIntOrDefault("server.port", 0))
	time.Sleep(10 * time.Second)
}
