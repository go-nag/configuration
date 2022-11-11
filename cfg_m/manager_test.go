package cfg_m

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestManager_Get(t *testing.T) {
	var manager ConfigStringGetter
	t.Log("\tWhen using Get")
	{
		t.Log("\t\tWhen value not present")
		{
			manager = &Manager{
				make(map[string]string),
			}
			t.Logf("\t\t\tShould return %s error", configurationValueNotFoundMessage)
			{
				cfgValue, err := manager.Get("test.value")

				assert.NotNil(t, err)
				assert.Equal(t, "", cfgValue)
				assert.True(t, strings.Contains(err.Error(), configurationValueNotFoundMessage))
			}
		}
		t.Log("\t\tWhen value present")
		{
			testMap := make(map[string]string)
			testMap["test.value"] = "some_value"
			manager = &Manager{
				testMap,
			}
			t.Log("\t\t\tShould return value")
			{
				cfgValue, err := manager.Get("test.value")
				assert.Nil(t, err)
				assert.Equal(t, "some_value", cfgValue)
			}
		}
	}
}

func TestManager_GetOrDefault(t *testing.T) {
	var manager ConfigStringGetterWithDefault
	t.Log("\tWhen using GetOrDefault")
	{
		t.Log("\t\tWhen value not present")
		{
			manager = &Manager{
				make(map[string]string),
			}
			t.Logf("\t\t\tShould return default value")
			{
				cfgValue := manager.GetOrDefault("test.value", "default")

				assert.Equal(t, "default", cfgValue)
			}
		}
		t.Log("\t\tWhen value present")
		{
			testMap := make(map[string]string)
			testMap["test.value"] = "some_value"
			manager = &Manager{
				testMap,
			}
			t.Log("\t\t\tShould return value")
			{
				cfgValue := manager.GetOrDefault("test.value", "default")

				assert.Equal(t, "some_value", cfgValue)
			}
		}
	}
}
