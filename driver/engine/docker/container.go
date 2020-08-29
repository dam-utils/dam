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
	"context"
	"io"
	"os"
	"path/filepath"

	"dam/config"
	"dam/driver/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func containerCreate(image string, name string) string {
	cli, err := client.NewClientWithOpts(client.WithVersion(config.DOCKER_API_VERSION))
	defer func() {
		if cli != nil {
			cli.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create new docker client")
	}

	var conf = container.Config{
		Image: image,
		Cmd:   []string{""},
		Tty:   true, //TODO check it
	}
	resp, err := cli.ContainerCreate(context.Background(), &conf,  nil, nil, name)
	if err != nil {
		logger.Fatal("Cannot build docker image with error: %s", err)
	}

	logger.Debug("Response ContainerCreate('%s'): %v", name, resp)
	return (resp.ID)[:12]
}

func copyFromContainer(containerID, sourcePath, destPath string) {
	cli, err := client.NewClientWithOpts(client.WithVersion(config.DOCKER_API_VERSION))
	defer func() {
		if cli != nil {
			cli.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create new docker client")
	}

	reader, _, err := cli.CopyFromContainer(context.Background(), containerID, sourcePath)
	if err != nil {
		logger.Fatal("Cannot copy from the container with ID '%s' with error: %s", containerID, err)
	}

	// Ex:
	//docker https://github.com/docker/engine-api/issues/308
	tr := tar.NewReader(reader)
	for {
		header, err := tr.Next()

		switch {

		// if no more files are found return
		case err == io.EOF:
			return

		// return any other error
		case err != nil:
			logger.Fatal("Cannot get '%s' from container tar archive with containerID '%s' with error: %s", header.Name, containerID, err)

		// if the header is nil, just skip it (not sure how this happens)
		case header == nil:
			continue
		}

		// the target location where the dir/file should be created
		target := filepath.Join(destPath, header.Name)

		// the following switch could also be done using fi.Mode(), not sure if there
		// a benefit of using one vs. the other.
		// fi := header.FileInfo()

		// check the file type
		switch header.Typeflag {

		// if its a dir and it doesn't exist create it
		case tar.TypeDir:
			if _, err := os.Stat(target); err != nil {
				if err := os.MkdirAll(target, 0755); err != nil {
					logger.Fatal("Cannot create target directory '%s' from containerID '%s' with error: %s", header.Name, containerID, err)
				}
			}

		// if it's a file create it
		case tar.TypeReg:
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				logger.Fatal("Cannot create target file '%s' from containerID '%s' with error: %s", header.Name, containerID, err)
			}

			// copy over contents
			if _, err := io.Copy(f, tr); err != nil {
				logger.Fatal("Cannot write to target file '%s' from containerID '%s' with error: %s", header.Name, containerID, err)
			}

			// manually close here after each file operation; defering would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}

}

func containerRemove(id string) {
	cli, err := client.NewClientWithOpts(client.WithVersion(config.DOCKER_API_VERSION))
	defer func() {
		if cli != nil {
			cli.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create new docker client")
	}

	var opts = types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}
	err = cli.ContainerRemove(context.Background(), id, opts)
	if err != nil {
		logger.Fatal("Cannot remove the container with ID '%s' with error: %s", id, err)
	}
}