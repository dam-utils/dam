package option

import "dam/config"

type Docker struct {}

func (o *Docker) GetAPIVersion() string {
	return config.DOCKER_API_VERSION
}
