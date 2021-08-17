package config

import (
	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
)

func Prepare() {
	switch config.DB_TYPE {
	case "files":
		reposDir := fs.GetDir(config.FILES_DB_REPOS)
		if !fs.IsExistDir(reposDir) {
			fs.MkDir(fs.GetDir(config.FILES_DB_REPOS))
		}
		fs.Touch(config.FILES_DB_REPOS)

		appDir := fs.GetDir(config.FILES_DB_APPS)
		if !fs.IsExistDir(appDir){
			fs.MkDir(fs.GetDir(config.FILES_DB_APPS))
		}
		fs.Touch(config.FILES_DB_APPS)
	default:
		logger.Fatal("Cannot supported db '%s'", config.DB_TYPE)
	}
}
