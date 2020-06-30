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
	"os"

	"dam/config"
	"dam/driver/db"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
)

type ExportSettings struct {
	All bool
}

var ExportFlags = new(ExportSettings)

func Export(arg string) {
	flag.ValidateFilePath(arg)

	fs.Touch(arg)

	f, err := os.OpenFile(arg, os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open apps file '%s' with error: %s", arg, err.Error())
	}

	apps := db.ADriver.GetApps()
	for _, app := range apps {
		newLine := app.ImageName+config.EXPORT_APP_SEPARATOR+app.ImageVersion+"\n"
		_, err := f.WriteString(newLine)
		if err != nil {
			logger.Fatal("Cannot write to apps file '%s' with error: %s", arg, err.Error())
		}
	}

	err = f.Sync()
	if err != nil {
		logger.Fatal("Cannot sync apps file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
	}
	err = f.Close()
	if err != nil {
		logger.Fatal("Cannot close from apps file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
	}

	logger.Info("Export file save to '%s'", arg)
}
