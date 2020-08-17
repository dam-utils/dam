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
	"dam/driver/docker"
	"dam/driver/docker/manifest"
	fs "dam/driver/filesystem"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/run/internal"
)

type SaveSettings struct {
	FilePath string
}

var SaveFlags = new(SaveSettings)

func Save(appFullName string) {
	var filePath, resultPrefixPath string

	flag.ValidateAppPlusVersion(appFullName)
	logger.Debug("Flags validated with success")

	logger.Debug("Parsing tag ...")
	_, name, version := internal.SplitTag(appFullName)

	logger.Debug("Getting archive path ...")
	// TODO refactoring
	if SaveFlags.FilePath != "" {
		flag.ValidateFilePath(SaveFlags.FilePath)
		filePath = SaveFlags.FilePath
	} else {
		filePath = fs.GetCurrentDir()+
			string(os.PathSeparator)+
			name+
			config.SAVE_FILE_SEPARATOR+
			version+config.SAVE_TMP_FILE_POSTFIX
		resultPrefixPath =  fs.GetCurrentDir()+
			string(os.PathSeparator)+
			name+
			config.SAVE_FILE_SEPARATOR+
			version+
			config.SAVE_OPTIONAL_SEPARATOR
	}

	logger.Debug("Saving archive ...")
	docker.SaveImage(docker.GetImageID(appFullName), filePath)

	logger.Debug("Preparing manifest ...")
	modifyManifest(filePath, appFullName)

	logger.Debug("Releasing archive ...")
	if SaveFlags.FilePath == "" {
		resultPath := resultPrefixPath+fs.HashFileCRC32(filePath)+config.SAVE_FILE_SEPARATOR+fs.FileSize(filePath)+config.SAVE_FILE_POSTFIX
		fs.MoveFile(filePath, resultPath)
		logger.Success("Created '%s' file.", resultPath)
	}
}

func modifyManifest(filePath, appFullName string) {
	dir := fs.Untar(filePath)
	manifestFile := dir+string(os.PathSeparator)+config.SAVE_MANIFEST_FILE

	manifest.ModifyRepoTags(manifestFile, appFullName)
	fs.Remove(filePath)
	fs.Gzip(dir, filePath)
	fs.Remove(dir)
}
