package run

import (
	"reflect"
	"testing"
)

func Test_getInstallTypeByArg(t *testing.T) {
	tests := []struct {
		name string
		arg string
		want installArgType
	}{
		{
			name: "Empty test",
			arg: "",
			want: unknownInstall,
		},
		{
			name: "Random string test",
			arg: "e6fd610e2aa8",
			want: unknownInstall,
		},
		{
			name: "Tag test",
			arg: "localhost:5000/my-app:1.0.7",
			want: tagInstall,
		},
		{
			name: "Tag test without port",
			arg: "localhost/my-app:1.0.7",
			want: tagInstall,
		},
		{
			name: "Tag test with two prefix",
			arg: "domain/localhost/my-app:1.0.7",
			want: tagInstall,
		},
		{
			name: "App test",
			arg: "my-app:1.0.7",
			want: appInstall,
		},
	}

	for _, tt := range tests {
		t.Logf("Test: %s\n", tt.name)
		result := getInstallTypeByArg(tt.arg)
		if !reflect.DeepEqual(tt.want, result) {
			t.Fatalf("Want install type '%v' with result '%v' not equal", tt.want, result)
		}
	}
}
