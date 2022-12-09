package cfgm

import (
	"errors"
	"fmt"
)

// error list
var (
	configurationValueNotFoundMessage = "configuration value not found"
	invalidValueType                  = "invalid value type"
)

// ConfigStringGetter is responsible for allowing getting of loaded configuration from the implemented types.
type ConfigStringGetter interface {
	// Get returns the configuration value that has been previously loaded.
	// In event that the value isn't present it returns an error.
	Get(configurationName string) (string, error)
}

// ConfigArrayGetter is responsible for parsing and getting config values that are represented as
type ConfigArrayGetter interface {
	// GetArr returns the configuration in an array form that has been previously loaded.
	// In event that the value isn't present it returns an empty array.
	GetArr(configurationName string) []string
}

// ConfigStringGetterWithDefault is responsible for allowing getting of loaded configuration from the implemented types.
// With the addition of not returning an error, rather the provided default value
type ConfigStringGetterWithDefault interface {
	// GetOrDefault returns the configuration value that has been previously loaded/
	// In event that the value isn't present it returns the provided default value.
	GetOrDefault(configurationName string, defaultValue string) string
}

// Represents currently supported types by the configuration
var (
	str configType = 0
	arr configType = 1
)

// configType represents the type of the configuration value
// see configValue
type configType = int

type configValue struct {
	value   string
	cfgType configType
}

// Manager is used as a loaded configuration store.
// It provides a set of methods
type Manager struct {
	loadedConfiguration map[string]configValue
}

func newManager(configuration map[string]configValue) *Manager {
	return &Manager{
		loadedConfiguration: configuration,
	}
}

func (m *Manager) Get(configurationName string) (string, error) {
	cfgValue, present := m.loadedConfiguration[configurationName]
	if present && cfgValue.cfgType == str {
		return cfgValue.value, nil
	} else if cfgValue.cfgType == str {
		return "", errors.New(fmt.Sprintf("%s - %s", configurationName, configurationValueNotFoundMessage))
	} else {
		return "", errors.New(invalidValueType)
	}
}

func (m *Manager) GetOrDefault(configurationName string, defaultValue string) string {
	cfgValue, present := m.loadedConfiguration[configurationName]
	if present && cfgValue.cfgType == str {
		return cfgValue.value
	} else {
		return defaultValue
	}
}
