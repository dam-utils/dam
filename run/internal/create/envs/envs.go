package envs

import (
	"dam/driver/conf/option"
	"dam/driver/logger"
	"dam/run/internal/label/servers"
)

type env struct {
	data map[string]string
}

func NewStorage(m map[string]string) *env {
	storage := new(env)
	storage.data = m
	return storage
}

func (e *env) InitAppName(def, flag string) {
	if flag != "" {
		e.data[option.Config.ReservedEnvs.GetAppNameEnv()] = flag
	}

	if e.data[option.Config.ReservedEnvs.GetAppNameEnv()] == "" {
		e.data[option.Config.ReservedEnvs.GetAppNameEnv()] = def
	}
}

func (e *env) InitAppVersion(def, flag string) {
	if flag != "" {
		e.data[option.Config.ReservedEnvs.GetAppVersionEnv()] = flag
	}

	if e.data[option.Config.ReservedEnvs.GetAppVersionEnv()] == "" {
		e.data[option.Config.ReservedEnvs.GetAppVersionEnv()] = def
	}
}

func (e *env) InitAppMultiversion(flag string) {
	if flag != "" {
		e.data[option.Config.ReservedEnvs.GetAppMultiversionEnv()] = flag
	}

	if e.data[option.Config.ReservedEnvs.GetAppMultiversionEnv()] == "" {
		e.data[option.Config.ReservedEnvs.GetAppMultiversionEnv()] = "false"
	}
}

func (e *env) InitAppFamily(flag string) {
	if flag != "" {
		e.data[option.Config.ReservedEnvs.GetAppFamilyEnv()] = flag
	}

	if e.data[option.Config.ReservedEnvs.GetAppFamilyEnv()] == "" {
		e.data[option.Config.ReservedEnvs.GetAppFamilyEnv()] = e.data[option.Config.ReservedEnvs.GetAppNameEnv()]
	}
}

func (e *env) InitAppTag(repo string) {
	if repo == "" {
		e.data[option.Config.ReservedEnvs.GetAppTagEnv()]=e.data[option.Config.ReservedEnvs.GetAppNameEnv()]+":"+e.data[option.Config.ReservedEnvs.GetAppVersionEnv()]
		return
	}
	e.data[option.Config.ReservedEnvs.GetAppTagEnv()]=repo+"/"+e.data[option.Config.ReservedEnvs.GetAppNameEnv()]+":"+e.data[option.Config.ReservedEnvs.GetAppVersionEnv()]
}

func (e *env) InitAppServers(def string) {
	storage := servers.NewLabel(e.data[option.Config.ReservedEnvs.GetAppServersEnv()])

	storage.AddRepo(def)

	err := storage.ValidateRepos()
	if err != nil {
		logger.Fatal("Failed validating servers label '%s' with error: %s", storage.String(), err)
	}

	e.data[option.Config.ReservedEnvs.GetAppServersEnv()] = storage.String()
}

func (e *env) Envs() map[string]string {
	return e.data
}

func (e *env) Labels() map[string]string {
	labels := make(map[string]string)
	labels[option.Config.ReservedEnvs.GetAppFamilyEnv()] = e.data[option.Config.ReservedEnvs.GetAppFamilyEnv()]
	labels[option.Config.ReservedEnvs.GetAppMultiversionEnv()] = e.data[option.Config.ReservedEnvs.GetAppMultiversionEnv()]
	labels[option.Config.ReservedEnvs.GetAppServersEnv()] = e.data[option.Config.ReservedEnvs.GetAppServersEnv()]

	return labels
}

func (e *env) Tag() string {
	val, _ := e.data[option.Config.ReservedEnvs.GetAppTagEnv()]
	return val
}