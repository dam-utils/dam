package conf

import (
	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
)

func Prepare() {
	switch config.DB_TYPE {
	case "files":
		reposDir := fs.GetDir(config.FILES_DB_REPOS_FILENAME)
		if !fs.IsExistDir(reposDir) {
			fs.MkDir(fs.GetDir(config.FILES_DB_REPOS_FILENAME))
		}
		fs.Touch(config.FILES_DB_REPOS_FILENAME)

		appDir := fs.GetDir(config.FILES_DB_APPS_FILENAME)
		if !fs.IsExistDir(appDir){
			fs.MkDir(fs.GetDir(config.FILES_DB_APPS_FILENAME))
		}
		fs.Touch(config.FILES_DB_APPS_FILENAME)
	default:
		logger.Fatal("Cannot supported db '%s'", config.DB_TYPE)
	}
}
