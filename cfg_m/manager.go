package cfg_m

import "errors"

var (
	configurationValueNotFoundMessage = "configuration value not found"
)

// ConfigStringGetter is responsible for allowing getting of loaded configuration from the implemented types.
type ConfigStringGetter interface {
	// Get returns the configuration value that has been previously loaded/
	// In event that the value isn't present it returns an error.
	Get(configurationName string) (string, error)
}

// ConfigStringGetterWithDefault is responsible for allowing getting of loaded configuration from the implemented types.
// With the addition of not returning an error, rather the provided default value
type ConfigStringGetterWithDefault interface {
	// GetOrDefault returns the configuration value that has been previously loaded/
	// In event that the value isn't present it returns the provided default value.
	GetOrDefault(configurationName string, defaultValue string) string
}

// Manager is used as a loaded configuration store.
// It provides a set of methods
type Manager struct {
	loadedConfiguration map[string]string
}

func (m *Manager) Get(configurationName string) (string, error) {
	return "", errors.New("not implemented")
}

func (m *Manager) GetOrDefault(configurationName string, defaultValue string) string {
	return ""
}
