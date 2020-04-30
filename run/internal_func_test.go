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
package run_test

import (
	"dam/config"
	"dam/driver/db"
	"log"
	"os"
	"testing"
)

func init(){
	log.SetFlags(0)
	db.Init()
}

func setDefaultConfig() {
	config.DECORATE_MAX_DISPLAY_WIDTH = 100
	config.FILES_DB_REPOS ="Repos"
	config.FILES_DB_APPS ="Apps"
	config.FILES_DB_TMP =".db"
	config.OFFICIAL_REGISTRY_URL="https://registry-1.docker.io/"
}

func dropTestDB(t *testing.T) {
	for _, path := range [...]string{config.FILES_DB_REPOS, config.FILES_DB_REPOS, config.FILES_DB_TMP}{
		_, err := os.Stat(path)
		if err == nil {
			err = os.Remove(path)
			if err != nil{
				t.Fail()
			} else {
				t.Log("File '"+path+"' was removed")
			}
		} else {
			t.Log("Skip removing '"+path+"' file")
		}
	}
}

func printFailInfo(t *testing.T, expResult string, result string) {
	if result != expResult {
		t.Log("Expected result:")
		t.Log("`"+expResult+"`")
		t.Log("Result:")
		t.Log("`"+result+"`")
		t.Fail()
	}
}