package cfg_e

import (
	"fmt"
	"github.com/pavleprica/configuration/conf_loader"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGetEnv(t *testing.T) {
	conf_loader.LoadEnvFile("../.env.example")
	t.Log("When using GetEnv")
	{
		t.Log("\tWhen variable present")
		{
			t.Log("\t\tShould return value")
			{
				variable, err := GetEnv("TEST_STR")

				assert.Nil(t, err)
				assert.Equal(t, "This is some cool str", variable)
			}
		}
		t.Log("\tWhen variable not present")
		{
			t.Logf("\t\tShould return '%s' error", notFoundErrorMessage)
			{
				variable, err := GetEnv("TEST_STR_NOT_PRESENT")

				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("TEST_STR_NOT_PRESENT - %s", notFoundErrorMessage), err.Error())
				assert.Equal(t, "", variable)
			}
		}
	}
}

func TestGetEnvOrDefault(t *testing.T) {
	conf_loader.LoadEnvFile("../.env.example")
	t.Log("When using GetEnvOrDefault")
	{
		t.Log("\tWhen variable present")
		{
			t.Log("\t\tShould return variable value")
			{
				variable := GetEnvOrDefault("TEST_STR", "DEFAULT_VALUE")

				assert.Equal(t, "This is some cool str", variable)
				assert.NotEqual(t, "DEFAULT_VALUE", variable)
			}
		}
		t.Log("\tWhen variable not present")
		{
			t.Log("\t\tShould return default value")
			{
				variable := GetEnvOrDefault("TEST_STR_NOT_PRESENT", "DEFAULT_VALUE")

				assert.NotEqual(t, "This is some cool str", variable)
				assert.Equal(t, "DEFAULT_VALUE", variable)
			}
		}
	}
}

func TestGetEnvInt(t *testing.T) {
	conf_loader.LoadEnvFile("../.env.example")
	t.Log("When using GetEnvInt")
	{
		t.Log("\tWhen variable present")
		{
			t.Log("\t\tWhen variable not parsable")
			{
				t.Logf("\t\t\tShould return '%s' error message", notParsableErrorMessage)
				{
					variable, err := GetEnvInt("TEST_INT_CORRUPTED")

					assert.NotNil(t, err)
					assert.True(t, strings.Contains(err.Error(), fmt.Sprintf("TEST_INT_CORRUPTED - %s", notParsableErrorMessage)), err.Error())
					assert.Equal(t, 0, variable)
				}
			}
			t.Log("\t\tWhen variable parsable")
			{
				t.Log("\t\t\tShould return variable")
				{
					variable, err := GetEnvInt("TEST_INT")

					assert.Nil(t, err)
					assert.Equal(t, 1234, variable)
				}
			}
		}
		t.Log("\tWhen variable not present")
		{
			t.Logf("\t\tShould return '%s' error", notFoundErrorMessage)
			{
				variable, err := GetEnvInt("TEST_INT_NOT_PRESENT")

				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("TEST_INT_NOT_PRESENT - %s", notFoundErrorMessage), err.Error())
				assert.Equal(t, 0, variable)
			}
		}
	}
}

func TestGetEnvIntOrDefault(t *testing.T) {
	conf_loader.LoadEnvFile("../.env.example")
	t.Log("When using GetEnvIntOrDefault")
	{
		t.Log("\tWhen variable present")
		{
			t.Log("\t\tWhen variable not parsable")
			{
				t.Log("\t\t\tShould return default value")
				{
					variable := GetEnvIntOrDefault("TEST_INT_CORRUPTED", 5)

					assert.Equal(t, 5, variable)
				}
			}
			t.Log("\t\tWhen variable parsable")
			{
				t.Log("\t\t\tShould return variable")
				{
					variable := GetEnvIntOrDefault("TEST_INT", 5)

					assert.Equal(t, 1234, variable)
				}
			}
		}
		t.Log("\tWhen variable not present")
		{
			t.Log("\t\t\tShould return default value")
			{
				variable := GetEnvIntOrDefault("TEST_INT_NOT_PRESENT", 5)

				assert.Equal(t, 5, variable)
			}
		}
	}
}

func TestGetEnvBool(t *testing.T) {
	conf_loader.LoadEnvFile("../.env.example")
	t.Log("When using GetEnvBool")
	{
		t.Log("\tWhen variable present")
		{
			t.Log("\t\tWhen variable not parsable")
			{
				t.Logf("\t\t\tShould return '%s' error message", notParsableErrorMessage)
				{
					variable, err := GetEnvBool("TEST_BOOL_CORRUPTED")

					assert.NotNil(t, err)
					assert.True(t, strings.Contains(err.Error(), fmt.Sprintf("TEST_BOOL_CORRUPTED - %s", notParsableErrorMessage)), err.Error())
					assert.Equal(t, false, variable)
				}
			}
			t.Log("\t\tWhen variable parsable")
			{
				t.Log("\t\t\tShould return variable")
				{
					variable, err := GetEnvBool("TEST_BOOL")

					assert.Nil(t, err)
					assert.Equal(t, true, variable)
				}
			}
		}
		t.Log("\tWhen variable not present")
		{
			t.Logf("\t\tShould return '%s' error", notFoundErrorMessage)
			{
				variable, err := GetEnvBool("TEST_BOOL_NOT_PRESENT")

				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("TEST_BOOL_NOT_PRESENT - %s", notFoundErrorMessage), err.Error())
				assert.Equal(t, false, variable)
			}
		}
	}
}

func TestGetEnvBoolOrDefault(t *testing.T) {
	conf_loader.LoadEnvFile("../.env.example")
	t.Log("When using GetEnvBoolOrDefault")
	{
		t.Log("\tWhen variable present")
		{
			t.Log("\t\tWhen variable not parsable")
			{
				t.Log("\t\t\tShould return default value")
				{
					variable := GetEnvBoolOrDefault("TEST_BOOL_CORRUPTED", true)

					assert.Equal(t, true, variable)
				}
			}
			t.Log("\t\tWhen variable parsable")
			{
				t.Log("\t\t\tShould return variable")
				{
					variable := GetEnvBoolOrDefault("TEST_BOOL", false)

					assert.Equal(t, true, variable)
				}
			}
		}
		t.Log("\tWhen variable not present")
		{
			t.Log("\t\t\tShould return default value")
			{
				variable := GetEnvBoolOrDefault("TEST_BOOL_NOT_PRESENT", true)

				assert.Equal(t, true, variable)
			}
		}
	}
}
