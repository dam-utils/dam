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