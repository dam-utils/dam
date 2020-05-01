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
package apps

import (
	"bufio"
	"dam/config"
	"dam/driver/storage"
	"os"
	"strconv"
	"strings"
)


func GetApps() *[]storage.App {
	var apps []storage.App

	fileHandle, _ := os.Open(config.FILES_DB_APPS)
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		NewLine := fileScanner.Text()
		apps = append(apps, *str2app(NewLine))
	}
	return &apps
}

func str2app(app string) *storage.App {
	// Ex: 1|fd78216a9d61|test_image|2.4.7_18|0||test_image
	App := new(storage.App)

	ParseApp := strings.Split(app, config.FILES_DB_SEPARATOR)
	App.Id, _ = strconv.Atoi(ParseApp[0])
	App.DockerID = ParseApp[1]
	App.ImageName = ParseApp[2]
	App.ImageVersion = ParseApp[3]
	App.RepoID, _ = strconv.Atoi(ParseApp[4])
	if ParseApp[5] == config.FILES_DB_BOOL_FLAG {
		App.MultiVersion = true
	} else {
		App.MultiVersion = false
	}
	App.Family = ParseApp[6]
	return App
}