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
package run

type CreateAppSettings struct {
	MetaPath       string
	DockerFilePath string
	EnvFilePath    string
}

var CreateAppFlags = new(CreateAppSettings)

func CreateApp(){
	//app := new(storage.App)
	//
	//pwd := getCurrentDir()
	//// Validate filesystem
	//metaDir := checkMeta(CreateAppFlags.MetaPath, pwd)
	//dockerFile := checkDockerfile(CreateAppFlags.DockerFilePath, pwd)
	//
	//// Create environment map
	//envs := checkEnvironment(metaDir, CreateAppFlags.EnvFilePath)
	//
	//// Пока не знаю заменять ли текущую мета или генерировать новую и добавлять в Dockerfile
	//releaseMetaPath := createMeta(metaDir, envs)
	//
	//app = createApp(releaseMetaPath, dockerFile, envs)
	//
	//// Написать, что приложение успешно создано
}