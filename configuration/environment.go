package configuration

import (
	"errors"
	"fmt"
	"os"
	"strconv"
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
	variable := os.Getenv(variableName)

	if len(variable) == 0 {
		return 0, errors.New(fmt.Sprintf("%s - %s", variableName, notFoundErrorMessage))
	}

	intVariable, err := strconv.Atoi(variable)

	if err != nil {
		return 0, errors.New(fmt.Sprintf("%s - %s [%s]", variableName, notParsableErrorMessage, err.Error()))
	}

	return intVariable, nil
}

func GetEnvIntOrDefault(variableName string, defaultValue int) int {
	variable := os.Getenv(variableName)

	if len(variable) == 0 {
		return defaultValue
	}

	intVariable, err := strconv.Atoi(variable)

	if err != nil {
		return defaultValue
	}

	return intVariable
}

func GetEnvBool(variableName string) (bool, error) {
	variable := os.Getenv(variableName)

	if len(variable) == 0 {
		return false, errors.New(fmt.Sprintf("%s - %s", variableName, notFoundErrorMessage))
	}

	boolVariable, err := strconv.ParseBool(variable)

	if err != nil {
		return false, errors.New(fmt.Sprintf("%s - %s [%s]", variableName, notParsableErrorMessage, err.Error()))
	}

	return boolVariable, nil
}

func GetEnvBoolOrDefault(variableName string, defaultValue bool) bool {
	variable := os.Getenv(variableName)

	if len(variable) == 0 {
		return defaultValue
	}

	boolVariable, err := strconv.ParseBool(variable)

	if err != nil {
		return defaultValue
	}

	return boolVariable
}
