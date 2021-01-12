package envs

import "dam/config"

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
		e.data[config.APP_NAME_ENV] = flag
	}

	if e.data[config.APP_NAME_ENV] == "" {
		e.data[config.APP_NAME_ENV] = def
	}
}

func (e *env) InitAppVersion(def, flag string) {
	if flag != "" {
		e.data[config.APP_VERS_ENV] = flag
	}

	if e.data[config.APP_VERS_ENV] == "" {
		e.data[config.APP_VERS_ENV] = def
	}
}

func (e *env) InitAppMultiversion(flag string) {
	if flag != "" {
		e.data[config.APP_MULTIVERSION_ENV] = flag
	}

	if e.data[config.APP_MULTIVERSION_ENV] == "" {
		e.data[config.APP_MULTIVERSION_ENV] = "false"
	}
}

func (e *env) InitAppFamily(flag string) {
	if flag != "" {
		e.data[config.APP_FAMILY_ENV] = flag
	}

	if e.data[config.APP_FAMILY_ENV] == "" {
		e.data[config.APP_FAMILY_ENV] = e.data[config.APP_NAME_ENV]
	}
}

func (e *env) InitAppTag(repo string) {
	if repo == "" {
		e.data[config.APP_TAG_ENV]=e.data[config.APP_NAME_ENV]+":"+e.data[config.APP_VERS_ENV]
		return
	}
	e.data[config.APP_TAG_ENV]=repo+"/"+e.data[config.APP_NAME_ENV]+":"+e.data[config.APP_VERS_ENV]
}

func (e *env) InitAppServers(def string) {
	if e.data[config.APP_SERVERS] == "" {
		e.data[config.APP_SERVERS] = def
	}
}

func (e *env) Envs() map[string]string {
	return e.data
}

func (e *env) Labels() map[string]string {
	labels := make(map[string]string)
	labels[config.APP_FAMILY_ENV] = e.data[config.APP_FAMILY_ENV]
	labels[config.APP_MULTIVERSION_ENV] = e.data[config.APP_MULTIVERSION_ENV]

	return labels
}

func (e *env) Tag() string {
	val, _ := e.data[config.APP_TAG_ENV]
	return val
}