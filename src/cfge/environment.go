package cfge

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

// GetEnv returns a string value of the environment variable name provided.
// It returns an error in the event of the variable not existing.
func GetEnv(variableName string) (string, error) {
	variable := os.Getenv(variableName)

	if len(variable) == 0 {
		return "", errors.New(fmt.Sprintf("%s - %s", variableName, notFoundErrorMessage))
	}
	return variable, nil
}

// GetEnvOrDefault returns a string value of the environment variable name provided.
// In case there is an error, or the variable doesn't exist, it will return the default provided value.
func GetEnvOrDefault(variableName string, defaultValue string) string {
	variable := os.Getenv(variableName)

	if len(variable) == 0 {
		return defaultValue
	}
	return variable
}

// GetEnvInt returns an int value of the environment variable name provided.
// It returns an error in the event of the variable not existing or parsing error.
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

// GetEnvIntOrDefault returns an int value of the environment variable name provided.
// In case there is an error, or the variable doesn't exist, it will return the default provided value.
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

// GetEnvBool returns a bool value of the environment variable name provided.
// It returns an error in the event of the variable not existing or parsing error.
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

// GetEnvBoolOrDefault returns a bool value of the environment variable name provided.
// In case there is an error, or the variable doesn't exist, it will return the default provided value.
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
