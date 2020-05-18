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
	"net/http"
	"time"

	"dam/config"
	"dam/driver/storage"
)

var SessionURL string

func CheckRepo(repo *storage.Repo, protocol string) error {
	tr := &http.Transport{
		MaxIdleConns:    config.SEARCH_MAX_CONNECTS,
		IdleConnTimeout: time.Duration(config.SEARCH_TIMEOUT_MS) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}
	SessionURL = protocol + "://" + repo.Server + "/"

	req, err := http.NewRequest("GET", SessionURL, nil)
	if err != nil {
		return err
	}
	if repo.Username != "" {
		req.SetBasicAuth(repo.Username, repo.Password)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	return resp.Body.Close()
}
