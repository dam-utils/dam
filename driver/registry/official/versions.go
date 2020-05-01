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
package registry_official

import (
	"dam/config"
	"dam/driver/logger"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func GetAppVersions(app string) *[]string {
	url := config.OFFICIAL_REGISTRY_URL+"/v2/"+app+"/tags/list"

	tr := &http.Transport{
		MaxIdleConns:    config.SEARCH_MAX_CONNECTS,
		IdleConnTimeout: time.Duration(config.SEARCH_TIMEOUT_MS) * time.Millisecond,
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println()
		logger.Fatal(err.Error())
	}
	token := GetBearerToken(app)
	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		log.Println()
		logger.Fatal("Cannot send request to URL: '" + url + "'")
	}
	defer resp.Body.Close()

	type AppVersionsResponse struct {
		Tags []string `json:"tags"`
	}
	var body AppVersionsResponse
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		log.Println()
		logger.Debug(err.Error())
		logger.Fatal("Cannot parse app versions in the body from URL: '" + url + "'")
	}
	vers := body.Tags
	return &vers
}

