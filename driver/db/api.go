package db

import (
	"dam/driver/conf/option"
	"dam/driver/db/files/apps"
	"dam/driver/db/files/repos"
	"dam/driver/logger"
)

var (
	RDriver RProvider
	ADriver AProvider
)

func Init() {
	switch option.Config.DB.GetType() {
	case "files":
		RDriver = repos.NewProvider()
		ADriver = apps.NewProvider()
	default:
		logger.Fatal("Config option DB_TYPE='%s' not valid. DB type is bad", option.Config.DB.GetType())
	}
}
