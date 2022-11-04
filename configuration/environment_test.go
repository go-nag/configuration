package configuration

import (
	"fmt"
	"github.com/pavleprica/configuration/loader"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetEnv(t *testing.T) {
	loader.LoadEnvFile("../.env.example")
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
	loader.LoadEnvFile("../.env.example")
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
