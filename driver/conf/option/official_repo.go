package option

import "dam/config"

type OfficialRepo struct {}

func (o *OfficialRepo) GetAuthURL() string {
	return config.OFFICIAL_REGISTRY_AUTH_URL
}

func (o *OfficialRepo) GetURL() string {
	return config.OFFICIAL_REGISTRY_URL
}

func (o *OfficialRepo) GetName() string {
	return config.OFFICIAL_REGISTRY_NAME
}