package option

import "dam/config"

type Export struct {}

func (o *Export) GetAppSeparator() string {
	return config.EXPORT_APP_SEPARATOR
}

func (o *Export) GetAppsFileName() string {
	return config.EXPORT_APPS_FILE_NAME
}