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
	"bytes"
	"dam/config"
	"encoding/hex"
	"hash/crc32"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"dam/driver/logger"

	"github.com/docker/docker/pkg/system"
)

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		logger.Fatal("Cannot get current dir with error: %s", err)
	}
	return dir
}

func IsExistDir(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		logger.Debug("Cannot check directory '%s' with error: %s", path, err)
		return false
	}
	return info.IsDir()
}

func IsExistFile(path string) bool {
	info, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !info.IsDir()
}

func GetBaseName(path string) string {
	return filepath.Base(path)
}

func GetFileList(path string, agg []string) []string {
	if IsExistDir(path) {
		for _, p := range Ls(path) {
			return append(agg, GetFileList(p, agg)...)
		}
	} else {
		return append(agg, path)
	}
	return agg
}

func Ls(dir string) []string {
	files, err := filepath.Glob(dir + string(os.PathSeparator) + "*")
	if err != nil {
		logger.Fatal("Cannot check files in path '%s' with error: %s", dir, err)
	}
	return files
}

func MkDir(dir string) {
	err := os.MkdirAll(dir, 0755)
	if err != nil {
		logger.Fatal("Cannot create directory '%s/' with error: %s", dir, err)
	}
}

func Remove(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		logger.Warn("Cannot check path '%s' with error: %s", path, err)
	}

	err = os.RemoveAll(path)
	if err != nil {
		logger.Warn("Cannot remove path '%s' with error: %s", path, err)
	}
	return false
}

func MoveFile(oldLocation, newLocation string) {
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		logger.Fatal("Cannot move file '%s' to '%s' with error: %s", oldLocation, newLocation, err)
	}
}

func CopyFile(sourceFile, destFile string) {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		logger.Fatal("Cannot read file '%s' with error: %s", sourceFile, err)
	}

	err = ioutil.WriteFile(destFile, input, 0644)
	if err != nil {
		logger.Fatal("Cannot write to tmp file '%s' with error: %s", destFile, err)
	}
}

func GetAbsolutePath(path string) string {
	p, err := filepath.Abs(path)
	{
		if err != nil {
			logger.Fatal("Cannot get absolute path for '%s' with error: %s", path, err)
		}
	}
	return p
}

// TODO заменить на ChmodPlusX()
func Chmod777(path string) {
	if err := os.Chmod(path, 0777); err != nil {
		logger.Fatal("Cannot chmod 777 '%s' with error: %s", path, err)
	}
}

func Chdir(path string) {
	err := os.Chdir(path)
	if err != nil {
		logger.Fatal("Cannot change home dir to '%s' with error: %s", path, err)
	}
}

// https://stackoverflow.com/questions/40670228/how-to-run-binary-files-inside-golang-program
func RunFile(runFile string) {
	pwd := GetCurrentDir()
	defer Chdir(pwd)

	homeDir := filepath.Dir(runFile)
	Chdir(homeDir)

	c := exec.Command(runFile)
	c.Dir = homeDir   //TODO delete?
	// set var to get the output
	var outb, errb bytes.Buffer

	// set the output to our variable
	c.Stdout = &outb
	c.Stderr = &errb
	err := c.Run()
	if err != nil {
		logger.Warn(errb.String())
		logger.Fatal("Cannot execute file '%s' with error: %s", runFile, err)
	}
	logger.Info(outb.String())
}

func Touch(file string) {
	if !IsExistFile(file) {
		emptyFile, err := os.Create(file)
		defer func() {
			if emptyFile != nil {
				emptyFile.Close()
			}
		}()
		if err != nil {
			logger.Fatal("Cannot create file '%s' with error: %s", file, err)
		}
	}
}

func HashFileCRC32(filePath string) string {
	f, err := os.Open(filePath)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open file '%s' with error: %s", filePath, err)
	}

	tablePolynomial := crc32.MakeTable(config.SAVE_POLYNOMIAL_CKSUM)
	hash := crc32.New(tablePolynomial)
	if _, err := io.Copy(hash, f); err != nil {
		logger.Fatal("Cannot check hash file '%s' with error: %s", filePath, err)
	}
	hashInBytes := hash.Sum(nil)[:]
	return hex.EncodeToString(hashInBytes)
}

func FileSize(filePath string) string {
	fi, err := os.Stat(filePath)
	if err != nil {
		logger.Fatal("Cannot check file '%s' with error: %s", filePath, err)
	}

	return strconv.FormatInt(fi.Size(), 10)
}

func EraceDataCreation(path string) {
	// https://github.com/moby/moby/blob/e9b4655bc98563602d961c72fc62cb20cc143515/image/tarexport/save.go#L187
	if err := system.Chtimes(path, time.Unix(0, 0), time.Unix(0, 0)); err != nil {
		logger.Fatal("Cannot erase metadata for manifest file with error: %s", err)
	}
}