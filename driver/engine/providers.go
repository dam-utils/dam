package engine

import (
	"dam/driver/structures"
)

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

