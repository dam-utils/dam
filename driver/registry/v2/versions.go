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
package registry_v2

import (
	"encoding/json"
	"net/http"
	"time"

	"dam/config"
	"dam/driver/db/storage"
	"dam/driver/logger"
)

func GetAppVersions(repo *storage.Repo, appName string) *[]string {
	tr := &http.Transport{
		MaxIdleConns:    config.SEARCH_MAX_CONNECTS,
		IdleConnTimeout: time.Duration(config.SEARCH_TIMEOUT_MS) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}

	url := SessionURL + "v2/" + appName + "/tags/list"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil || req == nil {
		logger.Fatal("Cannot create new request for get URL '%s' with error: %s", url, err)
	}
	if repo.Username != "" {
		req.SetBasicAuth(repo.Username, repo.Password)
	}

	resp, err := client.Do(req)
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot get response from URL '%s' with error: %s", url, err)
	}

	type AppVersionsResponse struct {
		Tags []string `json:"tags"`
	}
	var body AppVersionsResponse

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		logger.Fatal("Cannot parse app versions in the body from URL '%s' with error: %s", url, err)
	}
	vers := body.Tags
	return &vers
}