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
package run

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"dam/driver/db"
	"dam/driver/decorate"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/driver/logger/color"
	"dam/driver/storage"
)

func Import(arg string) {
	flag.ValidateFilePath(arg)

	appAllList := getListFromApps(db.ADriver.GetApps())

	appImportList := appsFromFile(arg)

	appDeleteList, appInstallList, appSkipList := matchLists(appAllList, appImportList)

	decorate.PrintAppList("Skip apps:\n", appSkipList, color.Yellow)
	decorate.PrintAppList("Install apps:\n", appInstallList, color.Green)
	decorate.PrintAppList("Delete apps:\n", appDeleteList, color.Red)

	answer := questionYesNo()
	if answer == false {
		logger.Success("Stop import.")
		os.Exit(0)
	}

	for _, app := range appDeleteList {
		RemoveApp(app.Name)
	}

	for _, app := range appInstallList {
		InstallApp(app.CurrentName())
	}

	logger.Success("Import was successful")
}



func getListFromApps(apps []*storage.App) []*storage.ImportApp {
	result := make([]*storage.ImportApp, 0)

	for _, a := range apps {
		result = append(result, &storage.ImportApp{
			Name:    a.ImageName,
			Version: a.ImageVersion,
		})
	}

	return result
}

func appsFromFile(path string) []*storage.ImportApp {
	result := make([]*storage.ImportApp, 0)

	f, err := os.Open(path)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open file '%s' with error: %s", path, err)
	}

	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		newLine := fileScanner.Text()
		result = append(result, str2importApp(newLine))
	}

	return result
}

func str2importApp(str string) *storage.ImportApp {
	strArray := strings.Split(str, ":")
	if len(strArray) != 2 {
		logger.Fatal("Import file is bad. String '%s' not equal '<app>:<version>'", str)
	}

	return &storage.ImportApp{
		Name:    strArray[0],
		Version: strArray[1],
	}
}

func matchLists(allApps, importApps []*storage.ImportApp) (appDeleteList, appInstallList, appSkipList []*storage.ImportApp) {
	for _, iApp := range importApps {
		var flagExist = false
		for _, aApp := range allApps {
			if iApp.CurrentName() == aApp.CurrentName() {
				appSkipList = append(appSkipList, iApp)
				flagExist = true
			}
		}
		if flagExist == false {
			appInstallList = append(appInstallList, iApp)
		}
	}

	for _, aApp := range allApps {
		var flagExist = false
		for _, iApp := range importApps {
			if aApp.CurrentName() == iApp.CurrentName() {
				flagExist = true
			}
		}
		if flagExist == false {
			appDeleteList = append(appDeleteList, aApp)
		}
	}

	return
}

func questionYesNo() bool {
	fmt.Print("Are you sure you want to apply the changes to the system? [yY/nN]:")

	reader := bufio.NewReader(os.Stdin)
	char, _, err := reader.ReadRune()
	if err != nil {
		logger.Fatal("Cannot read answer rune in Yes/No question.")
	}

	fmt.Println()

	switch char {
	case 'Y','y':
		return true
	case 'N', 'n':
		return false
	default:
		logger.Warn("Unknown symbol '%v'. Stop import.", char)
	}

	return false
}