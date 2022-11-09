package conf_loader

import (
	"errors"
	"github.com/pavleprica/configuration/cfg_m"
)

var (
	configFileNotFound = "configuration file not found"
)

// LoadConfigFile will take the `config-<environment>.yaml` file and
// provide the configuration manager allowing access to configuration data.
func LoadConfigFile(environment string) (*cfg_m.Manager, error) {
	return nil, errors.New("not implemented")
}
