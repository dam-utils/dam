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
package docker

import (
	"context"
	"dam/config"
	d_log "dam/decorate/log"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/registry"
	"github.com/docker/docker/client"
)

func SearchAppNames(mask string) *[]string {
	if len(mask) <2 || len(mask)> 100 {
		d_log.Fatal("Search mask for official registry not valid. It must be 2-24 symbols")
	}

	cli, err := client.NewClientWithOpts(client.WithVersion(config.DOCKER_API_VERSION))
	if err != nil {
		d_log.Fatal("Cannot create new docker client")
	}

	searchOpts := types.ImageSearchOptions {}
	searchOpts.Limit = config.OFFICIAL_REPO_SEARCH_APPS_LIMIT
	var results []registry.SearchResult
	results, err = cli.ImageSearch(context.Background(), mask,  searchOpts)
	if err != nil {
		d_log.Fatal("Cannot results of docker search. " + err.Error())
	}

	var appNames []string
	for _, result := range results {
		appNames = append(appNames, result.Name)
	}
	return &appNames
}
