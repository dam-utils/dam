package internal

import (
	"dam/config"
	"dam/driver/engine"
	"strings"
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

	imageFamily, ok := engine.VDriver.GetImageLabel(imageId, config.APP_FAMILY_ENV)
	_, imageName, _ := SplitTag(tag)

	if !ok {
		imageFamily = imageName
	}

	return imageFamily
}
