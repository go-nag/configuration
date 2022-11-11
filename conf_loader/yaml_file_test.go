package conf_loader

import (
	"github.com/pavleprica/configuration/cfg_m"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestLoadConfigFile(t *testing.T) {
	t.Log("When loading config-<environment>.yaml files")
	{
		t.Log("\tWhen loading config-local.yaml file")
		{
			environment := "local"
			t.Log("\t\tWhen file not present")
			{
				t.Logf("\t\t\tShould return %s for environment %s not found", configFileNotFound, environment)
				{
					manager, err := LoadConfigFile(environment + "wrong")

					assert.Nil(t, manager)
					assert.NotNil(t, err)
					assert.True(t, strings.Contains(err.Error(), configFileNotFound))
				}
			}

			t.Log("\t\tWhen file present")
			{
				t.Log("\t\t\tShould return cfg_m.Manager with accessible data")
				{
					manager, err := LoadConfigFile(environment)

					assert.Nil(t, err)
					assert.NotNil(t, manager)

					validateLocalConfigFileValues(t, manager)
				}
			}
		}
		t.Log("\tWhen loading config-dev.yaml file")
		{
			environment := "dev"
			t.Log("\t\tWhen file not present")
			{
				t.Logf("\t\t\tShould return %s for environment %s not found", configFileNotFound, environment)
				{
					manager, err := LoadConfigFile(environment + "wrong")

					assert.Nil(t, manager)
					assert.NotNil(t, err)
					assert.True(t, strings.Contains(err.Error(), configFileNotFound))
				}
			}

			t.Log("\t\tWhen file present")
			{
				LoadEnvFile("../config.dev.env")
				t.Log("\t\t\tShould return cfg_m.Manager with accessible data that are populated from system environment")
				{
					manager, err := LoadConfigFile(environment)

					assert.Nil(t, err)
					assert.NotNil(t, manager)

					validateDevConfigFileValues(t, manager)
				}
			}
		}
	}
}

func validateLocalConfigFileValues(t *testing.T, manager *cfg_m.Manager) {
	v, err := manager.Get("database.host")
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost:5042", v)

	v, err = manager.Get("database.username")
	assert.Nil(t, err)
	assert.Equal(t, "user", v)

	v, err = manager.Get("database.password")
	assert.Nil(t, err)
	assert.Equal(t, "my-secret-pw", v)

	v, err = manager.Get("kafka.url")
	assert.Nil(t, err)
	assert.Equal(t, "http://localhost:5555", v)

	v, err = manager.Get("kafka.clientId")
	assert.Nil(t, err)
	assert.Equal(t, "localApp", v)

	v, err = manager.Get("not.existing")
	assert.True(t, len(v) == 0)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "configuration value not found"))
}

func validateDevConfigFileValues(t *testing.T, manager *cfg_m.Manager) {
	v, err := manager.Get("database.host")
	assert.Nil(t, err)
	assert.Equal(t, "http://remote-database:5042", v)

	v, err = manager.Get("database.username")
	assert.Nil(t, err)
	assert.Equal(t, "database_username", v)

	v, err = manager.Get("database.password")
	assert.Nil(t, err)
	assert.Equal(t, "super-secret-password", v)

	v, err = manager.Get("kafka.url")
	assert.Nil(t, err)
	assert.Equal(t, "http://remote-kafka:5555", v)

	v, err = manager.Get("kafka.clientId")
	assert.Nil(t, err)
	assert.Equal(t, "dev_client", v)

	v, err = manager.Get("not.existing")
	assert.True(t, len(v) == 0)
	assert.NotNil(t, err)
	assert.True(t, strings.Contains(err.Error(), "configuration value not found"))
}
