package mapimp

import (
	"fmt"
	"time"
)

type DefaultConfig struct {
	m map[string]interface{}
}

// New ...
func New() (*DefaultConfig, error) {
	return &DefaultConfig{
		m: make(map[string]interface{}),
	}, nil
}

func (dc *DefaultConfig) GetProfile() string {
	p, ok := dc.m["profile"]
	if !ok {
		return ""
	}
	return p.(string)
}

// AllSettings ...
func (dc *DefaultConfig) AllSettings() map[string]interface{} {
	return dc.m
}

// IsSet ...
func (dc *DefaultConfig) IsSet(key string) bool {
	_, ok := dc.m[key]
	return ok
}

// Set ..
func (dc *DefaultConfig) Set(key string, val interface{}) error {
	dc.m[key] = val
	return nil
}

// Get ...
func (dc *DefaultConfig) Get(key string) (interface{}, error) {
	v, ok := dc.m[key]
	if !ok {
		return nil, fmt.Errorf("config not exist: %s", key)
	}
	return v, nil
}

// GetKeyOrDefault ...
func (dc *DefaultConfig) GetOrDefault(key string, d interface{}) interface{} {
	v, err := dc.Get(key)
	if err != nil {
		return d
	}
	return v
}

// GetBool ...
func (dc *DefaultConfig) GetBool(key string) (bool, error) {
	v, ok := dc.m[key]
	if !ok {
		return false, fmt.Errorf("config not exist: %s", key)
	}
	nv, ok := v.(bool)
	if !ok {
		return false, fmt.Errorf("value not match type: %s", key)
	}
	return nv, nil
}

// GetBoolOrDefault ...
func (dc *DefaultConfig) GetBoolOrDefault(key string, d bool) bool {
	v, err := dc.GetBool(key)
	if err != nil {
		return d
	}
	return v
}

// GetFloat64 ...
func (dc *DefaultConfig) GetFloat64(key string) (float64, error) {
	v, ok := dc.m[key]
	if !ok {
		return 0, fmt.Errorf("config not exist: %s", key)
	}
	nv, ok := v.(float64)
	if !ok {
		return 0, fmt.Errorf("value not match type: %s", key)
	}
	return nv, nil
}

// GetFloat64OrDefault ...
func (dc *DefaultConfig) GetFloat64OrDefault(key string, d float64) float64 {
	v, err := dc.GetFloat64(key)
	if err != nil {
		return d
	}
	return v
}

// GetInt ...
func (dc *DefaultConfig) GetInt(key string) (int, error) {
	v, ok := dc.m[key]
	if !ok {
		return 0, fmt.Errorf("config not exist: %s", key)
	}
	nv, ok := v.(int)
	if !ok {
		return 0, fmt.Errorf("value not match type: %s", key)
	}
	return nv, nil
}

// GetIntOrDefault ..
func (dc *DefaultConfig) GetIntOrDefault(key string, d int) int {
	v, err := dc.GetInt(key)
	if err != nil {
		return d
	}
	return v
}

// GetIntSlice ...
func (dc *DefaultConfig) GetIntSlice(key string) ([]int, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetIntSliceOrDefault ...
func (dc *DefaultConfig) GetIntSliceOrDefault(key string, d []int) []int {
	v, err := dc.GetIntSlice(key)
	if err != nil {
		return d
	}
	return v
}

// GetString ...
func (dc *DefaultConfig) GetString(key string) (string, error) {
	v, ok := dc.m[key]
	if !ok {
		return "", fmt.Errorf("config not exist: %s", key)
	}
	nv, ok := v.(string)
	if !ok {
		return "", fmt.Errorf("value not match type: %s", key)
	}
	return nv, nil
}

// GetStringOrDefault ...
func (dc *DefaultConfig) GetStringOrDefault(key string, d string) string {
	v, err := dc.GetString(key)
	if err != nil {
		return d
	}
	return v
}

// GetStringMap ...
func (dc *DefaultConfig) GetStringMap(key string) (map[string]interface{}, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetStringMapOrDefault ...
func (dc *DefaultConfig) GetStringMapOrDefault(key string, d map[string]interface{}) map[string]interface{} {
	v, err := dc.GetStringMap(key)
	if err != nil {
		return d
	}
	return v
}

// GetStringMapString ...
func (dc *DefaultConfig) GetStringMapString(key string) (map[string]string, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetStringMapStringOrDefault ...
func (dc *DefaultConfig) GetStringMapStringOrDefault(key string, d map[string]string) map[string]string {
	v, err := dc.GetStringMapString(key)
	if err != nil {
		return d
	}
	return v
}

// GetStringSlice ...
func (dc *DefaultConfig) GetStringSlice(key string) ([]string, error) {
	return nil, fmt.Errorf("not implemented")
}

// GetStringSliceOrDefault ...
func (dc *DefaultConfig) GetStringSliceOrDefault(key string, d []string) []string {
	v, err := dc.GetStringSlice(key)
	if err != nil {
		return d
	}
	return v
}

// GetTime ...
func (dc *DefaultConfig) GetTime(key string) (time.Time, error) {
	return time.Now(), fmt.Errorf("not implemented")
}

// GetTimeOrDefault ...
func (dc *DefaultConfig) GetTimeOrDefault(key string, d time.Time) time.Time {
	v, err := dc.GetTime(key)
	if err != nil {
		return d
	}
	return v
}

// GetDuration ...
func (dc *DefaultConfig) GetDuration(key string) (time.Duration, error) {
	return time.Duration(1), fmt.Errorf("not implemented")
}

// GetDurationOrDefault ..
func (dc *DefaultConfig) GetDurationOrDefault(key string, d time.Duration) time.Duration {
	v, err := dc.GetDuration(key)
	if err != nil {
		return d
	}
	return v
}
