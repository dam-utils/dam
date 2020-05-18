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
	// app name
	val, ok := envs[config.APP_NAME_ENV]
	if !ok || val == "" {
		val = config.DEF_APP_NAME
	}
	envs[config.APP_NAME_ENV]=val

	//app vers
	val, ok = envs[config.APP_VERS_ENV]
	if !ok || val == "" {
		val = config.DEF_APP_VERS
	}
	envs[config.APP_VERS_ENV]=val

	//app family
	val, ok = envs[config.APP_FAMILY]
	if !ok || val == "" {
		val = envs[config.APP_NAME_ENV]
	}
	envs[config.APP_FAMILY]=val
	return envs
}
