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
	"dam/driver/decorate"
	"dam/driver/engine"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/run/internal"
)

func InfoApp(tag string) {
	flag.ValidateAppPlusVersion(tag)
	logger.Debug("Flags validated with success")

	logger.Debug("Getting meta ...")
	tmpDir := internal.PrepareTmpMetaPath(config.TMP_META_PATH)
	defer fs.Remove(tmpDir)

	containerId := engine.VDriver.ContainerCreate(tag, "")
	engine.VDriver.CopyFromContainer(containerId, string(os.PathSeparator)+config.META_DIR_NAME, tmpDir)
	engine.VDriver.ContainerRemove(containerId)

	logger.Debug("Printing description ...")
	decorate.Println()
	decorate.PrintDescription(filepath.Join(tmpDir, config.META_DIR_NAME, config.DESCRIPTION_FILE_NAME))
	decorate.Println()

	logger.Debug("Printing family label ...")
	family := internal.GetFamily(tag)
	decorate.PrintLabel(family)

	logger.Debug("Printing multiversion label ...")
	imageId := engine.VDriver.GetImageID(tag)
	multiVersion, _ := engine.VDriver.GetImageLabel(imageId, config.APP_MULTIVERSION_ENV)
	if multiVersion != config.MULTIVERSION_TRUE_FLAG {
		multiVersion = config.MULTIVERSION_FALSE_FLAG
	}
	decorate.PrintLabel(multiVersion)
	decorate.Println()
}