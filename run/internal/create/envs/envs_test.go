package envs

import (
	"reflect"
	"testing"

	"dam/config"
)

func TestEnv_Envs(t *testing.T) {
	test := []struct {
		name       string
		init       map[string]string
		flagName   string
		flagVers   string
		flagFamily string
		flagMulti  string
		repo       string
		want       map[string]string
	}{
		{
			name:       "Empty env map",
			init:       map[string]string{},
			flagName:   "",
			flagVers:   "",
			flagFamily: "",
			flagMulti:  "",
			repo:       "",
			want: map[string]string{
				config.APP_NAME_ENV:         config.DEF_APP_NAME,
				config.APP_VERS_ENV:         config.DEF_APP_VERS,
				config.APP_FAMILY_ENV:       config.DEF_APP_NAME,
				config.APP_MULTIVERSION_ENV: "false",
				config.APP_TAG_ENV:          config.DEF_APP_NAME + ":" + config.DEF_APP_VERS,
			},
		},
		{
			name:       "Test app name and version",
			init:       map[string]string{
				config.APP_NAME_ENV:         "app-name",
				config.APP_VERS_ENV:         "1.0.0",
			},
			flagName:   "",
			flagVers:   "",
			flagFamily: "",
			flagMulti:  "",
			repo:       "",
			want: map[string]string{
				config.APP_NAME_ENV:         "app-name",
				config.APP_VERS_ENV:         "1.0.0",
				config.APP_FAMILY_ENV:       "app-name",
				config.APP_MULTIVERSION_ENV: "false",
				config.APP_TAG_ENV:          "app-name:1.0.0",
			},
		},
		{
			name:       "Test app name and version with flag",
			init:       map[string]string{
				config.APP_NAME_ENV:         "app-name",
				config.APP_VERS_ENV:         "1.0.0",
			},
			flagName:   "flag-name",
			flagVers:   "2.0.0",
			flagFamily: "",
			flagMulti:  "",
			repo:       "",
			want: map[string]string{
				config.APP_NAME_ENV:         "flag-name",
				config.APP_VERS_ENV:         "2.0.0",
				config.APP_FAMILY_ENV:       "flag-name",
				config.APP_MULTIVERSION_ENV: "false",
				config.APP_TAG_ENV:          "flag-name:2.0.0",
			},
		},
		{
			name:       "Test app name and version with flag without envs",
			init:       map[string]string{},
			flagName:   "flag-name",
			flagVers:   "2.0.0",
			flagFamily: "",
			flagMulti:  "",
			repo:       "",
			want: map[string]string{
				config.APP_NAME_ENV:         "flag-name",
				config.APP_VERS_ENV:         "2.0.0",
				config.APP_FAMILY_ENV:       "flag-name",
				config.APP_MULTIVERSION_ENV: "false",
				config.APP_TAG_ENV:          "flag-name:2.0.0",
			},
		},
		{
			name:       "Test family",
			init:       map[string]string{
				config.APP_FAMILY_ENV: "test-family",
			},
			flagName:   "flag-name",
			flagVers:   "2.0.0",
			flagFamily: "",
			flagMulti:  "",
			repo:       "",
			want: map[string]string{
				config.APP_NAME_ENV:         "flag-name",
				config.APP_VERS_ENV:         "2.0.0",
				config.APP_FAMILY_ENV:       "test-family",
				config.APP_MULTIVERSION_ENV: "false",
				config.APP_TAG_ENV:          "flag-name:2.0.0",
			},
		},
		{
			name:       "Test flag family",
			init:       map[string]string{
				config.APP_FAMILY_ENV: "test-family",
			},
			flagName:   "flag-name",
			flagVers:   "2.0.0",
			flagFamily: "flag-family",
			flagMulti:  "",
			repo:       "",
			want: map[string]string{
				config.APP_NAME_ENV:         "flag-name",
				config.APP_VERS_ENV:         "2.0.0",
				config.APP_FAMILY_ENV:       "flag-family",
				config.APP_MULTIVERSION_ENV: "false",
				config.APP_TAG_ENV:          "flag-name:2.0.0",
			},
		},
		{
			name:       "Test multiversion",
			init:       map[string]string{},
			flagName:   "flag-name",
			flagVers:   "2.0.0",
			flagFamily: "flag-family",
			flagMulti:  "true",
			repo:       "",
			want: map[string]string{
				config.APP_NAME_ENV:         "flag-name",
				config.APP_VERS_ENV:         "2.0.0",
				config.APP_FAMILY_ENV:       "flag-family",
				config.APP_MULTIVERSION_ENV: "true",
				config.APP_TAG_ENV:          "flag-name:2.0.0",
			},
		},
		{
			name:       "Test tag",
			init:       map[string]string{},
			flagName:   "flag-name",
			flagVers:   "2.0.0",
			flagFamily: "flag-family",
			flagMulti:  "true",
			repo:       "new-repo",
			want: map[string]string{
				config.APP_NAME_ENV:         "flag-name",
				config.APP_VERS_ENV:         "2.0.0",
				config.APP_FAMILY_ENV:       "flag-family",
				config.APP_MULTIVERSION_ENV: "true",
				config.APP_TAG_ENV:          "new-repo/flag-name:2.0.0",
			},
		},
	}

	for _, tt := range test {
		t.Log("Starting test: " + tt.name)
		s := NewStorage(tt.init)
		s.InitAppName(config.DEF_APP_NAME, tt.flagName)
		s.InitAppVersion(config.DEF_APP_VERS, tt.flagVers)
		s.InitAppFamily(tt.flagFamily)
		s.InitAppMultiversion(tt.flagMulti)
		s.InitAppTag(tt.repo)
		if !reflect.DeepEqual(tt.want, s.Envs()) {
			t.Fatalf("Map '%v' not equal '%v'", tt.init, tt.want)
		}
	}
}
