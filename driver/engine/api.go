package engine

import (
	"dam/driver/conf/option"
	"dam/driver/engine/docker"
	"dam/driver/logger"
)

var (
	VDriver VProvider
)

func Init() {
	switch option.Config.Virtualization.GetType() {
	case "docker":
		VDriver = docker.NewProvider()
	default:
		logger.Fatal("Config option VIRTUALIZATION_TYPE='%s' not valid. DB type is bad", option.Config.Virtualization.GetType())
	}
}
