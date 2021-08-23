package option

import "dam/config"

type FilesDB struct {}

func (o *FilesDB) GetSeparator() string {
	return config.FILES_DB_SEPARATOR
}

func (o *FilesDB) GetBoolFlagSymbol() string {
	return config.FILES_DB_BOOL_FLAG_SYMBOL
}

func (o *FilesDB) GetReposFilename() string {
	return config.FILES_DB_REPOS_FILENAME
}

func (o *FilesDB) GetAppsFilename() string {
	return config.FILES_DB_APPS_FILENAME
}

func (o *FilesDB) GetTmp() string {
	return config.FILES_DB_TMP
}

