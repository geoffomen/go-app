package myconfig

import (
	"time"
)

type MyConfigIface interface {
	GetProfile() string
	AllSettings() map[string]interface{}
	IsSet(key string) bool
	Set(key string, val interface{}) error
	Get(key string) (interface{}, error)
	GetOrDefault(key string, d interface{}) interface{}
	GetBool(key string) (bool, error)
	GetBoolOrDefault(key string, d bool) bool
	GetFloat64(key string) (float64, error)
	GetFloat64OrDefault(key string, d float64) float64
	GetInt(key string) (int, error)
	GetIntOrDefault(key string, d int) int
	GetIntSlice(key string) ([]int, error)
	GetIntSliceOrDefault(key string, d []int) []int
	GetString(key string) (string, error)
	GetStringOrDefault(key string, d string) string
	GetStringMap(key string) (map[string]interface{}, error)
	GetStringMapOrDefault(key string, d map[string]interface{}) map[string]interface{}
	GetStringMapString(key string) (map[string]string, error)
	GetStringMapStringOrDefault(key string, d map[string]string) map[string]string
	GetStringSlice(key string) ([]string, error)
	GetStringSliceOrDefault(key string, d []string) []string
	GetTime(key string) (time.Time, error)
	GetTimeOrDefault(key string, d time.Time) time.Time
	GetDuration(key string) (time.Duration, error)
	GetDurationOrDefault(key string, d time.Duration) time.Duration
}
