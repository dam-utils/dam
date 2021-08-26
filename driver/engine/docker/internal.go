package docker

import (
	"dam/driver/conf/option"
	"dam/driver/logger"

	"github.com/docker/docker/client"
)

func (p *provider) connect() {
	var err error
	p.client, err = client.NewClientWithOpts(client.WithVersion(option.Config.Docker.GetAPIVersion()))
	if err != nil {
		logger.Fatal("Cannot create a new docker client with error: %s", err)
	}
}

func (p *provider) close() {
	if p.client != nil {
		err := p.client.Close()
		if err != nil {
			logger.Fatal("Cannot close for docker client with error: %s", err)
		}
	}
}