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
package docker

import (
	"archive/tar"
	"compress/gzip"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"dam/config"
	"dam/driver/logger"
	"dam/driver/storage"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Pull(tag string, repo *storage.Repo) {
	cli, err := client.NewClientWithOpts(client.WithVersion(config.DOCKER_API_VERSION))
	defer func() {
		if cli != nil {
			cli.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create new docker client")
	}

	authConfig := types.AuthConfig{
		Username: repo.Username,
		Password: repo.Password,
	}
	encodedJSON, err := json.Marshal(authConfig)
	if err != nil {
		panic(err)
	}
	authStr := base64.URLEncoding.EncodeToString(encodedJSON)

	var pullOpts = types.ImagePullOptions{
		//Platform TODO ?
		RegistryAuth: authStr,
	}
	out, err := cli.ImagePull(context.Background(), tag, pullOpts)
	defer func() {
		if out != nil {
			out.Close()
		}
	}()
	if err != nil {
		logger.Warn("Cannot pull docker image with error: %s", err.Error())
		return
	}

	_, err = io.Copy(os.Stdout, out)
	if err != nil {
		logger.Fatal("Cannot print docker stdout with error: %s", err.Error())
	}
}

// TODO refactoring
func GetImageID(tag string) string {
	cli, err := client.NewClientWithOpts(client.WithVersion(config.DOCKER_API_VERSION))
	defer func() {
		if cli != nil {
			cli.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create new docker client")
	}

	var opts = types.ImageListOptions{}
	imageSum, err := cli.ImageList(context.Background(),opts)
	if err != nil {
		logger.Fatal("Cannot get images list")
	}
	for _, img := range imageSum {
		for _, sourceTag := range img.RepoTags {
			if sourceTag == tag {
				return img.ID
			}
		}
	}
	logger.Fatal("Cannot found images tag '%s' in images list", tag)
	return ""
}

// TODO refactoring
func GetImageLabel(tag, labelName string) string {
	cli, err := client.NewClientWithOpts(client.WithVersion(config.DOCKER_API_VERSION))
	defer func() {
		if cli != nil {
			cli.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create new docker client")
	}

	var opts = types.ImageListOptions{}
	imageSum, err := cli.ImageList(context.Background(),opts)
	if err != nil {
		logger.Fatal("Cannot get images list")
	}
	for _, img := range imageSum {
		for _, sourceTag := range img.RepoTags {
			if sourceTag == tag {
				for key, value := range img.Labels {
					if key == labelName {
						return value
					}
				}
				logger.Warn("Cannot found image label '%s'", labelName)
				return ""
			}
		}
	}
	logger.Warn("Cannot found image label '%s'", labelName)
	return ""
}

func SaveImage(imageId, filePath string) {
	cli, err := client.NewClientWithOpts(client.WithVersion(config.DOCKER_API_VERSION))
	defer func() {
		if cli != nil {
			cli.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create new docker client with error: '%s'")
	}

	readCloser, err := cli.ImageSave(context.Background(), []string{imageId})
	defer func() {
		if readCloser != nil {
			readCloser.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot save image with id '%s' with error: '%s'", imageId, err.Error())
	}

	saveToFile(filePath, readCloser)
}

func saveToFile(srcFile string, r io.ReadCloser) {
	f, err := os.Open(srcFile)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open archive file '%s' with error: '%s'", srcFile, err.Error())
	}

	gzf, err := gzip.NewReader(r)
	if err != nil {
		logger.Fatal("Cannot open archive gzip reader for file '%s' with error: '%s'", srcFile, err.Error())
	}
	tarReader := tar.NewReader(gzf)


	for true {
		h, err := tarReader.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Fatal("Cannot write image to gzip archive '%s' with error: '%s'", srcFile, err.Error())
		}
		if h == nil {
			logger.Fatal("Internal error. Header is nil in saving file '%s'", srcFile)
		}
		name := h.Name

		switch h.Typeflag {
		case tar.TypeDir:
			os.Mkdir(name, 0755)
		case tar.TypeReg:
			data := make([]byte, h.Size)
			_, err := tarReader.Read(data)
			if err != nil {
				logger.Fatal("Cannot get data from reader with error: '%s'", err.Error())
			}

			err = ioutil.WriteFile(name, data, 0755)
			if err != nil {
				logger.Fatal("Cannot write data to file '%s' with error: '%s'", name, err.Error())
			}
		default:
			logger.Fatal("Unknown type '%s' for file '%s'", h.Typeflag, name)
		}
	}
}