package option

import "dam/config"

type Sort struct {}

func (o *Sort) GetAppType() string {
	return config.SORT_APP_TYPE
}

func (o *Sort) GetVersionType() string {
	return config.SORT_VERSION_TYPE
}