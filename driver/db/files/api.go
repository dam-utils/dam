package files_db

import (
	"dam/driver/storage"
)

type provider struct {
	//GetRepos() *[]storage.Repo
	//GetRepoById(id int) *storage.Repo
	//GetDefaultRepo() *storage.Repo
	//NewRepo(repo *storage.Repo)
	//ModifyRepo(repo *storage.Repo)
	//RemoveRepoById(id int)
	//GetRepoIdByName(name *string) int
	//ClearRepos()
}

func NewProvider() *provider {
	return &provider{}
}

func (p *provider) GetRepos() *[]storage.Repo {
	return GetRepos()
}

func (p *provider) GetRepoById(id int) *storage.Repo {
	return GetRepoById(id)

}

func (p *provider) GetDefaultRepo() *storage.Repo {
	return GetDefaultRepo()
}

func (p *provider) NewRepo(repo *storage.Repo) {
	NewRepo(repo)
}

func (p *provider) ModifyRepo(repo *storage.Repo) {
	ModifyRepo(repo)
}

func (p *provider) RemoveRepoById(id int) {
	RemoveRepoById(id)
}

func (p *provider) GetRepoIdByName(name *string) int {
	return GetRepoIdByName(name)
}

func (p *provider) ClearRepos() {
	ClearRepos()
}