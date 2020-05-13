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
package project

import (
	"os"

	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/filesystem/dockerfile"
	"dam/driver/logger"
)

func Prepare(path string)(string, string, string){
	meta := path+string(os.PathSeparator)+config.META_DIR_NAME
	if !fs.IsExistDir(meta) {
		logger.Fatal("Cannot found 'meta' for path: "+meta)
	}

	dockerFile := path+string(os.PathSeparator)+config.DOCKERFILE_NAME
	if !fs.IsExistFile(dockerFile) {
		logger.Fatal("Cannot found 'meta' for path: "+dockerFile)
	}

	install := meta+string(os.PathSeparator)+config.INSTALL_FILE_NAME
	if !fs.IsExistFile(install) {
		if !fs.IsExistFile(install+config.EXPAND_META_FILE) {
			logger.Fatal("Cannot found '" + install + "' or " + install + config.EXPAND_META_FILE + " files in meta directory")
		}
	}

	uninstall := meta+string(os.PathSeparator)+config.UNINSTALL_FILE_NAME
	if !fs.IsExistFile(uninstall) {
		if !fs.IsExistFile(uninstall+config.EXPAND_META_FILE) {
			logger.Fatal("Cannot found '" + uninstall + "' or " + uninstall + config.EXPAND_META_FILE + " files in meta directory")
		}
	}

	if !dockerfile.IsCopyMeta(dockerFile) {
		logger.Fatal("Not found COPY or ADD .. /meta command in Dockerfile: "+ dockerFile)
	}

	if !dockerfile.IsFamily(dockerFile) {
		logger.Warn("Not found label 'FAMILY' in Dockerfile: "+ dockerFile)
	}

	return meta, dockerFile, path+string(os.PathSeparator)+config.ENV_FILE_NAME
}

func ValidateTag(tag string) {

}