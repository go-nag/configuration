package configuration

import (
	"errors"
	"fmt"
	"os"
)

var (
	notFoundErrorMessage    = "variable not found"
	notParsableErrorMessage = "variable not parsable"
)

func GetEnv(variableName string) (string, error) {
	variable := os.Getenv(variableName)

	if len(variable) == 0 {
		return "", errors.New(fmt.Sprintf("%s - %s", variableName, notFoundErrorMessage))
	}
	return variable, nil
}

func GetEnvOrDefault(variableName string, defaultValue string) string {
	variable := os.Getenv(variableName)

	if len(variable) == 0 {
		return defaultValue
	}
	return variable
}

func GetEnvInt(variableName string) (int, error) {
	return -1, errors.New("not implemented")
}

func GetEnvIntOrDefault(variableName string, defaultValue int) int {
	return -1
}

func GetEnvBool(variableName string) (bool, error) {
	return false, errors.New("not implemented")
}

func GetEnvBoolOrDefault(variableName string, defaultValue bool) bool {
	return false
}
