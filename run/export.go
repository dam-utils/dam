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
	"dam/driver/engine"
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
	logger.Debug("Flags validated with success")

	if !ExportFlags.All {
		exportAppsListToFile(arg)
		logger.Success("Export app list to file '%s'", arg)
	} else {
		tmpDir := arg + "_tmp"
		fs.MkDir(tmpDir)
		defer fs.Remove(tmpDir)

		logger.Debug("Exporting images file to tmp dir ...")
		exportAppsListToFile(arg+string(os.PathSeparator)+config.EXPORT_APPS_FILE_NAME)
		logger.Debug("Exporting docker images to tmp dir ...")
		exportImagesToDir(tmpDir)

		logger.Debug("Creating general apps archive ...")
		fs.Gzip(tmpDir, arg, true)

		logger.Success("Export app list to apps archive '%s'", arg)
	}
}

func exportImagesToDir(tmpDir string) {
	for _, app := range db.ADriver.GetApps() {
		logger.Debug("Export image %s:%s", app.ImageName, app.ImageVersion)
		tmpFilePath := tmpDir+string(os.PathSeparator)+config.SAVE_TMP_FILE_POSTFIX
		tag := app.ImageName+":"+app.ImageVersion

		imageId := engine.VDriver.GetImageID(tag)
		engine.VDriver.SaveImage(imageId, tmpFilePath)

		modifyManifest(tmpFilePath, tag)
		resultPath := tmpDir +
			string(os.PathSeparator) +
			app.ImageName +
			config.SAVE_FILE_SEPARATOR +
			app.ImageVersion +
			config.SAVE_OPTIONAL_SEPARATOR +
			fs.HashFileCRC32(tmpFilePath) +
			config.SAVE_FILE_SEPARATOR +
			fs.FileSize(tmpFilePath) +
			config.SAVE_FILE_POSTFIX
		fs.MoveFile(tmpFilePath, resultPath)
	}
}

func exportAppsListToFile(path string) {
	fs.Touch(path)

	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open apps file '%s' with error: %s", path, err)
	}

	logger.Debug("Getting apps ...")
	apps := db.ADriver.GetApps()
	for _, app := range apps {
		newLine := app.ImageName + config.EXPORT_APP_SEPARATOR + app.ImageVersion + "\n"
		_, err := f.WriteString(newLine)
		if err != nil {
			logger.Fatal("Cannot write to apps file '%s' with error: %s", path, err)
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
}
