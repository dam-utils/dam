package internal

import (
	"context"
	"io"
	"io/ioutil"
	"strings"

	"dam/driver/conf/option"
	fs "dam/driver/filesystem"
	"dam/driver/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func GetImagesSum() []types.ImageSummary {
	cli, err := client.NewClientWithOpts(client.WithVersion(option.Config.Docker.GetAPIVersion()))
	defer func() {
		if cli != nil {
			cli.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot create new docker client")
	}

	var opts = types.ImageListOptions{}
	imageSum, err := cli.ImageList(context.Background(), opts)
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
