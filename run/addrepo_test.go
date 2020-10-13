package run_test

import (
	"bytes"
	"log"
	"os"
	"testing"

	"dam/driver/db"
	"dam/run"
)

func TestAddRepoOneRepo(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	run.ListReposFlags.Raw = true
	run.AddRepoFlags.Default = true
	run.AddRepoFlags.Name = "test"
	run.AddRepoFlags.Server = "localhost:5000"
	var buf bytes.Buffer

	stdout := `1||official|https://registry-1.docker.io/|
2|*|test|localhost:5000|
`
	run.AddRepo()
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestAddRepoWithNotDefaultFlag(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	run.ListReposFlags.Raw = true
	run.AddRepoFlags.Default = false
	run.AddRepoFlags.Name = "test"
	run.AddRepoFlags.Server = "localhost:5000"
	var buf bytes.Buffer

	stdout := `1|*|official|https://registry-1.docker.io/|
2||test|localhost:5000|
`
	run.AddRepo()
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestAddRepoTwoRepo(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	run.ListReposFlags.Raw = true

	run.AddRepoFlags.Default = true
	run.AddRepoFlags.Name = "test"
	run.AddRepoFlags.Server = "localhost:5000"
	run.AddRepoFlags.Username="test"
	run.AddRepo()

	run.AddRepoFlags.Default = true
	run.AddRepoFlags.Name = "test2"
	run.AddRepoFlags.Server = "localhost:5000"
	run.AddRepoFlags.Username="test2"
	run.AddRepo()

	var buf bytes.Buffer
	stdout := `1||official|https://registry-1.docker.io/|
2||test|localhost:5000|test
3|*|test2|localhost:5000|test2
`
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}