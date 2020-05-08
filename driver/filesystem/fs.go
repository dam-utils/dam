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
	"dam/driver/logger"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
)

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		logger.Fatal("Cannot get current dir with error: "+ err.Error())
	}
	return dir
}

func DirIsExist(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		logger.Fatal("Cannot check directory "+path+" with error: "+ err.Error())
	}
	return info.IsDir()
}

func FileIsExist(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		logger.Fatal("Cannot check file "+path+" with error: "+ err.Error())
	}
	return !info.IsDir()
}

func GetBaseName(path string) string {
	return filepath.Base(path)
}

func GetFileList(path string, agg *[]string) *[]string {
	if DirIsExist(path) {
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
		logger.Fatal("Cannot check files in path:"+dir+"/ with error: "+ err.Error())
	}
	return files
}

func RemoveFile(path string) {
	err := os.Remove(path)
	if err != nil {
		logger.Fatal("Cannot remove file in path:"+path+" with error: "+ err.Error())
	}
}

func MoveFile(oldLocation, newLocation string) {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		logger.Fatal("Cannot move file '"+oldLocation+"' to '"+newLocation+"' with error: "+ err.Error())
	}
}

func GenerateTmpFilePath() string {
	return GetCurrentDir() + ".tmp."+randSeq(6)
}

func randSeq(n int) string {
	var letters = []rune("abcdefghijklmnopqrstuvwxyz1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CopyFile(sourceFile, destFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		logger.Fatal("Cannot read file '"+sourceFile+"' with error: "+ err.Error())
	}

	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		logger.Fatal("Cannot write to tmp file '"+destFile+"' with error: "+ err.Error())
	}
}
