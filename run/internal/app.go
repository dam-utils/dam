package internal

import (
	"strings"

	"dam/config"
	"dam/driver/engine"
	"dam/driver/logger"
)

func SplitTag(tag string) (string, string, string) {
	n := strings.Split(tag, "/")
	nameWithVersion := n[len(n)-1]
	server := strings.Join(n[:len(n)-1], "/")

	v := strings.Split(nameWithVersion, ":")
	version := v[len(v)-1]
	name := v[0]

	return server, name, version
}

func GetFamily(tag string) string {
	imageId := engine.VDriver.GetImageID(tag)
	if imageId == "" {
		logger.Fatal("Image with tag '%s' not exist in the system", tag)
	}

	imageFamily, ok := engine.VDriver.GetImageLabel(imageId, config.APP_FAMILY_ENV)
	_, imageName, _ := SplitTag(tag)

	if !ok {
		imageFamily = imageName
	}

	return imageFamily
}

func GetMultiVersion(tag string) bool {
	imageId := engine.VDriver.GetImageID(tag)
	if imageId == "" {
		logger.Fatal("Image with tag '%s' not exist in the system", tag)
	}

	imageMultiVersion, ok := engine.VDriver.GetImageLabel(imageId, config.APP_MULTIVERSION_ENV)

	if !ok {
		imageMultiVersion = config.MULTIVERSION_FALSE_FLAG
	}

	return imageMultiVersion == config.MULTIVERSION_TRUE_FLAG
}

func GetServers(tag string) string {
	imageId := engine.VDriver.GetImageID(tag)
	if imageId == "" {
		logger.Fatal("Image with tag '%s' not exist in the system", tag)
	}

	servers, ok := engine.VDriver.GetImageLabel(imageId, config.APP_SERVERS_ENV)

	if !ok {
		logger.Warn("Label APP_SERVERS")
		servers = ""
	}

	return servers
}