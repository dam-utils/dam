package docker

import (
	"context"
	"os"

	"dam/driver/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/archive"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
)

func (p *provider) Build(imageTag, projectDir string, labels map[string]string) {
	logger.Debug("Building image with parameters: imageTag '%s', projectDir '%s', labels '%v'", imageTag, projectDir, labels)
	buildCtx, err := archive.TarWithOptions(projectDir, &archive.TarOptions{})
	if err != nil {
		logger.Fatal("Cannot create docker context (project files directory) with error: %s", err)
	}
	opts := types.ImageBuildOptions{
		Tags: []string{imageTag},
		Context : buildCtx,
		Labels: labels,

		//может пригодиться
		//PullParent: true,
		//BuildArgs: map[string]*string,
		//Platform: string,
	}

	p.connect()
	defer p.close()

	resp, err := p.client.ImageBuild(context.Background(), buildCtx, opts)
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
