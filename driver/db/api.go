package db

import (
	"dam/driver/conf/option"
	"dam/driver/db/files/apps"
	"dam/driver/db/files/repos"
	"dam/driver/logger"
	"dam/driver/structures"
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

type RProvider interface {
	GetRepos() []*structures.Repo
	GetRepoById(id int) *structures.Repo
	GetDefaultRepo() *structures.Repo
	NewRepo(repo *structures.Repo) int
	ModifyRepo(repo *structures.Repo)
	RemoveRepoById(id int)
	GetRepoIdByName(name *string) int
	ClearRepos()
}

type AProvider interface {
	GetApps() []*structures.App
	NewApp(app *structures.App)
	GetAppById(id int) *structures.App
	ExistFamily(family string) bool
	RemoveAppById(id int)
}
