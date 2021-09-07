package run_test

import (
	"bytes"
	"log"
	"os"
	"testing"

	"dam/driver/db"
	"dam/driver/structures"
	"dam/run"
)

func TestRemoveRepoWithOneLineByName(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	db.ADriver.ClearApps()
	run.ListReposFlags.Raw = true
	run.RemoveRepoFlags.Force = true
	var buf bytes.Buffer

	stdout := `1|*|official|https://registry-1.docker.io/|
`
	repoName := "test1"
	db.RDriver.NewRepo(&structures.Repo{Id: 2, Name: repoName, Server: "localhost:5000"})
	run.RemoveRepo(repoName)
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestRemoveRepoWithOneLineById(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	db.ADriver.ClearApps()
	run.ListReposFlags.Raw = true
	run.RemoveRepoFlags.Force = true
	var buf bytes.Buffer

	stdout := `1|*|official|https://registry-1.docker.io/|
`
	repoId := "2"
	db.RDriver.NewRepo(&structures.Repo{Name: "repoName", Server: "localhost:5000"})
	run.RemoveRepo(repoId)
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestRemoveRepoWithDefaultFlag(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	db.ADriver.ClearApps()
	run.ListReposFlags.Raw = true
	run.RemoveRepoFlags.Force = true
	var buf bytes.Buffer

	stdout := `1|*|official|https://registry-1.docker.io/|
2||repoName|localhost:5000|
`
	repoId := "3"
	db.RDriver.NewRepo(&structures.Repo{Id: 1, Name: "repoName", Server: "localhost:5000"})
	db.RDriver.NewRepo(&structures.Repo{Id: 2, Default: true, Name: "repoName2", Server: "localhost:5000", Username: "test"})
	run.RemoveRepo(repoId)
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestRemoveRepoDefaultRepo(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	db.ADriver.ClearApps()
	run.ListReposFlags.Raw = true
	run.RemoveRepoFlags.Force = true
	var buf bytes.Buffer

	stdout := `1|*|official|https://registry-1.docker.io/|
3||repoName2|localhost:5000|test
`
	repoId := "2"
	db.RDriver.NewRepo(&structures.Repo{Id: 1, Default: true, Name: "repoName", Server: "localhost:5000"})
	db.RDriver.NewRepo(&structures.Repo{Id: 2, Name: "repoName2", Server: "localhost:5000", Username: "test"})
	run.RemoveRepo(repoId)
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestRemoveRepoWithApps(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	db.ADriver.ClearApps()
	run.ListReposFlags.Raw = true
	run.RemoveRepoFlags.Force = true
	run.ListFlags.Raw = true
	var repoBuf, appBuf bytes.Buffer

	stdoutRepo := `1|*|official|https://registry-1.docker.io/|
3||repoName2|localhost:5000|test
`
	stdoutApp := `0|000000000000|test|1.0.0|0||test
`

	repoId := "2"
	db.RDriver.NewRepo(&structures.Repo{Id: 1, Default: true, Name: "repoName", Server: "localhost:5000"})
	db.RDriver.NewRepo(&structures.Repo{Id: 2, Name: "repoName2", Server: "localhost:5000", Username: "test"})
	db.ADriver.NewApp(&structures.App{
		Id: 1,
		DockerID: "000000000000",
		ImageName: "test",
		ImageVersion: "1.0.0",
		RepoID: 2,
		Family: "test",
	})
	run.RemoveRepo(repoId)
	log.SetOutput(&repoBuf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdoutRepo, repoBuf.String())

	log.SetOutput(&appBuf)
	run.List()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdoutApp, appBuf.String())

	dropTestDB(t)
}