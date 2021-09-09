package engine

import (
	"dam/driver/conf/option"
	"dam/driver/engine/docker"
	"dam/driver/logger"
	"dam/driver/structures"
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

type VProvider interface {
	Build(imageTag, projectDir string, labels map[string]string)
	LoadImage(file string)
	Pull(tag string, repo *structures.Repo)
	Images() *[]string
	GetImageID(tag string) string
	GetImageLabel(imageId, labelName string) (string, bool)
	SaveImage(imageId, filePath string)
	ContainerCreate(image string, name string) string
	CopyFromContainer(containerID, sourcePath, destPath string)
	ContainerRemove(id string)
	SearchAppNames(mask string) *[]string
	ImageRemove(dockerID string) bool
	CreateTag(imageId, tag string)
}
