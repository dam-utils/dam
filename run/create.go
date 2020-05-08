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

import (
	"dam/config"
	"dam/driver/db"
	"dam/driver/docker"
	fs "dam/driver/filesystem"
	"dam/driver/filesystem/dockerfile"
	"dam/driver/filesystem/env"
	"dam/driver/filesystem/meta"
	"dam/driver/logger"
)

type CreateAppSettings struct {
	MetaPath       string
	DockerFilePath string
	EnvFilePath    string
}

var CreateAppFlags = new(CreateAppSettings)

func CreateApp() {
	pwd := fs.GetCurrentDir()
	// Validate filesystem
	metaDir := meta.GetPath(CreateAppFlags.MetaPath, pwd)
	dockerFile := dockerfile.GetPath(CreateAppFlags.DockerFilePath, pwd)

	// Create environment map
	envFile := env.GetPath(metaDir, CreateAppFlags.EnvFilePath)
	envs := combineEnvs(envFile, dockerFile)
	preparedEnvs := env.PrepareAppVers(env.PrepareAppName(envs))

	meta.PrepareExpFiles(metaDir, preparedEnvs)
	tag := getImageTag(preparedEnvs)
	buildImage(dockerFile, metaDir, getImageTag(preparedEnvs))

	logger.Info("App "+tag+" was created.")
}

func buildImage(dockerFile, meta, imageTag string) {
	tmpFile := fs.GenerateTmpFilePath() + ".Dockerfile"

	fs.CopyFile(dockerFile, tmpFile)
	defer fs.RemoveFile(tmpFile)

	dockerfile.PrepareCopyMeta(tmpFile, meta, CreateAppFlags.MetaPath)
	docker.Build(imageTag, tmpFile)
}

func getImageTag(envs map[string]string) string {
	r := db.RDriver.GetDefaultRepo()
	return r.Name+"/"+envs[config.DEF_APP_NAME]+":"+envs[config.DEF_APP_VERS]
}

// Приоритеты замещения переменных по убыванию:
//- файл ENVIRONMENT
//- Dockerfile
//- переменных окружения, начинающихся с config.OS_ENV_PREFIX
func combineEnvs(envFile string, dockerFile string) map[string]string {
	dfEnv := env.GetDockerFileEnv(dockerFile)
	osEnv := env.GetOSEnv(config.OS_ENV_PREFIX)

	if envFile == "" {
		return env.MergeEnvs(osEnv, dfEnv)
	}

	fEnv := env.GetFileEnv(envFile)
	return env.MergeEnvs(env.MergeEnvs(osEnv, dfEnv), fEnv)
}