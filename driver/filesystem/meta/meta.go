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
package meta

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"

	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/filesystem/env"
	"dam/driver/logger"
)

func PrepareExpFiles(metaDir string, envs map[string]string) {
	files := fs.GetFileList(metaDir, &[]string{})
	for _, file := range *files {
		if strings.HasSuffix(file, config.EXPAND_META_FILE){
			prepareExpFile(file, envs)
		}
	}
}

func prepareExpFile(path string, envs map[string]string) {
	newPath := path[:4]

	f, err := os.Open(path)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open file '%s' with error: %s", path, err.Error())
	}

	newf, err := os.Create(newPath)
	defer func() {
		if newf != nil {
			f.Close()
		}
	}()
	if err != nil {
		newf.Close()
		logger.Fatal("Cannot create file '%s' with error: %s", newPath, err.Error())
	}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		newString := env.PrepareExpString(scanner.Text(), envs)
		_, err = f.WriteString(newString)
		if err != nil {
			logger.Fatal("Cannot write string to file '%s' with error: %s", newPath, err.Error())
		}
	}

	err = newf.Sync()
	if err != nil {
		logger.Fatal("Cannot sync of writable file '%s' with error: %s", newPath, err.Error())
	}

	err = newf.Close()
	if err != nil {
		logger.Fatal("Cannot close of writable file '%s' with error: %s", newPath, err.Error())
	}
}

func PrepareExecFiles(meta string) {
	fs.Chmod777(filepath.Join(meta, config.INSTALL_FILE_NAME))
	fs.Chmod777(filepath.Join(meta, config.UNINSTALL_FILE_NAME))
}
