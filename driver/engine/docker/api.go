package docker

import "github.com/docker/docker/client"

type provider struct {
	client *client.Client
}

func NewProvider() *provider {
	return &provider{}
}