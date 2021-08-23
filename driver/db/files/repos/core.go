package repos

import (
	"bufio"
	"os"

	"dam/driver/conf/option"
	"dam/driver/db/files/repos/internal"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/structures"
)

func (p *provider) NewRepo(repo *structures.Repo) int {
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

	return repo.Id
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
	f, err := os.OpenFile(option.Config.FilesDB.GetTmp(), os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open repo file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
	}

	_, err = f.WriteString("")
	if err != nil {
		logger.Fatal("Cannot write to repo file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
	}

	err = f.Sync()
	if err != nil {
		logger.Fatal("Cannot sync repo file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
	}
	err = f.Close()
	if err != nil {
		logger.Fatal("Cannot close from repo file '%s' with error: %s", option.Config.FilesDB.GetTmp(), err)
	}

	fs.MoveFile(option.Config.FilesDB.GetTmp(), option.Config.FilesDB.GetReposFilename())
}

func (p *provider) GetRepos() []*structures.Repo {
	// Ex: 2||auto_repo|packages.test.com|admin|YWRtaW4K|
	var repos []*structures.Repo

	p.connect()
	fileScanner := bufio.NewScanner(p.client)
	defer p.close()

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
