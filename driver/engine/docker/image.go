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
	"context"
	"dam/driver/engine/docker/internal"
	"dam/driver/structures"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"

	"dam/driver/logger"

	"github.com/docker/docker/api/types"
)

func (p *provider) LoadImage(file string) {
	p.connect()
	defer p.close()

	f, err := os.Open(file)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open the loaded file '%s' with error: %s", file, err)
	}

	out, err := p.client.ImageLoad(context.Background(), f, false)
	defer func() {
		if out.Body != nil {
			out.Body.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot pull docker image with error: %s", err)
	}

	_, err = io.Copy(os.Stdout, out.Body)
	if err != nil {
		logger.Fatal("Cannot print docker stdout with error: %s", err)
	}
}

func (p *provider) Pull(tag string, repo *structures.Repo) {
	p.connect()
	defer p.close()

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
	out, err := p.client.ImagePull(context.Background(), tag, pullOpts)
	defer func() {
		if out != nil {
			out.Close()
		}
	}()
	if err != nil {
		logger.Warn("Cannot pull docker image with error: %s", err)
		return
	}

	_, err = io.Copy(os.Stdout, out)
	if err != nil {
		logger.Fatal("Cannot print docker stdout with error: %s", err)
	}
}

// TODO refactoring
func (p *provider) GetImageID(tag string) string {
	imageSum := internal.GetImagesSum()

	for _, img := range imageSum {
		for _, sourceTag := range img.RepoTags {
			if sourceTag == tag {
				return internal.PrepareImageID(img.ID)
			}
		}
	}
	logger.Fatal("Cannot found images tag '%s' in images list", tag)
	return ""
}

func (p *provider) Images() *[]string {
	result := make([]string, 0)
	imageSum := internal.GetImagesSum()

	for _, img := range imageSum {
		result = append(result, internal.PrepareImageID(img.ID))
	}

	preparedResult := internal.RemoveDuplicates(result)

	return &preparedResult
}

// TODO refactoring
func (p *provider) GetImageLabel(tag, labelName string) string {
	p.connect()
	defer p.close()

	var opts = types.ImageListOptions{}
	imageSum, err := p.client.ImageList(context.Background(),opts)
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

func (p *provider) SaveImage(imageId, filePath string) {
	p.connect()
	defer p.close()

	readCloser, err := p.client.ImageSave(context.Background(), []string{imageId})
	defer func() {
		if readCloser != nil {
			readCloser.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot save image with id '%s' with error: '%s'", imageId, err)
	}

	internal.SaveToFile(filePath, readCloser)
}

