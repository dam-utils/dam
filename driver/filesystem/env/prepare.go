package env

import (
	"strings"

	"dam/config"
)

func PrepareExpString(s string, envs map[string]string) string {
	var newString string
	for envKey, envVal := range envs {
		newString = strings.ReplaceAll(s, "${"+envKey+"}", envVal)
	}
	return newString
}

func PrepareProjectEnvs(envs map[string]string) map[string]string {
	envs = setDefaultEnv(envs, config.APP_NAME_ENV, config.DEF_APP_NAME)
	envs = setDefaultEnv(envs, config.APP_VERS_ENV, config.DEF_APP_VERS)
	envs = setDefaultEnv(envs, config.APP_FAMILY_ENV, config.APP_NAME_ENV)
	return envs
}

func setDefaultEnv(envs map[string]string, env, defEnv string) map[string]string {
	val, ok := envs[env]
	if !ok || val == "" {
		val = defEnv
	}
	envs[env]=val

	return envs
}