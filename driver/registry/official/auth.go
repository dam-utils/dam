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
	d_log "dam/decorate/log"
	"encoding/json"
	"log"
	"net/http"
)

func GetBearerToken(app string) string {
	url := config.OFFICIAL_REGISTRY_AUTH_URL+ "&scope=repository:"+app+":pull"
	resp, err := http.Get(url)
	if err != nil {
		log.Println()
		d_log.Debug(err.Error())
		d_log.Fatal("Cannot get token from URL: '" + url + "'")
	}
	defer resp.Body.Close()

	type TokenResponse struct {
		AccessToken string `json:"access_token"`
	}
	var body TokenResponse

	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		log.Println()
		d_log.Debug(err.Error())
		d_log.Fatal("Cannot parse token in the body from URL: '" + url + "'. Err: "+err.Error())
	}
	return body.AccessToken
}

