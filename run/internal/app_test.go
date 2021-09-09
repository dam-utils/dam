package internal

import (
	"reflect"
	"testing"
)

func TestSplitTag(t *testing.T) {
	tests := []struct {
		name string
		arg string
		wantRepo string
		wantName string
		wantVersion string
	}{
		{
			name: "Empty test",
			arg: "",
			wantRepo: "",
			wantName: "",
			wantVersion: "",
		},
		{
			name: "Random string test",
			arg: "dk49dk58cdd5ske32",
			wantRepo: "",
			wantName: "",
			wantVersion: "",
		},
		{
			name: "Simple test",
			arg: "server/name:version",
			wantRepo: "server",
			wantName: "name",
			wantVersion: "version",
		},
		{
			name: "With two slash test",
			arg: "docker.io/library/ubuntu:14.04",
			wantRepo: "docker.io/library",
			wantName: "ubuntu",
			wantVersion: "14.04",
		},
		{
			name: "Without repo test",
			arg: "ubuntu:16.04",
			wantRepo: "",
			wantName: "ubuntu",
			wantVersion: "16.04",
		},
		{
			name: "With port in repo test",
			arg: "localhost:5000/ubuntu:18.04",
			wantRepo: "localhost:5000",
			wantName: "ubuntu",
			wantVersion: "18.04",
		},
	}

	for _, tt := range tests {
		t.Logf("Test: %s\n", tt.name)
		resultRepo, resultName, resultVersion := SplitTag(tt.arg)
		if !reflect.DeepEqual(tt.wantRepo, resultRepo) {
			t.Fatalf("Want repo '%v' with result '%v' not equal", tt.wantRepo, resultRepo)
		}
		if !reflect.DeepEqual(tt.wantName, resultName) {
			t.Fatalf("Want name '%v' with result '%v' not equal", tt.wantName, resultName)
		}
		if !reflect.DeepEqual(tt.wantVersion, resultVersion) {
			t.Fatalf("Want version '%v' with result '%v' not equal", tt.wantVersion, resultVersion)
		}
	}
}
