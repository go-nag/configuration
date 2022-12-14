package cfgm

import (
	"fmt"
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
				make(map[string]*configValue),
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
			testMap := make(map[string]*configValue)
			testMap["test.value"] = &configValue{
				value:   "some_value",
				cfgType: str,
			}
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
		t.Log("\tWhen accessing wrong value type")
		{
			testMap := make(map[string]*configValue)
			testMap["test.value"] = &configValue{
				value:   "some_value",
				cfgType: arr,
			}
			manager = &Manager{
				testMap,
			}
			t.Logf("\t\tShould return %s\n", invalidValueType)
			{
				cfgValue, err := manager.Get("test.value")
				assert.Equal(t, "", cfgValue)
				assert.NotNil(t, err)
				assert.Equal(t, invalidValueType, err.Error())
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
				make(map[string]*configValue),
			}
			t.Logf("\t\t\tShould return default value")
			{
				cfgValue := manager.GetOrDefault("test.value", "default")

				assert.Equal(t, "default", cfgValue)
			}
		}
		t.Log("\t\tWhen value present")
		{
			testMap := make(map[string]*configValue)
			testMap["test.value"] = &configValue{
				value:   "some_value",
				cfgType: str,
			}
			manager = &Manager{
				testMap,
			}
			t.Log("\t\t\tShould return value")
			{
				cfgValue := manager.GetOrDefault("test.value", "default")

				assert.Equal(t, "some_value", cfgValue)
			}
		}
		t.Log("\tWhen accessing wrong value type")
		{
			testMap := make(map[string]*configValue)
			testMap["test.value"] = &configValue{
				value:   "some_value",
				cfgType: arr,
			}
			manager = &Manager{
				testMap,
			}
			t.Log("\t\tShould return default value")
			{
				cfgValue := manager.GetOrDefault("test.value", "default")
				assert.NotNil(t, cfgValue)
				assert.Equal(t, "default", cfgValue)
			}
		}
	}
}

func TestManager_GetArr(t *testing.T) {
	var manager ConfigArrayGetter
	t.Log("\tWhen using GetArr")
	{
		t.Log("\t\tWhen value not present")
		{
			manager = &Manager{
				make(map[string]*configValue),
			}
			t.Log("\t\t\tShould return empty array")
			{
				arr := manager.GetArr("array.value")
				assert.Equal(t, 0, len(arr))
			}
		}
		t.Log("\t\tWhen value present")
		{
			testMap := make(map[string]*configValue)
			testMap["array.value"] = &configValue{
				value:   "something0;something1;something2;something3",
				cfgType: arr,
			}
			manager = &Manager{
				testMap,
			}
			t.Log("\t\t\tShould return array value")
			{
				arr := manager.GetArr("array.value")
				for i, value := range arr {
					assert.Equal(t, fmt.Sprintf("something%d", i), value)
				}
			}
		}
		t.Log("\tWhen accessing wrong value type")
		{
			testMap := make(map[string]*configValue)
			testMap["array.value"] = &configValue{
				value:   "something0;something1;something2;something3",
				cfgType: str,
			}
			manager = &Manager{
				testMap,
			}
			t.Log("\t\tShould return empty array")
			{
				arr := manager.GetArr("array.value")

				assert.Equal(t, 0, len(arr))
			}
		}
	}
}
