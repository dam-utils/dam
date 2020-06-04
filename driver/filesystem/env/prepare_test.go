package env

import (
	"reflect"
	"testing"
)

func Test_setDefaultEnv(t *testing.T){

	//"test change map"
	testMap := map[string]string {
		"test" : "test",
		"test2" : "test2",
	}
	resultMap := setDefaultEnv(testMap, "test", "default")
	expectedMap := map[string]string {
		"test" : "test",
		"test2" : "test2",
	}
	equalMap(t, resultMap, expectedMap)

	//"test change map with default"
	testMap = map[string]string {
		"test" : "test",
		"test2" : "test2",
	}
	resultMap = setDefaultEnv(testMap, "test3", "default")
	expectedMap = map[string]string {
		"test" : "test",
		"test2" : "test2",
		"test3" : "default",
	}
	equalMap(t, resultMap, expectedMap)


}

func equalMap(t *testing.T, map1, map2 map[string]string) {
	if !reflect.DeepEqual(map1, map2) {
		t.Fatalf("Map1:%v not equal Map2:%v", map1, map2)
	}
}
