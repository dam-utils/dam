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
	"os"

	"dam/config"
	"dam/driver/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
)

func Build(imageTag, projectDir string) {
	buildCtx, err := archive.TarWithOptions(projectDir, &archive.TarOptions{})
	if err != nil {
		logger.Fatal("Cannot create docker context (project files directory) with error: %s", err)
	}
	opts := types.ImageBuildOptions{
		Tags: []string{imageTag},
		Context : buildCtx,

		//может пригодиться
		//PullParent: true,
		//BuildArgs: map[string]*string,
		//Labels: map[string]string,
		//Platform: string,
	}

	cli, err := client.NewClientWithOpts(client.WithVersion(config.DOCKER_API_VERSION))
	defer func() {
		if cli != nil {
			cli.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create new docker client")
	}

	resp, err := cli.ImageBuild(context.Background(), buildCtx, opts)
	defer func() {
		if resp.Body != nil {
			resp.Body.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot build docker image with error: %s", err)
	}


	termFd, isTerm := term.GetFdInfo(os.Stderr)
	err = jsonmessage.DisplayJSONMessagesStream(resp.Body, os.Stderr, termFd, isTerm, nil)
	if err != nil {
		logger.Fatal("Cannot get output json for building image with error: %s", err)
	}
}