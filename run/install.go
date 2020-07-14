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

	"dam/config"
	"dam/driver/db"
	"dam/driver/docker"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/driver/storage"
	"dam/run/internal"
)

func InstallApp(appCurrentName string) {
	flag.ValidateAppPlusVersion(appCurrentName)

	logger.Success("Start '%s' installing to the system.", appCurrentName)

	tag := dockerPull(appCurrentName)
	tmpMeta := internal.PrepareTmpMetaPath(config.TMP_META_PATH)
	logger.Debug("tag: '%v', tmpMeta: '%v'", tag, tmpMeta)
	containerId := docker.ContainerCreate(tag, "")

	docker.CopyFromContainer(containerId, string(os.PathSeparator)+config.META_DIR_NAME, tmpMeta)
	docker.ContainerRemove(containerId)

	installMeta := filepath.Join(tmpMeta, config.META_DIR_NAME)
	install := getInstall(installMeta)

	fs.RunFile(install)
	fs.Remove(tmpMeta)

	saveInstallAppToDB(tag)
	logger.Success("App '%s' was installed.", appCurrentName)
}

func dockerPull(app string) string {
	defRepo := db.RDriver.GetDefaultRepo()
	if defRepo == nil {
		logger.Fatal("Internal error. Not found default repo")
	}

	var tag string
	if defRepo.Id == storage.OfficialRepo.Id {
		tag = app
	} else {
		tag = defRepo.Server + "/" + app
	}

	if defRepo.Id == storage.OfficialRepo.Id {
		tag = app
	}

	docker.Pull(tag, defRepo)

	return tag
}

func saveInstallAppToDB(tag string) {
	repo := db.RDriver.GetDefaultRepo()
	if repo == nil {
		logger.Fatal("Internal error. Not found default repo")
	}
	_, imageName, imageVersion := internal.SplitTag(tag)

	var app storage.App
	app.RepoID = repo.Id
	app.DockerID = docker.GetImageID(tag)
	app.ImageName = imageName
	app.ImageVersion = imageVersion
	app.Family = docker.GetImageLabel(tag, config.APP_FAMILY_ENV)

	db.ADriver.NewApp(&app)
}

func getInstall(meta string) string {
	inst := filepath.Join(meta, config.INSTALL_FILE_NAME)
	if !fs.IsExistFile(inst) {
		logger.Fatal("Not found '%s' file in meta '%s'", config.INSTALL_FILE_NAME, config.META_DIR_NAME)
	}
	return inst
}


