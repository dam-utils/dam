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
	"os"
	"strconv"
	"strings"

	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/storage"
)


func GetApps() []*storage.App {
	var apps []*storage.App

	fileHandle, _ := os.Open(config.FILES_DB_APPS)
	defer fileHandle.Close()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		NewLine := fileScanner.Text()
		apps = append(apps, str2app(NewLine))
	}
	return apps
}

func GetAppById(id int) *storage.App {
	apps := GetApps()
	for _, app := range apps {
		if app.Id == id {
			return app
		}
	}
	return nil
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

func NewApp(app *storage.App) {
	apps := GetApps()
	app.Id = getNewAppID(apps)

	newApps := append(apps, app)
	saveApps(newApps)
}

func getNewAppID(apps []*storage.App) int {
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

func saveApps(apps []*storage.App) {
	f, err := os.OpenFile(config.FILES_DB_TMP, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer f.Close()

	for _, app := range apps {
		newLine := app2str(app)
		_, err := f.WriteString(*newLine)
		if err != nil {
			logger.Fatal(err.Error())
		}
	}
	err = f.Sync()
	if err != nil {
		logger.Fatal(err.Error())
	}
	err = f.Close()
	if err != nil {
		logger.Fatal(err.Error())
	}

	logger.Debug("Move '%s' to '%s'", config.FILES_DB_TMP, config.FILES_DB_APPS)
	fs.MoveFile(config.FILES_DB_TMP, config.FILES_DB_APPS)
}

func app2str(app *storage.App) *string {
	var appStr string
	sep := config.FILES_DB_SEPARATOR

	multiVers := ""
	if app.MultiVersion {
		multiVers = config.DECORATE_BOOL_FLAG
	}

	fields := []string{
		strconv.Itoa(app.Id),
		app.DockerID, app.ImageName,
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


