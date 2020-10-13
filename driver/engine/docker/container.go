package docker

import (
	"archive/tar"
	"context"
	"io"
	"os"
	"path/filepath"

	"dam/driver/logger"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

func (p *provider) ContainerCreate(image string, name string) string {
	p.connect()
	defer p.close()

	var conf = container.Config{
		Image: image,
		Cmd:   []string{""},
		Tty:   true, //TODO check it
	}
	resp, err := p.client.ContainerCreate(context.Background(), &conf,  nil, nil, name)
	if err != nil {
		logger.Fatal("Cannot build docker image with error: %s", err)
	}

	logger.Debug("Response ContainerCreate('%s'): %v", name, resp)
	return (resp.ID)[:12]
}

func (p *provider) CopyFromContainer(containerID, sourcePath, destPath string) {
	p.connect()
	defer p.close()

	reader, _, err := p.client.CopyFromContainer(context.Background(), containerID, sourcePath)
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

			// manually close here after each file operation; deferring would cause each file close
			// to wait until all operations have completed.
			f.Close()
		}
	}

}

func (p *provider) ContainerRemove(id string) {
	p.connect()
	defer p.close()

	var opts = types.ContainerRemoveOptions{
		RemoveVolumes: true,
		Force:         true,
	}
	err := p.client.ContainerRemove(context.Background(), id, opts)
	if err != nil {
		logger.Fatal("Cannot remove the container with ID '%s' with error: %s", id, err)
	}
}