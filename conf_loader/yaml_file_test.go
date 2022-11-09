package conf_loader

import (
	"github.com/stretchr/testify/assert"
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
					assert.Fail(t, "not implemented")
				}
			}

			t.Log("\t\tWhen file present")
			{
				t.Log("\t\t\tShould return cfg_m.Manager with accessible data")
				{
					assert.Fail(t, "not implemented")
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
					assert.Fail(t, "not implemented")
				}
			}

			t.Log("\t\tWhen file present")
			{
				t.Log("\t\t\tShould return cfg_m.Manager with accessible data that are populated from system environment")
				{
					assert.Fail(t, "not implemented")
				}
			}
		}
	}
}
