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
	switch value := yamlContent.(type) {
	case map[string]interface{}:
		for k, v := range value {
			var base string
			if len(objectBase) == 0 {
				base = k
			} else {
				base = fmt.Sprintf("%s.%s", objectBase, k)
			}

			strValue, ok := v.(string)

			if ok {
				configuration[base] = strValue
			} else {
				unmarshalYamlContent(base, v, configuration)
			}
		}
	default:
		fmt.Println("Unsupported type")
		fmt.Println(value)
	}
}
