package run

import (
	"dam/config"
	"dam/driver/db"
	"dam/driver/engine"
	fs "dam/driver/filesystem"
	"dam/driver/filesystem/env"
	"dam/driver/filesystem/meta"
	"dam/driver/filesystem/project"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/driver/structures"
	"dam/run/internal"
)

type CreateAppSettings struct {
	Name   string
	Version string
	Family string
	MultiVersion bool
}

var CreateAppFlags = new(CreateAppSettings)

func CreateApp(path string) {
	flag.ValidateProjectDirectory(path)
	flag.ValidateAppName(CreateAppFlags.Name)
	flag.ValidateAppVersion(CreateAppFlags.Version)
	logger.Debug("Flags validated with success")

	logger.Debug("Preparing labels ...")
	labels := make(map[string]string)
	if CreateAppFlags.Family == "" {
		labels[config.APP_FAMILY_ENV]=CreateAppFlags.Name
	} else {
		flag.ValidateFamily(CreateAppFlags.Family)
		labels[config.APP_FAMILY_ENV]=CreateAppFlags.Family
	}

	labels[config.APP_MULTIVERSION_ENV]=internal.BoolToString(CreateAppFlags.MultiVersion)

	logger.Debug("Preparing envs ...")
	projectDir := fs.GetAbsolutePath(path)
	metaDir, dockerFile, envFile := project.Prepare(projectDir)

	// Create environment map
	envs := combineEnvs(envFile, dockerFile)
	preparedEnvs := env.PrepareProjectEnvs(envs)
	preparedEnvs = setEnvFlag(preparedEnvs, config.APP_NAME_ENV, CreateAppFlags.Name)
	preparedEnvs = setEnvFlag(preparedEnvs, config.APP_VERS_ENV, CreateAppFlags.Version)

	logger.Debug("Preparing metaDir ...")
	meta.PrepareExpFiles(metaDir, preparedEnvs)
	meta.PrepareExecFiles(metaDir)

	logger.Debug("Preparing tag ...")
	tag := getImageTag(preparedEnvs)
	project.ValidateTag(tag)

	logger.Debug("Building image ...")
	engine.VDriver.Build(getImageTag(preparedEnvs), projectDir, labels)

	logger.Success("App '%s' was created.", tag)
}

func getImageTag(envs map[string]string) string {
	var tag string

	r := db.RDriver.GetDefaultRepo()
	if r == nil {
		logger.Fatal("Internal error. Not found default repo")
	}

	if r.Id == structures.OfficialRepo.Id {
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

func setEnvFlag(envs map[string]string, env, envFlag string) map[string]string {
	if envFlag != "" {
		envs[env] = envFlag
	}
	return envs
}