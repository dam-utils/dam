package internal

import (
	"dam/driver/conf/option"
	"strconv"

	"dam/driver/db"
	"dam/driver/logger"
	"dam/driver/structures"
)

func GetPrefixRepo() string {
	var prefixRepo string
	repo := db.RDriver.GetDefaultRepo()
	if repo.Id != structures.OfficialRepo.Id {
		prefixRepo = repo.Server + "/"
	}
	return prefixRepo
}

func PrepareRepo(newRepo string) int {
	for _, r := range db.RDriver.GetRepos() {
		if r.Server == newRepo {
			return r.Id
		}
	}

	newRepoName := generateNewRepoName()
	logger.Info("Add new repository '%s'", newRepoName)

	return db.RDriver.NewRepo(&structures.Repo{
		Name: newRepoName,
		Server: newRepo,
		Username: "",
		Password: "",
	})
}

func generateNewRepoName() string {
	repoPrefix := option.Config.DefaultRepo.GetNewRepoPrefix()

	for i:=0; i <= option.Config.DefaultRepo.GetNewRepoPostfixLimit(); i++ {
		for _, r := range db.RDriver.GetRepos() {
			newName := repoPrefix+strconv.Itoa(i)
			if r.Name == newName {
				continue
			}
			return newName
		}

	}
	logger.Fatal("Cannot create auto repo. Limit generated name is full. See util config.")
	return ""
}
