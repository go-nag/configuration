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

func unmarshalYamlContent(objectBase string, yamlContent interface{}, configuration map[string]string) {
	v := reflect.ValueOf(yamlContent)

	for _, key := range v.MapKeys() {
		var base string
		if len(objectBase) == 0 {
			base = key.String()
		} else {
			base = fmt.Sprintf("%s.%s", objectBase, key)
		}
		value := v.MapIndex(key)

		if value.Kind() == reflect.Map {
			unmarshalYamlContent(base, value.Interface(), configuration)
		} else {
			configuration[base] = value.String()
		}
	}
}
