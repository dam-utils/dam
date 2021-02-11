package db

import "dam/driver/structures"

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