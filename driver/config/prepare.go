package config

import (
	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
)

func Prepare() {
	switch config.DB_TYPE {
	case "files":
		fs.Touch(config.FILES_DB_REPOS)
		fs.Touch(config.FILES_DB_APPS)
	default:
		logger.Fatal("Cannot supported db '%s'", config.DB_TYPE)
	}
}
