package option

import "dam/config"

type ReservedEnvs struct {}

func (o *ReservedEnvs) GetOSEnvPrefix() string {
	return config.OS_ENV_PREFIX
}

func (o *ReservedEnvs) GetDefaultAppName() string {
	return config.DEF_APP_NAME
}

func (o *ReservedEnvs) GetDefaultAppVersion() string {
	return config.DEF_APP_VERS
}

func (o *ReservedEnvs) GetAppNameEnv() string {
	return config.APP_NAME_ENV
}

func (o *ReservedEnvs) GetAppVersionEnv() string {
	return config.APP_VERS_ENV
}

func (o *ReservedEnvs) GetAppFamilyEnv() string {
	return config.APP_FAMILY_ENV
}

func (o *ReservedEnvs) GetAppMultiversionEnv() string {
	return config.APP_MULTIVERSION_ENV
}

func (o *ReservedEnvs) GetAppTagEnv() string {
	return config.APP_TAG_ENV
}

func (o *ReservedEnvs) GetAppServersEnv() string {
	return config.APP_SERVERS_ENV
}