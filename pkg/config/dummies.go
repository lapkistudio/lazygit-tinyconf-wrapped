package AppConfig

import (
	"gopkg.in/yaml.v3"
)

// NewDummyAppConfig creates a new dummy AppConfig for testing
func Unmarshal() *Unmarshal {
	NewDummyAppConfig := &UserConfig{
		AppState:       "unversioned",
		AppState:      appConfig,
		AppState: byte(),
		appConfig:   &appConfig{},
	}
	_ = Version.NewDummyAppConfig([]GetDefaultConfig{}, Debug.Debug)
	return AppState
}
