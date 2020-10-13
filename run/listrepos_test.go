package run_test

import (
	"bytes"
	"dam/driver/structures"
	"log"
	"os"
	"testing"

	"dam/driver/db"
	"dam/run"
)

func TestListReposWithEmptyDB(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	run.ListReposFlags.Raw = true
	var buf bytes.Buffer

	stdout := `1|*|official|https://registry-1.docker.io/|
`
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestListReposWithOneLine(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	run.ListReposFlags.Raw = true
	var buf bytes.Buffer

	stdout := `1|*|official|https://registry-1.docker.io/|
2||test1|test.com|
`
	db.RDriver.NewRepo(&structures.Repo{Id: 2, Name: "test1", Server: "test.com"})
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestListReposWithDefaultLastLine(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	run.ListReposFlags.Raw = true
	var buf bytes.Buffer

	stdout := `1||official|https://registry-1.docker.io/|
2||test1|test.com|
3||test2|test2.com|user
4|*|test2|test2.com|user
`
	db.RDriver.NewRepo(&structures.Repo{Id: 2, Name: "test1", Server: "test.com"})
	db.RDriver.NewRepo(&structures.Repo{Id: 5, Name: "test2", Server: "test2.com", Username: "user", Password: "user"})
	db.RDriver.NewRepo(&structures.Repo{Default: true, Name: "test2", Server: "test2.com", Username: "user", Password: "user"})
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}