package repos

import "os"

type provider struct {
	//GetRepos() []*storage.Repo
	//GetRepoById(id int) *storage.Repo
	//GetDefaultRepo() *storage.Repo
	//NewRepo(repo *storage.Repo) int
	//ModifyRepo(repo *storage.Repo)
	//RemoveRepoById(id int)
	//GetRepoIdByName(name *string) int
	//ClearRepos()

	client *os.File
}

func NewProvider() *provider {
	return &provider{}
}