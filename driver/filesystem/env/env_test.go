// Copyright 2020 The Docker Applications Manager Authors
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
package env

import (
	"os"
	"testing"
)

func TestGetFileEnv(t *testing.T) {
	const EnvFile = "../../../examples/env/ENVIRONMENT.example"
	var want = map[string]string {
		"TEST" : "test?",
		"TEST2 " : " test2",
		"TEST3" : "test3=test3",
	}

	realEnvTestResult := GetFileEnv(EnvFile)
	equalMaps(t, realEnvTestResult, want)
}

func TestGetDockerFileEnv(t *testing.T) {
	const DockerFile = "../../../examples/env/Dockerfile.example"
	var want = map[string]string {
		"TEST" : "test?",
		"TEST2 " : " test2",
		"TEST3" : "test3=test3",
		"TEST5" : "test5",
		"TEST6" : "test?",
		"TEST7" : "test7 test8",
	}

	realDockerTestResult := GetDockerFileEnv(DockerFile)
	equalMaps(t, realDockerTestResult, want)
}

func TestGetOSEnv(t *testing.T){
	envPrefix := "TEST_PREFFIX"

	os.Setenv(envPrefix+"TEST", "test")
	os.Setenv(envPrefix+"TEST", "test?")
	os.Setenv(envPrefix+"TEST2 ", " test2")
	os.Setenv(envPrefix+"TEST3", "test3 test4")
	os.Setenv(envPrefix+"TEST5", "test6=test7")

	resultEnv := GetOSEnv(envPrefix)
	if len(resultEnv) != 4 {
		t.Logf("resultEnv: %v",resultEnv)
		t.Fatalf("Not equal size maps in GetOSEnv(): 4 and %v", len(resultEnv))
	}
	checkOSEnv(t, envPrefix+"TEST", "test?")
	checkOSEnv(t, envPrefix+"TEST2 ", " test2")
	checkOSEnv(t, envPrefix+"TEST3", "test3 test4")
	checkOSEnv(t, envPrefix+"TEST5", "test6=test7")
}

func checkOSEnv(t *testing.T, name, value string) {
	if os.Getenv(name) != value {
		t.Fatalf("Not equal os.Getenv(%v) '%v' and value '%v' in OS Env",name, os.Getenv(name), value)
	}
}

func TestMergeEnvs(t *testing.T) {
var	map1 = map[string]string {
	"test1": "test1",
	"test2": "test2",
}
	var	map2 = map[string]string {
		"test1": "test?",
		"test3": "test3",
	}

	var resultMap = map[string]string {
		"test1": "test?",
		"test2": "test2",
		"test3": "test3",
	}

	equalMaps(t, MergeEnvs(map1, map2), resultMap)
}

func equalMaps(t *testing.T, map1, map2 map[string]string) {
	if len(map1) != len(map2) {
		t.Logf("map 1: %v",map1)
		t.Logf("map 2: %v",map2)
		t.Fatal("Not equal size maps")
	}
	for key1, var1 :=  range map1 {
		if map2[key1] != var1 {
			t.Logf("map 1: %v",map1)
			t.Logf("map 2: %v",map2)
			t.Fatalf("Not equal map2[key1] '%v' and value of map1 '%v' in reference result", map2[key1], var1)
		}
	}
	for key2, var2 :=  range map2 {
		if map1[key2] != var2 {
			t.Logf("map 1: %v",map1)
			t.Logf("map 2: %v",map2)
			t.Fatalf("Not equal map1[key2] '%v' and value of map2 '%v' in reference result", map1[key2], key2)
		}
	}
}