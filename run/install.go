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
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"dam/config"
	"dam/driver/db"
	"dam/driver/docker"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/storage"
)

type InstallAppSettings struct {

}

var InstallAppFlags = new(InstallAppSettings)

func InstallApp(appCurrentName string) {
	logger.Success("Start '%s' installation.", appCurrentName)

	tag := dockerPull(appCurrentName)
	tmpMeta := prepareTmpMetaPath(config.TMP_META_PATH)
	logger.Debug("tag: '%v', tmpMeta: '%v'", tag, tmpMeta)
	containerId := docker.ContainerCreate(tag, "")

	docker.CopyFromContainer(containerId, string(os.PathSeparator)+config.META_DIR_NAME, tmpMeta)
	docker.ContainerRemove(containerId)

	installMeta := filepath.Join(tmpMeta, config.META_DIR_NAME)
	install := getInstall(installMeta)

	runInstall(install)
	fs.Remove(tmpMeta)

	saveInstallAppToDB(tag)
	logger.Success("App '%s' was installed.", appCurrentName)
}

func dockerPull(app string) string {
	defRepo := db.RDriver.GetDefaultRepo()

	var tag string
	if defRepo.Id == storage.OfficialRepo.Id {
		tag = app
	} else {
		tag = defRepo.Server+"/"+app
	}

	if defRepo.Id == storage.OfficialRepo.Id {
		tag = app
	}

	docker.Pull(tag, defRepo)

	return tag
}

func prepareTmpMetaPath(meta string) string {
	path := fs.GetAbsolutePath(meta)
	fs.Remove(path)
	return path
}

func saveInstallAppToDB(tag string){
	repo := db.RDriver.GetDefaultRepo()
	_, imageName, imageVersion := splitTag(tag)

	var app storage.App
	app.RepoID = repo.Id
	app.DockerID = docker.GetImageID(tag)
	app.ImageName = imageName
	app.ImageVersion = imageVersion
	app.Family = docker.GetImageLabel(tag, config.APP_FAMILY)

	db.ADriver.NewApp(&app)
}

func getInstall(meta string) string {
	inst := filepath.Join(meta, config.INSTALL_FILE_NAME)
	if !fs.IsExistFile(inst) {
		logger.Fatal("Not found '%s' file in meta '%s'", config.INSTALL_FILE_NAME, config.META_DIR_NAME)
	}
	return inst
}

// https://stackoverflow.com/questions/40670228/how-to-run-binary-files-inside-golang-program
func runInstall(installFile string) {
	homeDir := filepath.Dir(installFile)
	err := os.Chdir(homeDir)
	if err != nil {
		logger.Fatal("Cannot change home dir to '%s' with error: %s", homeDir, err.Error())
	}

	c := exec.Command(installFile)
	c.Dir = homeDir
	// set var to get the output
	var outb, errb bytes.Buffer

	// set the output to our variable
	c.Stdout = &outb
	c.Stderr = &errb
	err = c.Run()
	if err != nil {
		logger.Warn(errb.String())
		logger.Fatal("Cannot execute file '%s' with error: %s", installFile, err.Error())
	}
	logger.Info(outb.String())
}

func splitTag(tag string) (string, string, string) {
	n := strings.Split(tag, "/")
	nameWithVersion := n[len(n)-1]
	server := strings.Join(n[:len(n)-1],"/")

	v := strings.Split(nameWithVersion, ":")
	version := v[len(v)-1]
	name := v[0]

	return server, name, version
}