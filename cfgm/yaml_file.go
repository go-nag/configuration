package cfgm

import (
	"errors"
	"flag"
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

	return loadConfigFile(err, fileContent)
}

// LoadConfigFileWithPath will take the path `config-<environment>.yaml` file and
// provide the configuration manager allowing access to configuration data.
func LoadConfigFileWithPath(pathToConfigFile string) (*Manager, error) {
	fmt.Printf("Loading %s\n", pathToConfigFile)

	fileContent, err := readFile(pathToConfigFile)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, errors.New(fmt.Sprintf("%s - %s", configFileNotFound, pathToConfigFile))
		}
		return nil, err
	}

	return loadConfigFile(err, fileContent)
}

func loadConfigFile(err error, fileContent []byte) (*Manager, error) {
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
	if isInTests() {
		return filepath.Join(workDir, "..", fmt.Sprintf("config-%s.yaml", environment))
	} else {
		return filepath.Join(workDir, fmt.Sprintf("config-%s.yaml", environment))
	}
}

func isInTests() bool {
	return flag.Lookup("test.v") != nil
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
