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
package docker

import (
	"dam/driver/structures"
)

type provider struct {
	//Build(imageTag, projectDir string)
	//LoadImage(file string)
	//Pull(tag string, repo *structures.Repo)
	//Images() []*string
	//GetImageID(tag string) string
	//GetImageLabel(tag, labelName string) string
	//SaveImage(imageId, filePath string)
	//ContainerCreate(image string, name string) string
	//CopyFromContainer(containerID, sourcePath, destPath string)
	//ContainerRemove(id string)
	//SearchAppNames(mask string) *[]string
}

func NewProvider() *provider {
	return &provider{}
}

func (p *provider) SearchAppNames(mask string) *[]string {
	return searchAppNames(mask)
}

func (p *provider) LoadImage(file string) {
	loadImage(file)
}

func (p *provider) Pull(tag string, repo *structures.Repo) {
	pull(tag, repo)
}

func (p *provider) GetImageID(tag string) string {
	return getImageID(tag)
}

func (p *provider) GetImageLabel(tag, labelName string) string {
	return getImageLabel(tag, labelName)
}

func (p *provider) SaveImage(imageId, filePath string) {
	saveImage(imageId, filePath)
}

func (p *provider) ContainerCreate(image string, name string) string {
	return containerCreate(image, name)
}

func (p *provider) CopyFromContainer(containerID, sourcePath, destPath string) {
	copyFromContainer(containerID, sourcePath, destPath)
}

func (p *provider) ContainerRemove(id string) {
	containerRemove(id)
}

func (p *provider) Build(imageTag, projectDir string) {
	build(imageTag, projectDir)
}

func (p *provider) Images() *[]string {
	return images()
}