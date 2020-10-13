package engine

import (
	"dam/config"
	"dam/driver/engine/docker"
	"dam/driver/logger"
)

var (
	VDriver VProvider
)

func Init() {
	switch config.VIRTUALIZATION_TYPE {
	case "docker":
		VDriver = docker.NewProvider()
	default:
		logger.Fatal("Config option VIRTUALIZATION_TYPE='%s' not valid. DB type is bad", config.VIRTUALIZATION_TYPE)
	}
}
