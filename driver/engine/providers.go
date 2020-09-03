// Copyright 2020 The Docker Applications Manager Authors
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
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
	GetImageLabel(tag, labelName string) string
	SaveImage(imageId, filePath string)
	ContainerCreate(image string, name string) string
	CopyFromContainer(containerID, sourcePath, destPath string)
	ContainerRemove(id string)
	SearchAppNames(mask string) *[]string
	ImageRemove(dockerID string) bool
}

