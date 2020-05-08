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
package registry

import (
	"strings"

	"dam/config"
	"dam/driver/docker"
	"dam/driver/logger"
	registry_official "dam/driver/registry/official"
	registry_v2 "dam/driver/registry/v2"
	"dam/driver/storage"
)

func CheckRepository(repo *storage.Repo) {
	if repo.Id ==1 {
		return
	}

	for _, protocol := range config.SEARCH_PROTOCOL_STRATEGY {
		err := registry_v2.CheckRepo(repo, protocol)
		if err != nil {
			// TODO create debug message
			//logger.Println("WARN: Cannot connect to default registry '" + repo.Name + "' for '" + protocol + "' protocol")
		} else {
			return
		}
	}
	logger.Fatal("Cannot connect to default registry '" + repo.Name + "'")
}

func GetAppNamesByMask(repo *storage.Repo, mask string) *[]string {
	if repo.Id == 1 {
		return docker.SearchAppNames(mask)
	}
	names := registry_v2.GetAppNames(repo)
	return filterAppNamesByMask(names, mask)
}

func filterAppNamesByMask(names *[]string, mask string) *[]string {
	var res []string
	for _, name := range *names {
		if strings.Contains(name, mask) {
			res = append(res, name)
		}
	}
	return &res
}

func GetVersions(repo *storage.Repo, appName string) *[]string {
	var versions *[]string
	if repo.Id == 1 {
		versions = registry_official.GetAppVersions(appName)
	} else {
		versions = registry_v2.GetAppVersions(repo, appName)
	}
	return versions
}