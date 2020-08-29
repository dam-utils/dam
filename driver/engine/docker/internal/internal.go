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
package internal

import (
	"context"
	"io"
	"io/ioutil"
	"strings"

	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetImagesSum() []types.ImageSummary {
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

	return imageSum
}

// From https://www.dotnetperls.com/duplicates-go
func RemoveDuplicates(elements []string) []string {
	encountered := map[string]bool{}
	result := make([]string, 0)

	for v := range elements {
		if encountered[elements[v]] != true {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}

	return result
}

// Incoming format: 'sha256:767d33...'
func PrepareImageID(id string) string {
	arr := strings.Split(id, ":")
	return arr[1][0:12]
}

func SaveToFile(srcFile string, r io.ReadCloser) {
	fs.Touch(srcFile)

	content, err := ioutil.ReadAll(r)
	if err != nil {
		logger.Fatal("Cannot open reader for file '%s' with error: '%s'", srcFile, err)
	}

	err = ioutil.WriteFile(srcFile, content, 0644)
	if err != nil {
		logger.Fatal("Cannot write image to file '%s' with error: '%s'", srcFile, err)
	}
}