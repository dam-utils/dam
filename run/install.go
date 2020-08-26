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
	"dam/driver/engine"
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"dam/config"
	"dam/driver/db"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/run/internal"
)

func InstallApp(appCurrentName string) {
	var isFileInstalling bool

	if fs.IsExistFile(appCurrentName) {
		isFileInstalling = true
		flag.ValidateFilePath(appCurrentName)
	} else {
		isFileInstalling = false
		flag.ValidateAppPlusVersion(appCurrentName)
	}
	logger.Debug("Flags validated with success")
	logger.Success("Start '%s' installing to the system.", appCurrentName)

	logger.Debug("Preparing docker image ...")
	var tag string
	if isFileInstalling {
		tag = getTagFromArchiveManifest(appCurrentName)
		engine.VDriver.LoadImage(appCurrentName)
	} else {
		tag = dockerPull(appCurrentName)
	}

	logger.Debug("Getting meta ...")
	tmpMeta := internal.PrepareTmpMetaPath(config.TMP_META_PATH)
	containerId := engine.VDriver.ContainerCreate(tag, "")
	engine.VDriver.CopyFromContainer(containerId, string(os.PathSeparator)+config.META_DIR_NAME, tmpMeta)
	engine.VDriver.ContainerRemove(containerId)

	logger.Debug("Installing meta ...")
	installMeta := filepath.Join(tmpMeta, config.META_DIR_NAME)
	install := getInstall(installMeta)

	logger.Debug("Removing tmp files ...")
	fs.RunFile(install)
	fs.Remove(tmpMeta)

	logger.Debug("Saving to DB ...")
	saveInstallAppToDB(tag)
	logger.Success("App '%s' was installed.", appCurrentName)
}

func dockerPull(app string) string {
	defRepo := db.RDriver.GetDefaultRepo()
	if defRepo == nil {
		logger.Fatal("Internal error. Not found default repo")
	}

	var tag string
	if defRepo.Id == db.OfficialRepo.Id {
		tag = app
	} else {
		tag = defRepo.Server + "/" + app
	}

	if defRepo.Id == db.OfficialRepo.Id {
		tag = app
	}

	engine.VDriver.Pull(tag, defRepo)

	return tag
}

func saveInstallAppToDB(tag string) {
	repo := db.RDriver.GetDefaultRepo()
	if repo == nil {
		logger.Fatal("Internal error. Not found default repo")
	}
	_, imageName, imageVersion := internal.SplitTag(tag)

	var app db.App
	app.RepoID = repo.Id
	app.DockerID = engine.VDriver.GetImageID(tag)
	app.ImageName = imageName
	app.ImageVersion = imageVersion
	app.Family = engine.VDriver.GetImageLabel(tag, config.APP_FAMILY_ENV)

	db.ADriver.NewApp(&app)
}

func getInstall(meta string) string {
	inst := filepath.Join(meta, config.INSTALL_FILE_NAME)
	if !fs.IsExistFile(inst) {
		logger.Fatal("Not found '%s' file in meta '%s'", config.INSTALL_FILE_NAME, config.META_DIR_NAME)
	}
	return inst
}

func getTagFromArchiveManifest(appCurrentName string) string {
	//TODO read manifest without archive uncompressing
	gzipFile := fs.Gunzip(appCurrentName)
	//defer fs.Remove(gzipFile)
	tarGzipDir := fs.Untar(gzipFile)
	//defer fs.Remove(tarGzipDir)

	manifestFile := tarGzipDir + string(filepath.Separator) + config.SAVE_MANIFEST_FILE

	content, err := os.Open(manifestFile)
	defer func() {
		if content != nil {
			content.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open the manifest file '%s' with error: %s", manifestFile, err)
	}

	type manifest struct {
		RepoTags []string `json:"RepoTags"`
	}

	result := make([]manifest, 0)
	byteValue, err := ioutil.ReadAll(content)
	if err != nil {
		logger.Fatal("Cannot read content in manifest file with error: %s", err)
	}

	err = json.Unmarshal(byteValue, &result)
	if err != nil {
		logger.Fatal("Cannot unmarshal manifest file with error: %s", err)
	}

	if len(result) > 0 {
		if len(result[0].RepoTags) > 0 {
			flag.ValidateAppPlusVersion(result[0].RepoTags[0])
			return result[0].RepoTags[0]
		}
	}

	logger.Fatal("Cannot get manifest tag from archive '%s'", appCurrentName)
	return ""
}
