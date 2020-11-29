package meta

import (
	"reflect"
	"testing"
)

func Test_prepareExpString(t *testing.T) {
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
		if !reflect.DeepEqual(tt.wantString, prepareExpString(tt.newString, tt.envs)) {
			t.Fatalf("String '%s' with map '%v' not equal string '%s'", tt.newString, tt.envs, tt.wantString)
		}
	}
}