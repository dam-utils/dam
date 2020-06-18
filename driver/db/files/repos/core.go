// Copyright 2020 The Docker Applications Manager Authors
//
//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
package repos

import (
	"bufio"
	"encoding/base64"
	"os"
	"strconv"
	"strings"

	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/storage"
	"dam/driver/validate"
)


func NewRepo(repo *storage.Repo) {
	repos := GetRepos()
	//preparedRepo := preparePassword(repo)
	repo.Id = getNewRepoID(repos)
	if len(repos) == 0 {
		repo.Default = true
		repo.Id = 2
	}
	var preparedRepos []*storage.Repo
	if repo.Default {
		preparedRepos = cleanDefaults(repos)
	} else {
		preparedRepos = repos
	}
	newRepos := append(preparedRepos, repo)
	saveRepos(newRepos)
}

func cleanDefaults(repos []*storage.Repo) []*storage.Repo {
	var newRepo []*storage.Repo
	for _, repo := range repos {
		repo.Default = false
		newRepo = append(newRepo, repo)
	}
	return newRepo
}

func ModifyRepo(mRepo *storage.Repo) {
	cleanRepos := cleanReposDefault(GetRepos())
	defRepo := GetDefaultRepo()

	if mRepo.Default {
		if mRepo.Id == defRepo.Id {
			mRepo.Default = true
			defRepo.Default = true
		} else {
			mRepo.Default = true
			defRepo.Default = false
		}
	} else {
		if mRepo.Id == defRepo.Id {
			mRepo.Default = false
			defRepo.Default = false
		} else {
			mRepo.Default = false
			defRepo.Default = true
		}
	}

	var newRepos []*storage.Repo
	for _, repo := range cleanRepos {
		if repo.Id == defRepo.Id {
			repo.Default = defRepo.Default
		}
		if repo.Id == mRepo.Id {
			newRepos = append(newRepos, mRepo)
		} else {
			newRepos = append(newRepos, repo)
		}
	}

	prepNewRepos := prepareClearRepos(newRepos)
	logger.Debug("ModifyRepo():prepareClearRepos(&newRepos): '%v'", prepNewRepos)
	saveRepos(prepNewRepos)
}

func prepareClearRepos(repos []*storage.Repo) []*storage.Repo {
	existingDefault := false
	for _, repo := range repos {
		if repo.Default {
			existingDefault = true
		}
	}
	if !existingDefault {
		newRepos := repos
		newRepos[0].Default = true
		return newRepos
	}
	return repos
}

func cleanReposDefault(repos []*storage.Repo) []*storage.Repo {
	var newRepos []*storage.Repo

	for _, repo := range repos {
			repo.Default = false
			newRepos = append(newRepos, repo)
	}
	return newRepos
}

func repo2str(repo *storage.Repo) *string {
	var def string
	if repo.Default {
		def = config.FILES_DB_BOOL_FLAG
	} else {
		def = ""
	}

	var repoStr string
	sep := config.FILES_DB_SEPARATOR
	fields := []string{strconv.Itoa(repo.Id), def, repo.Name, repo.Server, repo.Username, repo.Password}
	lenF := len(fields)
	for i, field := range fields {
		if i == lenF - 1 {
			repoStr = repoStr + field + "\n"
		} else {
			repoStr = repoStr + field + sep
		}
	}
	return &repoStr
}

func saveRepos(repos []*storage.Repo) {
	newRepos := preparePasswordRepos(repos)

	f, err := os.OpenFile(config.FILES_DB_TMP, os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open repo file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
	}

	for _, repo := range newRepos {
		newLine := repo2str(repo)
		_, err := f.WriteString(*newLine)
		if err != nil {
			logger.Fatal("Cannot write to repo file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
		}
	}
	err = f.Sync()
	if err != nil {
		logger.Fatal("Cannot sync repo file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
	}
	err = f.Close()
	if err != nil {
		logger.Fatal("Cannot close from repo file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
	}

	fs.MoveFile(config.FILES_DB_TMP, config.FILES_DB_REPOS)
}

func ClearRepos() {
	f, err := os.OpenFile(config.FILES_DB_TMP, os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open repo file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
	}

	_, err = f.WriteString("")
	if err != nil {
		logger.Fatal("Cannot write to repo file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
	}

	err = f.Sync()
	if err != nil {
		logger.Fatal("Cannot sync repo file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
	}
	err = f.Close()
	if err != nil {
		logger.Fatal("Cannot close from repo file '%s' with error: %s", config.FILES_DB_TMP, err.Error())
	}

	fs.MoveFile(config.FILES_DB_TMP, config.FILES_DB_REPOS)
}

func GetRepos() []*storage.Repo {
	// Ex: 2||auto_repo|packages.test.com|admin|YWRtaW4K|
	var repos []*storage.Repo
	f, err := os.Open(config.FILES_DB_REPOS)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open file '%s'", config.FILES_DB_REPOS)
	}


	fileScanner := bufio.NewScanner(f)
	for fileScanner.Scan() {
		NewLine := fileScanner.Text()
		repos = append(repos, str2Repo(NewLine))
	}

	offRepo := storage.OfficialRepo
	if len(repos) == 0 {
		repos = append(repos, &offRepo)
		saveRepos(repos)
	}

	internalValidatingReposDB(repos)

	return repos
}

func GetRepoById(id int) *storage.Repo {
	for _, repo := range GetRepos() {
		if repo.Id == id {
			return repo
		}
	}
	return nil
}

func GetDefaultRepo() *storage.Repo {
	for _, repo := range GetRepos() {
		if repo.Default {
			return repo
		}
	}
	return nil
}

func RemoveRepoById(id int) {
	var newRepos []*storage.Repo

	if id == 1 {
		logger.Fatal("Cannot remove official Repository. This is base repository in DB")
	}

	repos := GetRepos()
	for _, repo := range repos {
		if repo.Id != id {
			newRepos = append(newRepos, repo)
		}
	}
	if len(newRepos) < len(repos) {
		preparedRepos := prepareDefaultInRepos(newRepos)
		saveRepos(preparedRepos)
	} else {
		logger.Fatal("Not found Id or Name of Repository")
	}
}

func GetRepoIdByName(name *string) int {
	for _, repo := range GetRepos() {
		if repo.Name == *name {
			return repo.Id
		}
	}

	logger.Fatal("Internal error. Not found repo ID for name '%s'", *name)
	return -1
}

func prepareDefaultInRepos(repos []*storage.Repo) []*storage.Repo {
	def := false
	for _, repo := range repos {
		if repo.Default {
			def = true
		}
	}
	if !def {
		repos[0].Default = true
	}

	return repos
}

func getNewRepoID(repos []*storage.Repo) int {
	Res := 0

	if len(repos) == 0 {
		return 0
	}
	for _, repo := range repos {
		if repo.Id >= Res {
			Res = repo.Id
		}
	}
	return Res +1
}

func str2Repo(str string) *storage.Repo {
	repoArray := new(storage.Repo)
	strRepo := strings.Split(str, config.FILES_DB_SEPARATOR)

	pass, err := base64ToStr(strRepo[5])
	if err != nil {
		logger.Fatal("Internal error. Cannot read the password of user '%s' in line '%s'", repoArray.Username, str)
	}

	if validate.CheckID(strRepo[0]) != nil {
		logger.Fatal("Internal error. Cannot parse the repo id in line '%s'", str)
	}
	if validate.CheckBool(strRepo[1]) != nil {
		logger.Fatal("Internal error. Cannot parse the default flag in line '%s'", str)
	}
	if validate.CheckRepoName(strRepo[2]) != nil {
		logger.Fatal("Internal error. Cannot parse the repo name in line '%s'", str)
	}
	if validate.CheckServer(strRepo[3]) != nil {
		logger.Fatal("Internal error. Cannot parse the server in line '%s'", str)
	}
	if validate.CheckLogin(strRepo[4]) != nil {
		logger.Fatal("Internal error. Cannot parse the username in line '%s'", str)
	}
	if validate.CheckPassword(pass) != nil {
		logger.Fatal("Internal error. Cannot parse the password in line '%s'", str)
	}


	repoArray.Id, _ = strconv.Atoi(strRepo[0])
	if strRepo[1] == config.FILES_DB_BOOL_FLAG {
		repoArray.Default = true
	}
	repoArray.Name = strRepo[2]
	repoArray.Server = strRepo[3]
	repoArray.Username = strRepo[4]
	repoArray.Password = pass

	return repoArray
}

func preparePasswordRepos(repos []*storage.Repo) []*storage.Repo {
	var newRepos []*storage.Repo

	for _, repo := range repos {
		repo.Password = strToBase64(repo.Password)
		newRepos = append(newRepos, repo)
	}
	return newRepos
}

func strToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func base64ToStr(str string) (string, error) {
	sDec, err := base64.StdEncoding.DecodeString(str)
	return string(sDec), err
}

func internalValidatingReposDB(repos []*storage.Repo) {
	defRepo := false
	for _, repo := range repos{
		if repo.Default {
			if defRepo {
				logger.Fatal("Internal error. Found many default repositories in DB. Default repository must be only one")
			} else {
				defRepo = true
			}
		}

		if validate.CheckRepoName(repo.Name) != nil {
			logger.Fatal("Internal error. Repo name '%s' is invalid in DB", repo.Name)
		}
		if validate.CheckLogin(repo.Username) != nil {
			logger.Fatal("Internal error. Repo login '%s' is invalid in DB", repo.Username)
		}
		if validate.CheckServer(repo.Server) != nil {
			logger.Fatal("Internal error. Repo server '%s' is invalid in DB", repo.Server)
		}

	}
	if !defRepo {
		logger.Fatal("Internal error. Not found default repository in DB")
	}
}

