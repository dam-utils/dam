package env

import (
	"reflect"
	"testing"
)

func TestPrepareExpString(t *testing.T) {
	test := []struct {
		newString  string
		envs       map[string]string
		wantString string
	}{
		{
			newString:  "",
			envs:       make(map[string]string),
			wantString: "",
		},
		{
			newString:  "${ENV1}",
			envs:       map[string]string{"ENV1": "env1"},
			wantString: "env1",
		},
		{
			newString:  "ENV1",
			envs:       map[string]string{"ENV1": "env1"},
			wantString: "ENV1",
		},
		{
			newString:  "$ENV1",
			envs:       map[string]string{"ENV1": "env1"},
			wantString: "$ENV1",
		},
		{
			newString:  "${ENV1}:${ENV2}",
			envs:       map[string]string{"ENV1": "env1", "ENV2":"env2"},
			wantString: "env1:env2",
		},
	}

	for _, tt := range test {
		if !reflect.DeepEqual(tt.wantString, PrepareExpString(tt.newString, tt.envs)) {
			t.Fatalf("String '%s' with map '%v' not equal string '%s'", tt.newString, tt.envs, tt.wantString)
		}
	}
}

func Test_setDefaultEnv(t *testing.T) {
	//"test change map"
	testMap := map[string]string{
		"test":  "test",
		"test2": "test2",
	}
	resultMap := setDefaultEnv(testMap, "test", "default")
	expectedMap := map[string]string{
		"test":  "test",
		"test2": "test2",
	}
	equalMap(t, resultMap, expectedMap)

	//"test change map with default"
	testMap = map[string]string{
		"test":  "test",
		"test2": "test2",
	}
	resultMap = setDefaultEnv(testMap, "test3", "default")
	expectedMap = map[string]string{
		"test":  "test",
		"test2": "test2",
		"test3": "default",
	}
	equalMap(t, resultMap, expectedMap)

}

func equalMap(t *testing.T, map1, map2 map[string]string) {
	if !reflect.DeepEqual(map1, map2) {
		t.Fatalf("Map1:%v not equal Map2:%v", map1, map2)
	}
}
