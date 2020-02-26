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
	"dam/config"
	d_log "dam/decorate/log"
	"dam/driver/storage"
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

type ResponseGetAppNames struct {
	Repositories []string `json:"repositories"`
}

func GetAppNames(repo *storage.Repo) *[]string {
	tr := &http.Transport{
		MaxIdleConns:    config.SEARCH_MAX_CONNECTS,
		IdleConnTimeout: time.Duration(config.SEARCH_TIMEOUT_MS) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}

	url := SessionURL + "v2/_catalog?n="+strconv.Itoa(config.INTERNAL_REPO_SEARCH_APPS_LIMIT)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		d_log.Fatal("Cannot create new request for get URL: '"+url+"'")
	}
	if repo.Username != "" {
		req.SetBasicAuth(repo.Username, repo.Password)
	}

	resp, err := client.Do(req)
	if err != nil {
		d_log.Fatal("Cannot get response from URL: '"+url+"'")
	}
	defer resp.Body.Close()

	var body ResponseGetAppNames
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		d_log.Fatal("Cannot parse response from URL: '"+url+"'")
	}
	return &body.Repositories
}



