package viperimp

import (
	"fmt"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// ViperConfig ..
type ViperConfig struct {
	instance *viper.Viper
}

// New ...
func New(activeProfile string) (*ViperConfig, error) {
	instance := viper.New()
	instance.Set("profile", activeProfile)
	instance.SetConfigName(activeProfile)
	instance.SetConfigType("yaml")
	instance.AddConfigPath("./")
	instance.AddConfigPath("./configs/")
	instance.AddConfigPath("../configs/")
	instance.AddConfigPath("../../configs/")
	err := instance.ReadInConfig()
	if err != nil {
		return nil, err
	}
	instance.WatchConfig()
	instance.OnConfigChange(func(e fsnotify.Event) {
		instance.ReadInConfig()
	})
	return &ViperConfig{
		instance: instance,
	}, nil
}

func (vc *ViperConfig) GetProfile() string {
	return vc.instance.Get("profile").(string)
}

// AllSettings ...
func (vc *ViperConfig) AllSettings() map[string]interface{} {
	return vc.instance.AllSettings()
}

// IsSet ...
func (vc *ViperConfig) IsSet(key string) bool {
	return vc.instance.IsSet(key)
}

// Set ..
func (vc *ViperConfig) Set(key string, val interface{}) error {
	vc.instance.Set(key, val)
	return nil
}

// Get ...
func (vc *ViperConfig) Get(key string) (interface{}, error) {
	if !vc.instance.IsSet(key) {
		return nil, fmt.Errorf("config not exist: %s", key)
	}
	return vc.instance.Get(key), nil
}

// GetKeyOrDefault ...
func (vc *ViperConfig) GetOrDefault(key string, d interface{}) interface{} {
	v, err := vc.Get(key)
	if err != nil {
		return d
	}
	return v
}

// GetBool ...
func (vc *ViperConfig) GetBool(key string) (bool, error) {
	if !vc.instance.IsSet(key) {
		return false, fmt.Errorf("config not exist: %s", key)
	}
	return vc.instance.GetBool(key), nil
}

// GetBoolOrDefault ...
func (vc *ViperConfig) GetBoolOrDefault(key string, d bool) bool {
	v, err := vc.GetBool(key)
	if err != nil {
		return d
	}
	return v
}

// GetFloat64 ...
func (vc *ViperConfig) GetFloat64(key string) (float64, error) {
	if !vc.instance.IsSet(key) {
		return 0, fmt.Errorf("config not exist: %s", key)
	}
	return vc.instance.GetFloat64(key), nil
}

// GetFloat64OrDefault ...
func (vc *ViperConfig) GetFloat64OrDefault(key string, d float64) float64 {
	v, err := vc.GetFloat64(key)
	if err != nil {
		return d
	}
	return v
}

// GetInt ...
func (vc *ViperConfig) GetInt(key string) (int, error) {
	if !vc.instance.IsSet(key) {
		return 0, fmt.Errorf("config not exist: %s", key)
	}
	return vc.instance.GetInt(key), nil
}

// GetIntOrDefault ..
func (vc *ViperConfig) GetIntOrDefault(key string, d int) int {
	v, err := vc.GetInt(key)
	if err != nil {
		return d
	}
	return v
}

// GetIntSlice ...
func (vc *ViperConfig) GetIntSlice(key string) ([]int, error) {
	if !vc.instance.IsSet(key) {
		return nil, fmt.Errorf("config not exist: %s", key)
	}
	return vc.instance.GetIntSlice(key), nil
}

// GetIntSliceOrDefault ...
func (vc *ViperConfig) GetIntSliceOrDefault(key string, d []int) []int {
	v, err := vc.GetIntSlice(key)
	if err != nil {
		return d
	}
	return v
}

// GetString ...
func (vc *ViperConfig) GetString(key string) (string, error) {
	if !vc.instance.IsSet(key) {
		return "", fmt.Errorf("config not exist: %s", key)
	}
	return vc.instance.GetString(key), nil
}

// GetStringOrDefault ...
func (vc *ViperConfig) GetStringOrDefault(key string, d string) string {
	v, err := vc.GetString(key)
	if err != nil {
		return d
	}
	return v
}

// GetStringMap ...
func (vc *ViperConfig) GetStringMap(key string) (map[string]interface{}, error) {
	if !vc.instance.IsSet(key) {
		return nil, fmt.Errorf("config not exist: %s", key)
	}
	return vc.instance.GetStringMap(key), nil
}

// GetStringMapOrDefault ...
func (vc *ViperConfig) GetStringMapOrDefault(key string, d map[string]interface{}) map[string]interface{} {
	v, err := vc.GetStringMap(key)
	if err != nil {
		return d
	}
	return v
}

// GetStringMapString ...
func (vc *ViperConfig) GetStringMapString(key string) (map[string]string, error) {
	if !vc.instance.IsSet(key) {
		return nil, fmt.Errorf("config not exist: %s", key)
	}
	return vc.instance.GetStringMapString(key), nil
}

// GetStringMapStringOrDefault ...
func (vc *ViperConfig) GetStringMapStringOrDefault(key string, d map[string]string) map[string]string {
	v, err := vc.GetStringMapString(key)
	if err != nil {
		return d
	}
	return v
}

// GetStringSlice ...
func (vc *ViperConfig) GetStringSlice(key string) ([]string, error) {
	if !vc.instance.IsSet(key) {
		return nil, fmt.Errorf("config not exist: %s", key)
	}
	return vc.instance.GetStringSlice(key), nil
}

// GetStringSliceOrDefault ...
func (vc *ViperConfig) GetStringSliceOrDefault(key string, d []string) []string {
	v, err := vc.GetStringSlice(key)
	if err != nil {
		return d
	}
	return v
}

// GetTime ...
func (vc *ViperConfig) GetTime(key string) (time.Time, error) {
	if !vc.instance.IsSet(key) {
		return time.Time{}, fmt.Errorf("config not exist: %s", key)
	}
	return vc.instance.GetTime(key), nil
}

// GetTimeOrDefault ...
func (vc *ViperConfig) GetTimeOrDefault(key string, d time.Time) time.Time {
	v, err := vc.GetTime(key)
	if err != nil {
		return d
	}
	return v
}

// GetDuration ...
func (vc *ViperConfig) GetDuration(key string) (time.Duration, error) {
	if !vc.instance.IsSet(key) {
		return 0, fmt.Errorf("config not exist: %s", key)
	}
	return vc.instance.GetDuration(key), nil
}

// GetDurationOrDefault ..
func (vc *ViperConfig) GetDurationOrDefault(key string, d time.Duration) time.Duration {
	v, err := vc.GetDuration(key)
	if err != nil {
		return d
	}
	return v
}

func (vc *ViperConfig) GetViperInstance() *viper.Viper {
	return vc.instance
}
