package run

import (
	"reflect"
	"testing"
)

func Test_getSaveTypeByArg(t *testing.T) {
	tests := []struct {
		name string
		arg string
		want saveArgType
	}{
		{
			name: "Empty test",
			arg: "",
			want: unknownSave,
		},
		{
			name: "Random string test",
			arg: "e6fd610e2aa8",
			want: unknownSave,
		},
		{
			name: "Tag test",
			arg: "localhost:5000/my-app:1.0.7",
			want: tagSave,
		},
		{
			name: "Tag test without port",
			arg: "localhost/my-app:1.0.7",
			want: tagSave,
		},
		{
			name: "Tag test with two prefix",
			arg: "domain/localhost/my-app:1.0.7",
			want: tagSave,
		},
		{
			name: "App test",
			arg: "my-app:1.0.7",
			want: appSave,
		},
	}

	for _, tt := range tests {
		t.Logf("Test: %s\n", tt.name)
		result := getSaveTypeByArg(tt.arg)
		if !reflect.DeepEqual(tt.want, result) {
			t.Fatalf("Want save type '%v' with result '%v' not equal", tt.want, result)
		}
	}
}
