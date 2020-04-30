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
package filesdb

import (
	"bufio"
	"dam/config"
	d_log "dam/driver/logger"
	"dam/driver/storage"
	"dam/exam"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"strings"
)


func NewRepo(repo *storage.Repo) {
	repos := GetRepos()
	//preparedRepo := preparePassword(repo)
	repo.Id = getNewRepoID(repos)
	if len(*repos) == 0 {
		repo.Default = true
		repo.Id = 2
	}
	var preparedRepos []storage.Repo
	if repo.Default {
		preparedRepos = cleanDefaults(repos)
	} else {
		preparedRepos = *repos
	}
	newRepos := append(preparedRepos, *repo)
	saveRepos(&newRepos)
}

func cleanDefaults(repos *[]storage.Repo) []storage.Repo {
	var newRepo []storage.Repo
	for _, repo := range *repos {
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

	var newRepos []storage.Repo
	for _, repo := range *cleanRepos {
		if repo.Id == defRepo.Id {
			repo.Default = defRepo.Default
		}
		if repo.Id == mRepo.Id {
			newRepos = append(newRepos, *mRepo)
		} else {
			newRepos = append(newRepos, repo)
		}
	}

	prepNewRepos := prepareClearRepos(&newRepos)
	d_log.Debug(fmt.Sprintf("ModifyRepo():prepareClearRepos(&newRepos): '%v'", prepNewRepos))
	saveRepos(prepNewRepos)
}

func prepareClearRepos(repos *[]storage.Repo) *[]storage.Repo {
	existingDefault := false
	for _, repo := range *repos {
		if repo.Default {
			existingDefault = true
		}
	}
	if !existingDefault {
		newRepos := *repos
		newRepos[0].Default = true
		return &newRepos
	}
	return repos
}

func cleanReposDefault(repos *[]storage.Repo) *[]storage.Repo {
	var newRepos []storage.Repo
	for _, repo := range *repos {
			repo.Default = false
			newRepos = append(newRepos, repo)
	}
	return &newRepos
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

func saveRepos(repos *[]storage.Repo) {
	newRepos := preparePasswordRepos(repos)

	f, err := os.OpenFile(config.FILES_DB_TMP, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		d_log.Fatal(err.Error())
	}
	defer f.Close()

	for _, repo := range *newRepos {
		newLine := repo2str(&repo)
		_, err := f.WriteString(*newLine)
		if err != nil {
			d_log.Fatal(err.Error())
		}
	}
	err = f.Sync()
	if err != nil {
		d_log.Fatal(err.Error())
	}
	err = f.Close()
	if err != nil {
		d_log.Fatal(err.Error())
	}

	err = moveFile(config.FILES_DB_TMP, config.FILES_DB_REPOS)
	if err != nil {
		d_log.Fatal(err.Error())
	}
}

func ClearRepos() {
	f, err := os.OpenFile(config.FILES_DB_TMP, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		d_log.Fatal(err.Error())
	}
	defer f.Close()
	_, err = f.WriteString("")
	if err != nil {
		d_log.Fatal(err.Error())
	}

	err = f.Sync()
	if err != nil {
		d_log.Fatal(err.Error())
	}
	err = f.Close()
	if err != nil {
		d_log.Fatal(err.Error())
	}

	err = moveFile(config.FILES_DB_TMP, config.FILES_DB_REPOS)
	if err != nil {
		d_log.Fatal(err.Error())
	}
}

func GetRepos() *[]storage.Repo {
	// Ex: 2||auto_repo|packages.test.com|admin|YWRtaW4K|
	var Repos []storage.Repo
	fileHandle, err := os.Open(config.FILES_DB_REPOS)
	if err != nil {
		d_log.Fatal("Cannot open file '"+config.FILES_DB_REPOS+"'"  )
	}
	defer func (){
		err := fileHandle.Close()
		if err != nil {
			d_log.Fatal(fmt.Sprintf("Cannot close file '%v'",config.FILES_DB_REPOS))
		}
	}()

	fileScanner := bufio.NewScanner(fileHandle)
	for fileScanner.Scan() {
		NewLine := fileScanner.Text()
		Repos = append(Repos, *str2Repo(NewLine))
	}

	if len(Repos) == 0 {
		Repos = append(Repos, storage.OfficialRepo)
		saveRepos(&Repos)
	}

	if config.FILES_DB_VALIDATE {
		internalValidatingReposDB(&Repos)
	}
	return &Repos
}

func GetRepoById(id int) *storage.Repo {
	repos := GetRepos()
	for _, repo := range *repos {
		if repo.Id == id {
			return &repo
		}
	}
	return nil
}

func GetDefaultRepo() *storage.Repo {
	repos := GetRepos()
	for _, repo := range *repos {
		if repo.Default {
			return &repo
		}
	}
	return nil
}

func RemoveRepoById(id int) {
	if id == 1 {
		d_log.Fatal("Cannot remove official Repository. This is base repository in DB")
	}

	NewRepos := new([]storage.Repo)
	repos := GetRepos()
	for _, repo := range *repos {
		if repo.Id != id {
			*NewRepos = append(*NewRepos, repo)
		}
	}
	if len(*NewRepos) < len(*repos) {
		preparedRepos := prepareDefaultInRepos(*NewRepos)
		saveRepos(preparedRepos)
	} else {
		d_log.Fatal("Not found Id or Name of Repository")
	}
}

func GetRepoIdByName(name *string) int {
	repos := GetRepos()
	for _, repo := range *repos {
		if repo.Name == *name {
			return repo.Id
		}
	}
	return 0
}

func prepareDefaultInRepos(repos []storage.Repo) *[]storage.Repo {
	def := false
	for _, repo := range repos {
		if repo.Default {
			def = true
		}
	}
	if !def {
		repos[0].Default = true
	}

	return &repos
}

func getNewRepoID(repos *[]storage.Repo) int {
	Res := 0

	if len(*repos) == 0 {
		return 0
	}
	for _, repo := range *repos {
		if repo.Id >= Res {
			Res = repo.Id
		}
	}
	return Res +1
}

func str2Repo(repo string) *storage.Repo {
	// Ex: 2||auto_repo|localhost:5000|admin|YWRtaW4K
	Repo := new(storage.Repo)
	ParseRepo := strings.Split(repo, config.FILES_DB_SEPARATOR)
	Repo.Id, _ = strconv.Atoi(ParseRepo[0])
	if ParseRepo[1] == config.FILES_DB_BOOL_FLAG {
		Repo.Default = true
	} else {
		Repo.Default = false
	}
	Repo.Name = ParseRepo[2]
	Repo.Server = ParseRepo[3]
	Repo.Username = ParseRepo[4]
	var err error
	Repo.Password, err = base64ToStr(ParseRepo[5])
	if err != nil {
		d_log.Fatal("Cannot read the password of user '"+Repo.Username+"'")
	}
	return Repo
}

func preparePasswordRepos(repos *[]storage.Repo) *[]storage.Repo {
	var NewRepos []storage.Repo

	for _, repo := range *repos {
		repo.Password = strToBase64(repo.Password)
		NewRepos = append(NewRepos, repo)
	}
	return &NewRepos
}

func strToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

func base64ToStr(str string) (string, error) {
	sDec, err := base64.StdEncoding.DecodeString(str)
	return string(sDec), err
}

func internalValidatingReposDB(repos *[]storage.Repo) {
	defRepo := false
	for _, repo := range *repos{
		if repo.Default {
			if defRepo {
				d_log.Fatal("Internal error. Found many default repositories in DB. Default repository must be only one")
			} else {
				defRepo = true
			}
		}
		if exam.CheckRepoName(repo.Name) != nil {
			d_log.Fatal("Internal error. Repo name '"+repo.Name+"' is invalid in DB")
		}
		if exam.CheckLogin(repo.Username) != nil {
			d_log.Fatal("Internal error. Repo login '"+repo.Username+"' is invalid in DB")
		}
		if exam.CheckServer(repo.Server) != nil {
			d_log.Fatal("Internal error. Repo server '"+repo.Server+"' is invalid in DB")
		}
	}
	if !defRepo {
		d_log.Fatal("Internal error. Not found default repository in DB")
	}
}

