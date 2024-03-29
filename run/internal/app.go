package internal

import (
	"strings"

	"dam/driver/conf/option"
	"dam/driver/engine"
	"dam/driver/logger"
)

func SplitTag(tag string) (string, string, string) {
	n := strings.Split(tag, "/")
	nameWithVersion := n[len(n)-1]
	server := strings.Join(n[:len(n)-1], "/")

	v := strings.Split(nameWithVersion, ":")
	if len(v) != 2 {
		logger.Debug("Cannot found one name and one version app in string '%s'", v)
		return "", "", ""
	}
	version := v[1]
	name := v[0]

	return server, name, version
}

func GetFamily(tag string) string {
	imageId := engine.VDriver.GetImageID(tag)
	if imageId == "" {
		logger.Fatal("Image with tag '%s' not exist in the system", tag)
	}

	imageFamily, ok := engine.VDriver.GetImageLabel(imageId, option.Config.ReservedEnvs.GetAppFamilyEnv())
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

	imageMultiVersion, ok := engine.VDriver.GetImageLabel(imageId, option.Config.ReservedEnvs.GetAppMultiversionEnv())

	if !ok {
		imageMultiVersion = option.Config.Multiversion.GetFalseFlag()
	}

	return imageMultiVersion == option.Config.Multiversion.GetTrueFlag()
}

func GetServersByTag(tag string) string {
	imageId := engine.VDriver.GetImageID(tag)
	if imageId == "" {
		logger.Fatal("Image with tag '%s' not exist in the system", tag)
	}

	servers, ok := engine.VDriver.GetImageLabel(imageId, option.Config.ReservedEnvs.GetAppServersEnv())

	if !ok {
		logger.Warn("Label APP_SERVERS is empty")
		servers = ""
	}

	return servers
}
