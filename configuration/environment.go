package configuration

import (
	"errors"
)

var (
	notFoundErrorMessage string = "variable not found"
)

func GetEnv(variableName string) (string, error) {
	return "", errors.New("not implemented")
}

func GetEnvOrDefault(variableName string, defaultValue string) string {
	return "not implemented"
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
