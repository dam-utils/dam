package env

import (
	"os"
	"testing"
)

func TestGetFileEnv(t *testing.T) {
	const EnvFile = "../../../examples/env/ENVIRONMENT.example"
	var want = map[string]string{
		"TEST":   "test?",
		"TEST2 ": " test2",
		"TEST3":  "test3=test3",
	}

	realEnvTestResult := GetFileEnv(EnvFile)
	equalMaps(t, realEnvTestResult, want)
}

func TestGetDockerFileEnv(t *testing.T) {
	const DockerFile = "../../../examples/env/Dockerfile.example"
	var want = map[string]string{
		"TEST":   "test?",
		"TEST2 ": " test2",
		"TEST3":  "test3=test3",
		"TEST5":  "test5",
		"TEST6":  "test?",
		"TEST7":  "test7 test8",
	}

	realDockerTestResult := GetDockerFileEnv(DockerFile)
	equalMaps(t, realDockerTestResult, want)
}

func TestGetOSEnv(t *testing.T) {
	envPrefix := "TEST_PREFFIX"

	os.Setenv(envPrefix+"TEST", "test")
	os.Setenv(envPrefix+"TEST", "test?")
	os.Setenv(envPrefix+"TEST2 ", " test2")
	os.Setenv(envPrefix+"TEST3", "test3 test4")
	os.Setenv(envPrefix+"TEST5", "test6=test7")

	resultEnv := GetOSEnv(envPrefix)
	if len(resultEnv) != 4 {
		t.Logf("resultEnv: %v", resultEnv)
		t.Fatalf("Not equal size maps in GetOSEnv(): 4 and %v", len(resultEnv))
	}
	checkOSEnv(t, envPrefix+"TEST", "test?")
	checkOSEnv(t, envPrefix+"TEST2 ", " test2")
	checkOSEnv(t, envPrefix+"TEST3", "test3 test4")
	checkOSEnv(t, envPrefix+"TEST5", "test6=test7")
}

func checkOSEnv(t *testing.T, name, value string) {
	if os.Getenv(name) != value {
		t.Fatalf("Not equal os.Getenv(%v) '%v' and value '%v' in OS Env", name, os.Getenv(name), value)
	}
}

func TestMergeEnvs(t *testing.T) {
	tests := []struct{
		name string
		map1 map[string]string
		map2 map[string]string
		want map[string]string
	}{
		{
			name: "Empty test",
			map1: map[string]string{},
			map2: map[string]string{},
			want: map[string]string{},
		},
		{
			name: "Simple test",
			map1: map[string]string{
				"test1": "test1",
				"test2": "test2",
			},
			map2: map[string]string{
				"test1": "test?",
				"test3": "test3",
			},
			want: map[string]string{
				"test1": "test?",
				"test2": "test2",
				"test3": "test3",
			},
		},
		{
			name: "Test without first map",
			map1: map[string]string{},
			map2: map[string]string{
				"test1": "test?",
				"test3": "test3",
			},
			want: map[string]string{
				"test1": "test?",
				"test3": "test3",
			},
		},
		{
			name: "Test without second map",
			map1: map[string]string{
				"test1": "test1",
				"test2": "test2",
			},
			map2: map[string]string{},
			want: map[string]string{
				"test1": "test1",
				"test2": "test2",
			},
		},
	}

	for _, tt := range tests {
		t.Logf("Test name '%s'\n", tt.name)
		equalMaps(t, MergeEnvs(tt.map1, tt.map2), tt.want)
	}
}

func equalMaps(t *testing.T, map1, map2 map[string]string) {
	if len(map1) != len(map2) {
		t.Logf("map 1: %v", map1)
		t.Logf("map 2: %v", map2)
		t.Fatal("Not equal size maps")
	}
	for key1, var1 := range map1 {
		if map2[key1] != var1 {
			t.Logf("map 1: %v", map1)
			t.Logf("map 2: %v", map2)
			t.Fatalf("Not equal map2[key1] '%v' and value of map1 '%v' in reference result", map2[key1], var1)
		}
	}
	for key2, var2 := range map2 {
		if map1[key2] != var2 {
			t.Logf("map 1: %v", map1)
			t.Logf("map 2: %v", map2)
			t.Fatalf("Not equal map1[key2] '%v' and value of map2 '%v' in reference result", map1[key2], key2)
		}
	}
}
