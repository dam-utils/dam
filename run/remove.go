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
	"path/filepath"
	"strings"

	"dam/config"
	"dam/driver/db"
	"dam/driver/docker"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/storage"
	"dam/driver/validate"
	"dam/run/internal"
)

type RemoveAppSettings struct {
	Force   bool
}

var RemoveAppFlags = new(RemoveAppSettings)

func RemoveApp(name string) {
	validate.AppName(name)

	app := getAppIdByName(name)
	tag := getTagFormApp(app)

	logger.Success("Start app '%s:%s' removing from the system.",app.ImageName, app.ImageVersion)

	tmpMeta := internal.PrepareTmpMetaPath(config.TMP_META_PATH)
	logger.Debug("tmpMeta: '%v'", tmpMeta)
	containerId := docker.ContainerCreate(tag, "")

	docker.CopyFromContainer(containerId, string(os.PathSeparator)+config.META_DIR_NAME, tmpMeta)
	docker.ContainerRemove(containerId)

	uninstallMeta := filepath.Join(tmpMeta, config.META_DIR_NAME)
	uninstall := getUninstall(uninstallMeta)

	fs.RunFile(uninstall)
	fs.Remove(tmpMeta)

	logger.Success("App '%s:%s' was removed.",app.ImageName, app.ImageVersion)
}

func getAppIdByName(name string) *storage.App {
	var id int

	apps := db.ADriver.GetApps()
	for _, app := range apps {
		if app.ImageName == name {
			id = app.Id
		}
	}

	if id == 0 {
		logger.Fatal("Not found app with name '%s' in DB", name)
	}

	app := db.ADriver.GetAppById(id)
	logger.Debug("Remove app '%s'", app)
	if app == nil {
		logger.Fatal("Not found app with name '%s' in DB", name)
	}

	return app
}

func getTagFormApp(app *storage.App) string {
	var tag strings.Builder

	if app.RepoID != storage.OfficialRepo.Id {
		repo := db.RDriver.GetRepoById(app.RepoID)
		tag.WriteString(repo.Server)
		tag.WriteString("/")
	}
	tag.WriteString(app.ImageName)
	tag.WriteString(":")
	tag.WriteString(app.ImageVersion)

	return tag.String()
}

func getUninstall(meta string) string {
	uninst := filepath.Join(meta, config.UNINSTALL_FILE_NAME)
	if !fs.IsExistFile(uninst) {
		logger.Fatal("Not found '%s' file in meta '%s'", config.UNINSTALL_FILE_NAME, config.META_DIR_NAME)
	}
	return uninst
}