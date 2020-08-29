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
	"os"

	"dam/config"
	"dam/driver/db/files/repos/internal"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/structures"
)

func (p *provider) NewRepo(repo *structures.Repo) {
	repos := p.GetRepos()
	//preparedRepo := preparePassword(repo)
	repo.Id = internal.GetNewRepoID(repos)
	if len(repos) == 0 {
		repo.Default = true
		repo.Id = 2
	}
	var preparedRepos []*structures.Repo
	if repo.Default {
		preparedRepos = internal.CleanDefaults(repos)
	} else {
		preparedRepos = repos
	}
	newRepos := append(preparedRepos, repo)
	internal.SaveRepos(newRepos)
}

func (p *provider) ModifyRepo(mRepo *structures.Repo) {
	cleanRepos := internal.CleanReposDefault(p.GetRepos())
	defRepo := p.GetDefaultRepo()

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

	var newRepos []*structures.Repo
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

	prepNewRepos := internal.PrepareClearRepos(newRepos)
	logger.Debug("ModifyRepo():prepareClearRepos(&newRepos): '%v'", prepNewRepos)
	internal.SaveRepos(prepNewRepos)
}

func (p *provider) ClearRepos() {
	f, err := os.OpenFile(config.FILES_DB_TMP, os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open repo file '%s' with error: %s", config.FILES_DB_TMP, err)
	}

	_, err = f.WriteString("")
	if err != nil {
		logger.Fatal("Cannot write to repo file '%s' with error: %s", config.FILES_DB_TMP, err)
	}

	err = f.Sync()
	if err != nil {
		logger.Fatal("Cannot sync repo file '%s' with error: %s", config.FILES_DB_TMP, err)
	}
	err = f.Close()
	if err != nil {
		logger.Fatal("Cannot close from repo file '%s' with error: %s", config.FILES_DB_TMP, err)
	}

	fs.MoveFile(config.FILES_DB_TMP, config.FILES_DB_REPOS)
}

func (p *provider) GetRepos() []*structures.Repo {
	// Ex: 2||auto_repo|packages.test.com|admin|YWRtaW4K|
	var repos []*structures.Repo
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
		repos = append(repos, internal.Str2Repo(NewLine))
	}

	offRepo := structures.OfficialRepo
	if len(repos) == 0 {
		repos = append(repos, &offRepo)
		internal.SaveRepos(repos)
	}

	internal.ValidatingReposDB(repos)

	return repos
}

func (p *provider) GetRepoById(id int) *structures.Repo {
	for _, repo := range p.GetRepos() {
		if repo.Id == id {
			return repo
		}
	}
	return nil
}

func (p *provider) GetDefaultRepo() *structures.Repo {
	for _, repo := range p.GetRepos() {
		if repo.Default {
			return repo
		}
	}
	return nil
}

func (p *provider) RemoveRepoById(id int) {
	var newRepos []*structures.Repo

	if id == 1 {
		logger.Fatal("Cannot remove official Repository. This is base repository in DB")
	}

	repos := p.GetRepos()
	for _, repo := range repos {
		if repo.Id != id {
			newRepos = append(newRepos, repo)
		}
	}
	if len(newRepos) < len(repos) {
		preparedRepos := internal.PrepareDefaultInRepos(newRepos)
		internal.SaveRepos(preparedRepos)
	} else {
		logger.Fatal("Not found Id or Name of Repository")
	}
}

func (p *provider) GetRepoIdByName(name *string) int {
	for _, repo := range p.GetRepos() {
		if repo.Name == *name {
			return repo.Id
		}
	}

	logger.Fatal("Internal error. Not found repo ID for name '%s'", *name)
	return -1
}