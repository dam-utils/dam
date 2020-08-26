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
	"dam/driver/structures"

	"os"
	"strconv"
	"strings"

	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/validate"
)


func GetApps() []*structures.App {
	var apps []*structures.App

	f, err := os.Open(config.FILES_DB_APPS)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open file '%s' with error: %s", config.FILES_DB_APPS, err)
	}

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		newLine := fileScanner.Text()
		apps = append(apps, str2app(newLine))
	}
	return apps
}

func GetAppById(id int) *structures.App {
	apps := GetApps()
	for _, app := range apps {
		if app.Id == id {
			return app
		}
	}
	return nil
}

func str2app(str string) *structures.App {
	app := new(structures.App)
	strArray := strings.Split(str, config.FILES_DB_SEPARATOR)

	if validate.CheckAppID(strArray[0]) != nil {
		logger.Fatal("Internal error. Cannot parse the app ID in line '%s'", str)
	}
	if validate.CheckDockerID(strArray[1]) != nil {
		logger.Fatal("Internal error. Cannot parse the docker ID in line '%s'", str)
	}
	if validate.CheckAppName(strArray[2]) != nil {
		logger.Fatal("Internal error. Cannot parse the app name in line '%s'", str)
	}
	if validate.CheckVersion(strArray[3]) != nil {
		logger.Fatal("Internal error. Cannot parse the app version in line '%s'", str)
	}
	if validate.CheckRepoID(strArray[4]) != nil {
		logger.Fatal("Internal error. Cannot parse the repo id in line '%s'", str)
	}
	if validate.CheckBool(strArray[5]) != nil {
		logger.Fatal("Internal error. Cannot parse the multiversion flag in line '%s'", str)
	}
	if validate.CheckLabel(strArray[6]) != nil {
		logger.Fatal("Internal error. Cannot parse the family flag in line '%s'", str)
	}

	app.Id, _ = strconv.Atoi(strArray[0])
	app.DockerID = strArray[1]
	app.ImageName = strArray[2]
	app.ImageVersion = strArray[3]
	app.RepoID, _ = strconv.Atoi(strArray[4])
	if strArray[5] == config.FILES_DB_BOOL_FLAG {
		app.MultiVersion = true
	}
	app.Family = strArray[6]
	return app
}

func NewApp(app *structures.App) {
	apps := GetApps()
	app.Id = getNewAppID(apps)

	newApps := append(apps, app)
	saveApps(newApps)
}

func getNewAppID(apps []*structures.App) int {
	res := 0

	if len(apps) == 0 {
		return 0
	}
	for _, app := range apps {
		if app.Id >= res {
			res = app.Id
		}
	}
	return res + 1
}

func saveApps(apps []*structures.App) {
	f, err := os.OpenFile(config.FILES_DB_TMP, os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open apps file '%s' with error: %s", config.FILES_DB_TMP, err)
	}

	for _, app := range apps {
		newLine := app2str(app)
		_, err := f.WriteString(*newLine)
		if err != nil {
			logger.Fatal("Cannot write to apps file '%s' with error: %s", config.FILES_DB_TMP, err)
		}
	}
	err = f.Sync()
	if err != nil {
		logger.Fatal("Cannot sync apps file '%s' with error: %s", config.FILES_DB_TMP, err)
	}
	err = f.Close()
	if err != nil {
		logger.Fatal("Cannot close from apps file '%s' with error: %s", config.FILES_DB_TMP, err)
	}

	logger.Debug("Move '%s' to '%s'", config.FILES_DB_TMP, config.FILES_DB_APPS)
	fs.MoveFile(config.FILES_DB_TMP, config.FILES_DB_APPS)
}

func app2str(app *structures.App) *string {
	var appStr string
	sep := config.FILES_DB_SEPARATOR

	multiVers := ""
	if app.MultiVersion {
		multiVers = config.DECORATE_BOOL_FLAG
	}

	fields := []string{
		strconv.Itoa(app.Id),
		app.DockerID,
		app.ImageName,
		app.ImageVersion,
		strconv.Itoa(app.RepoID),
		multiVers,
		app.Family,
	}
	lenF := len(fields)
	for i, field := range fields {
		if i == lenF - 1 {
			appStr = appStr + field + "\n"
		} else {
			appStr = appStr + field + sep
		}
	}
	return &appStr
}

func ExistFamily(family string) bool {
	apps := GetApps()
	for _, a := range apps {
		if a.Family == family {
			return true
		}
	}
	return false
}

func RemoveAppById(id int) {
	newApps := make([]*structures.App, 0)

	apps := GetApps()
	for _, a := range apps {
		if a.Id != id {
			newApps = append(newApps, a)
		}
	}
	if len(newApps) < len(apps) {
		saveApps(newApps)
	} else {
		logger.Fatal("Not found Id of the App in DB")
	}
}
