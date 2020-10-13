package internal

import (
	"encoding/base64"
	"os"
	"strconv"
	"strings"

	"dam/config"
	fs "dam/driver/filesystem"
	"dam/driver/logger"
	"dam/driver/structures"
	"dam/driver/validate"
)

func CleanDefaults(repos []*structures.Repo) []*structures.Repo {
	var newRepo []*structures.Repo
	for _, repo := range repos {
		repo.Default = false
		newRepo = append(newRepo, repo)
	}
	return newRepo
}

func PrepareClearRepos(repos []*structures.Repo) []*structures.Repo {
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

func CleanReposDefault(repos []*structures.Repo) []*structures.Repo {
	var newRepos []*structures.Repo

	for _, repo := range repos {
		repo.Default = false
		newRepos = append(newRepos, repo)
	}
	return newRepos
}

func repo2str(repo *structures.Repo) *string {
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

func SaveRepos(repos []*structures.Repo) {
	newRepos := preparePasswordRepos(repos)

	f, err := os.OpenFile(config.FILES_DB_TMP, os.O_WRONLY|os.O_CREATE, 0644)
	defer func() {
		if f != nil {
			f.Close()
		}
	}()
	if err != nil {
		logger.Fatal("Cannot open repo file '%s' with error: %s", config.FILES_DB_TMP, err)
	}

	for _, repo := range newRepos {
		newLine := repo2str(repo)
		_, err := f.WriteString(*newLine)
		if err != nil {
			logger.Fatal("Cannot write to repo file '%s' with error: %s", config.FILES_DB_TMP, err)
		}
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

func PrepareDefaultInRepos(repos []*structures.Repo) []*structures.Repo {
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

func GetNewRepoID(repos []*structures.Repo) int {
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

func Str2Repo(str string) *structures.Repo {
	repoArray := new(structures.Repo)
	strRepo := strings.Split(str, config.FILES_DB_SEPARATOR)

	pass, err := base64ToStr(strRepo[5])
	if err != nil {
		logger.Fatal("Internal error. Cannot read the password of user '%s' in line '%s'", repoArray.Username, str)
	}

	if validate.CheckRepoID(strRepo[0]) != nil {
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

func preparePasswordRepos(repos []*structures.Repo) []*structures.Repo {
	var newRepos []*structures.Repo

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

func ValidatingReposDB(repos []*structures.Repo) {
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
