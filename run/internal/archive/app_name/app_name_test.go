package app_name

import (
	"dam/config"
	"dam/driver/conf/option"
	"dam/driver/logger"
	"fmt"
	"reflect"
	"testing"
)

func TestName(t *testing.T) {
	logger.DebugMode = false
	config.DEF_APP_NAME = "def_app_name"
	config.DEF_APP_VERS = "1.2.3"

	tests := []struct {
		name string
		str  string
		want string
		err  error
	}{
		{
			name: "Empty string",
			str:  "",
			want: "",
			err:  fmt.Errorf("cannot match string by mask '%s'", getRegexpMask()),
		},
		{
			name: "Template",
			str:  "<имя приложения>-<версия приложения>.<8 символов контрольной суммы>.dam",
			want: "<имя приложения>-<версия приложения>.<8 символов контрольной суммы>.dam",
			err:  nil,
		},
		{
			name: "Name with '-' and '.'",
			str:  "def.app-name-1.2.3.0.dam",
			want: "def.app-name-1.2.3.0.dam",
			err:  nil,
		},
		{
			name: "Hash is empty",
			str:  "def.app-name-1.2.3..dam",
			want: "",
			err:  fmt.Errorf("string has two consecutive '%s'", option.Config.Save.GetOptionalSeparator()),
		},
		{
			name: "Bad format",
			str:  "s,.c83muff93md,032",
			want: "",
			err:  fmt.Errorf("cannot match string by mask '%s'", getRegexpMask()),
		},
		{
			name: "Bad format with double '-'",
			str:  "def--app-name-1.2.3.0.dam",
			want: "",
			err:  fmt.Errorf("string has two consecutive '%s'", option.Config.Save.GetFileSeparator()),
		},
		{
			name: "Bug for parsing hash in string",
			str:  "my-app-1.0.0.639533ca.dam",
			want: "my-app-1.0.0.639533ca.dam",
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Logf("Test: %s\n", tt.name)
		inf := NewInfo()
		err := inf.FromString(tt.str)
		if !reflect.DeepEqual(tt.err, err) {
			t.Fatalf("Want error '%v' with result error '%v' not equal", tt.err, err)
		}
		if err != nil {
			continue
		}
		result := inf.FullNameToString()
		if !reflect.DeepEqual(tt.want, result) {
			t.Fatalf("Want string '%v' with result '%v' not equal", tt.want, result)
		}
	}
}

func TestDefaultString(t *testing.T) {
	config.DEF_APP_NAME = "def_app_name"
	config.DEF_APP_VERS = "1.2.3"

	inf := NewInfo()
	want := "def_app_name-1.2.3.0.dam"
	if !reflect.DeepEqual(inf.FullNameToString(), want) {
		t.Fatalf("Result %s not equal string %s", inf.FullNameToString(), want)
	}
}
