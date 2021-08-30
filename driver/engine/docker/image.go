package docker

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"os"

	"dam/driver/engine/docker/internal"
	"dam/driver/logger"
	"dam/driver/structures"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/pkg/jsonmessage"
	"github.com/docker/docker/pkg/term"
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

	if out.Body != nil && out.JSON {
		outFd, isTerminalOut := term.GetFdInfo(os.Stdout)
		err := jsonmessage.DisplayJSONMessagesStream(out.Body, os.Stdout, outFd, isTerminalOut, nil)
		if err != nil {
			logger.Fatal("Cannot pull docker image with error: %s", err)
		}
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
	logger.Warn("Cannot find images tag '%s' in images list", tag)
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
func (p *provider) GetImageLabel(id, labelName string) (string, bool) {
	p.connect()
	defer p.close()

	var opts = types.ImageListOptions{}
	imageSum, err := p.client.ImageList(context.Background(), opts)
	if err != nil {
		logger.Fatal("Cannot get images list")
	}
	for _, img := range imageSum {
		if internal.PrepareImageID(img.ID) == id {
			for key, value := range img.Labels {
				if key == labelName {
					return value, true
				}
			}
			logger.Warn("Cannot find image label '%s'", labelName)
			return "", false
		}
	}
	logger.Warn("Cannot get image '%s' with label '%s'", id, labelName)
	return "", false
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

func (p *provider) ImageRemove(imageID string) bool {
	p.connect()
	defer p.close()

	var opts = types.ImageRemoveOptions{
		Force:         true,
		PruneChildren: true,
	}

	// response: ([]types.ImageDeleteResponseItem, error)
	_, err := p.client.ImageRemove(context.Background(), imageID, opts)
	if err != nil {
		logger.Warn("Cannot remove image with id '%s' with error: '%s'", imageID, err)
		return false
	}
	return true
}

func (p *provider) CreateTag(imageId, tag string) {
	p.connect()
	defer p.close()

	err := p.client.ImageTag(context.Background(), imageId, tag)
	if err != nil {
		logger.Warn("Cannot create tag '%s' for image id '%s' with error: '%s'", tag, imageId, err)
	}
}