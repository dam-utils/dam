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
	"dam/driver/filesystem/env"
	"dam/driver/filesystem/meta"
	"dam/driver/filesystem/project"
	"dam/driver/logger"
	"dam/driver/storage"
)

type CreateAppSettings struct {

}

var CreateAppFlags = new(CreateAppSettings)

func CreateApp(path string) {
	projectDir := fs.GetAbsolutePath(path)
	metaDir, dockerFile, envFile := project.Prepare(projectDir)

	// Create environment map
	envs := combineEnvs(envFile, dockerFile)
	preparedEnvs := env.PrepareProjectEnvs(envs)

	meta.PrepareExpFiles(metaDir, preparedEnvs)

	tag := getImageTag(preparedEnvs)
	project.ValidateTag(tag)

	docker.Build(getImageTag(preparedEnvs), projectDir)
	logger.Info("App '"+tag+"' was created.")
}

func getImageTag(envs map[string]string) string {
	var tag string

	r := db.RDriver.GetDefaultRepo()
	if r.Id == storage.OfficialRepo.Id {
		tag = envs[config.APP_NAME_ENV]+":"+envs[config.APP_VERS_ENV]
	} else {
		tag = r.Name+"/"+envs[config.APP_NAME_ENV]+":"+envs[config.APP_VERS_ENV]
	}
	return tag
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