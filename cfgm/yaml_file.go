package cfgm

import (
	"errors"
	"fmt"
	"github.com/go-nag/configuration/cfge"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var (
	configFileNotFound = "configuration file not found"
	configTestRunKey   = "CONFIGURATION_TEST_RUN"
)

// LoadConfigFile will take the `config-<environment>.yaml` file and
// provide the configuration manager allowing access to configuration data.
func LoadConfigFile(environment string) (*Manager, error) {
	fmt.Printf("Loading config-%s.yaml\n", environment)

	workDir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	configFilePath := getConfigFilePath(environment, workDir)
	fileContent, err := readFile(configFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New(fmt.Sprintf("%s - %s", configFileNotFound, configFilePath))
		}
		return nil, err
	}

	var unmarshalledFileContent interface{}
	err = yaml.Unmarshal(fileContent, &unmarshalledFileContent)
	if err != nil {
		return nil, err
	}

	configuration := make(map[string]string)
	unmarshalYamlContent("", unmarshalledFileContent, configuration)
	populateConfigurationWithEnvironmentVariables(configuration)

	return newManager(configuration), nil
}

func readFile(configFilePath string) ([]byte, error) {
	fileContent, err := os.ReadFile(configFilePath)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) && strings.HasSuffix(configFilePath, ".yaml") {
			alternativeYamlExtensionPath := configFilePath[:strings.Index(configFilePath, ".yaml")] + ".yml"
			return readFile(alternativeYamlExtensionPath)
		}
		return nil, err
	}
	return fileContent, nil
}

func getConfigFilePath(environment string, workDir string) string {
	if cfge.GetEnvBoolOrDefault(configTestRunKey, false) {
		return filepath.Join(workDir, "..", fmt.Sprintf("config-%s.yaml", environment))
	} else {
		return filepath.Join(workDir, fmt.Sprintf("config-%s.yaml", environment))
	}
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

			switch res := v.(type) {
			case string:
				configuration[base] = res
			case int:
				configuration[base] = strconv.Itoa(res)
			case bool:
				configuration[base] = strconv.FormatBool(res)
			default:
				unmarshalYamlContent(base, v, configuration)
			}
		}
	default:
		fmt.Println("Unsupported type")
		fmt.Println(value)
	}
}

func populateConfigurationWithEnvironmentVariables(configuration map[string]string) {
	for k, v := range configuration {
		if strings.HasPrefix(v, "${") && strings.HasSuffix(v, "}") {
			envName := v[2 : len(v)-1]
			configuration[k] = cfge.GetEnvOrDefault(envName, "")
		}
	}
}
