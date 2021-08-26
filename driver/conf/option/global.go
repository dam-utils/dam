package option

import "dam/config"

type Global struct {}

func (o *Global) GetProjectName() string {
	return config.PROJECT_NAME
}

func (o *Global) GetProjectVersion() string {
	return config.PROJECT_VERSION
}
