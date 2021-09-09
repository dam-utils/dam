package run

import (
	"dam/driver/conf/option"
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
	createEnvs "dam/run/internal/create/envs"
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

	logger.Debug("Flags validated with success")
	if CreateAppFlags.Family != "" {
		flag.ValidateFamily(CreateAppFlags.Family)
	}

	logger.Debug("Preparing envs ...")
	projectDir := fs.GetAbsolutePath(path)
	metaDir, dockerFile, envFile := project.Prepare(projectDir)

	// Create environment map
	envs := combineEnvs(envFile, dockerFile)
	logger.Debug("Envs: '%v'", envs)
	envStorage := createEnvs.NewStorage(envs)
	// Строгая последовательность инициализации
	envStorage.InitAppName(option.Config.ReservedEnvs.GetDefaultAppName(), CreateAppFlags.Name)
	envStorage.InitAppVersion(option.Config.ReservedEnvs.GetDefaultAppVersion(), CreateAppFlags.Version)
	envStorage.InitAppFamily(CreateAppFlags.Family)
	envStorage.InitAppMultiversion(internal.BoolToString(CreateAppFlags.MultiVersion))
	defRepo := getRepo()
	envStorage.InitAppTag(defRepo)
	envStorage.InitAppServers(defRepo)

	logger.Debug("Preparing metaDir ...")
	meta.PrepareExpFiles(metaDir, envStorage.Envs())
	meta.PrepareExecFiles(metaDir)

	logger.Debug("Building image ...")
	engine.VDriver.Build(envStorage.Tag(), projectDir, envStorage.Labels())

	logger.Success("App '%s' was created.", envStorage.Tag())
}

func getRepo() string {
	r := db.RDriver.GetDefaultRepo()
	if r == nil {
		logger.Fatal("Internal error. Not found default repo")
	}

	if r.Id == structures.OfficialRepo.Id {
		return ""
	} else {
		return r.Server
	}
}

// Приоритеты замещения переменных по убыванию:
//- файл ENVIRONMENT
//- Dockerfile
//- переменных окружения, начинающихся с config.OS_ENV_PREFIX
func combineEnvs(envFile string, dockerFile string) map[string]string {
	dfEnv := env.GetDockerFileEnv(dockerFile)
	logger.Debug("Find Dockerfile envs: %v", dfEnv)
	osEnv := env.GetOSEnv(option.Config.ReservedEnvs.GetOSEnvPrefix())
	logger.Debug("Find OS envs: %v", osEnv)

	if envFile == "" {
		return env.MergeEnvs(osEnv, dfEnv)
	}

	fEnv := env.GetFileEnv(envFile)
	logger.Debug("Find Environment file envs: %v", fEnv)
	return env.MergeEnvs(env.MergeEnvs(osEnv, dfEnv), fEnv)
}
