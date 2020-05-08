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
	"dam/config"
	"dam/driver/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func Build(imageTag, dFile string) {
	opts := types.ImageBuildOptions {
		Tags: []string{imageTag},
		//PullParent: true, надо ли?
		Dockerfile: dFile,
		//BuildArgs: map[string]*string TODO
		//Labels: map[string]string //может пригодиться
		//Platform: string //может пригодиться
	}

	cli, err := client.NewClientWithOpts(client.WithVersion(config.DOCKER_API_VERSION))
	if err != nil {
		logger.Fatal("Cannot create new docker client")
	}

	_, err = cli.ImageBuild(context.Background(), nil, opts)
	if err != nil {
		logger.Fatal("Cannot build docker image with error: " + err.Error())
	}
}
