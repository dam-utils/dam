package db

import (
	"dam/config"
	"dam/driver/db/files/apps"
	"dam/driver/db/files/repos"
	"dam/driver/logger"
)

var (
	RDriver RProvider
	ADriver AProvider
)

func Init() {
	switch config.DB_TYPE {
	case "files":
		RDriver = repos.NewProvider()
		ADriver = apps.NewProvider()
	default:
		logger.Fatal("Config option DB_TYPE='%s' not valid. DB type is bad", config.DB_TYPE)
	}
}
