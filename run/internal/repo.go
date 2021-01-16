package internal

import (
	"dam/config"
	"dam/driver/db"
	"dam/driver/logger"
	"dam/driver/structures"
	"dam/driver/validate"
	"strconv"
	"strings"
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
	repoPrefix := config.NEW_REPO_PREFIX

	for i:=0; i <= config.NEW_REPO_POSTFIX_LIMIT; i++ {
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

func ValidateReposLabelString(value string) {
	str  := strings.Split(value, config.LABEL_REPOS_SEPARATOR)
	for _, s := range str {
		if s != "" {
			err := validate.CheckServer(s)
			if err != nil {
				logger.Fatal("Not validate registry server '%s' in string '%s' with error: %s", s, value, err)
			}
		}
	}
}