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
package filesystem

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"dam/driver/logger"
)

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		logger.Fatal("Cannot get current dir with error: %s", err.Error())
	}
	return dir
}

func IsExistDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		logger.Fatal("Cannot check directory '%s' with error: %s", path, err.Error())
	}
	return info.IsDir()
}

func IsExistFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		logger.Fatal("Cannot check file '%s' with error: %s", path, err.Error())
	}
	return !info.IsDir()
}

func GetBaseName(path string) string {
	return filepath.Base(path)
}

func GetFileList(path string, agg *[]string) *[]string {
	if IsExistDir(path) {
		for _, p := range Ls(path) {
			GetFileList(p, agg)
		}
	} else {
		newAgg := append(*agg, path)
		return &newAgg
	}
	return agg
}

func Ls(dir string) []string {
	files, err := filepath.Glob("dir/*")
	if err != nil {
		logger.Fatal("Cannot check files in path '%s/' with error: %s", dir, err.Error())
	}
	return files
}

func Remove(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		logger.Warn("Cannot check path '%s' with error: %s", path, err.Error())
	}

	err = os.RemoveAll(path)
	if err != nil {
		logger.Warn("Cannot remove path '%s' with error: %s", path, err.Error())
	}
	return false
}

func MoveFile(oldLocation, newLocation string) {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		logger.Fatal("Cannot move file '%s' to '%s' with error: %s", oldLocation, newLocation, err.Error())
	}
}

func CopyFile(sourceFile, destFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		logger.Fatal("Cannot read file '%s' with error: %s", sourceFile, err.Error())
	}

	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		logger.Fatal("Cannot write to tmp file '%s' with error: %s", destFile, err.Error())
	}
}

func GetAbsolutePath(path string) string {
	p, err := filepath.Abs(path)
	{
		if err != nil {
			logger.Fatal("Cannot get absolute path for '%s' with error: %s", path, err.Error())
		}
	}
	return p
}

// TODO заменить на ChmodPlusX()
func Chmod777(path string) {
	if err := os.Chmod(path, 0777); err != nil {
		logger.Fatal("Cannot chmod 777 '%s' with error: %s", path, err.Error())
	}
}
