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

	configuration := make(map[string]*configValue)
	unmarshalYamlContent("", unmarshalledFileContent, configuration)
	populateConfigurationWithEnvironmentVariables(configuration)

	return newManager(configuration), nil
}

// LoadConfigFileWithPath will take the path `config-<environment>.yaml` file and
// provide the configuration manager allowing access to configuration data.
func LoadConfigFileWithPath(pathToConfigFile string) (*Manager, error) {
	return nil, errors.New("not implemented")
}

// LoadConfigFileWithDirScanning will try to locate the `config-<environment>.yaml` file and
// provide the configuration manager allowing access to configuration data.
// We provide it with the environment name, and the number of dirs to go upwards.
// In the event we are running a test file, we want to check one upward package as well.
// Example:
/*
project/
- config-local.yml
- package/
--  some_file.go
--  some_file_test.go
*/
// See issue https://github.com/go-nag/configuration/issues/6
func LoadConfigFileWithDirScanning(pathToConfigFile string) (*Manager, error) {
	return nil, errors.New("not implemented")
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

func unmarshalYamlContent(objectBase string, yamlContent interface{}, configuration map[string]*configValue) {
	switch value := yamlContent.(type) {
	// Single values
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
				configuration[base] = &configValue{
					value:   res,
					cfgType: str,
				}
			case int:
				configuration[base] = &configValue{
					value:   strconv.Itoa(res),
					cfgType: str,
				}
			case bool:
				configuration[base] = &configValue{
					value:   strconv.FormatBool(res),
					cfgType: str,
				}
			default:
				unmarshalYamlContent(base, v, configuration)
			}
		}
	// Array
	case []interface{}:
		var builder strings.Builder

		for _, v := range value {
			switch res := v.(type) {
			case string:
				builder.WriteString(res)
			case int:
				builder.WriteString(strconv.Itoa(res))
			case bool:
				builder.WriteString(strconv.FormatBool(res))
			default:
				fmt.Println("Unsupported type")
				fmt.Println(res)
			}

			builder.WriteString(";")
		}
		configuration[objectBase] = &configValue{
			value:   builder.String(),
			cfgType: arr,
		}
	default:
		fmt.Println("Unsupported type")
		fmt.Println(value)
	}
}

func populateConfigurationWithEnvironmentVariables(configuration map[string]*configValue) {
	for k, v := range configuration {
		if v.cfgType == str && strings.HasPrefix(v.value, "${") && strings.HasSuffix(v.value, "}") {
			envName := v.value[2 : len(v.value)-1]
			cfgValue := configuration[k]
			cfgValue.value = cfge.GetEnvOrDefault(envName, "")
		}
	}
}
