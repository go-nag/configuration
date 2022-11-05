package configuration

import (
	"fmt"
	"github.com/pavleprica/configuration/configuration_loader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEnv(t *testing.T) {
	configuration_loader.LoadEnvFile("../.env.example")
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
	configuration_loader.LoadEnvFile("../.env.example")
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
	configuration_loader.LoadEnvFile("../.env.example")
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
					assert.Equal(t, fmt.Sprintf("TEST_INT_CORRUPTED - %s", notParsableErrorMessage), err.Error())
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
				variable, err := GetEnvInt("TEST_INT_NOT_RESENT")

				assert.NotNil(t, err)
				assert.Equal(t, fmt.Sprintf("TEST_INT_NOT_PRESENT - %s", notFoundErrorMessage), err.Error())
				assert.Equal(t, 0, variable)
			}
		}
	}
}
