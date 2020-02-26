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
package run_test

import (
	"bytes"
	"dam/driver/db"
	"dam/driver/storage"
	"dam/run"
	"log"
	"os"
	"testing"
)

func TestRemoveRepoWithOneLineByName(t *testing.T) {
	setDefaultConfig()
	db.ClearRepos()
	run.ListReposFlags.Raw = true
	run.RemoveRepoFlags.Force = true
	var buf bytes.Buffer

	stdout := `1|*|official|https://registry-1.docker.io/|
`
	repoName := "test1"
	db.NewRepo(&storage.Repo{2, false, repoName, "localhost:5000", "", ""})
	run.RemoveRepo(repoName)
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestRemoveRepoWithOneLineById(t *testing.T) {
	setDefaultConfig()
	db.ClearRepos()
	run.ListReposFlags.Raw = true
	run.RemoveRepoFlags.Force = true
	var buf bytes.Buffer

	stdout := `1|*|official|https://registry-1.docker.io/|
`
	repoId := "2"
	db.NewRepo(&storage.Repo{0, false, "repoName", "localhost:5000", "", ""})
	run.RemoveRepo(repoId)
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestRemoveRepoWithDefaultFlag(t *testing.T) {
	setDefaultConfig()
	db.ClearRepos()
	run.ListReposFlags.Raw = true
	run.RemoveRepoFlags.Force = true
	var buf bytes.Buffer

	stdout := `1|*|official|https://registry-1.docker.io/|
2||repoName|localhost:5000|
`
	repoId := "3"
	db.NewRepo(&storage.Repo{1, false, "repoName", "localhost:5000", "", ""})
	db.NewRepo(&storage.Repo{2, true, "repoName2", "localhost:5000", "test", ""})
	run.RemoveRepo(repoId)
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestRemoveRepoDefaultRepo(t *testing.T) {
	setDefaultConfig()
	db.ClearRepos()
	run.ListReposFlags.Raw = true
	run.RemoveRepoFlags.Force = true
	var buf bytes.Buffer

	stdout := `1|*|official|https://registry-1.docker.io/|
3||repoName2|localhost:5000|test
`
	repoId := "2"
	db.NewRepo(&storage.Repo{1, true, "repoName", "localhost:5000", "", ""})
	db.NewRepo(&storage.Repo{2, false, "repoName2", "localhost:5000", "test", ""})
	run.RemoveRepo(repoId)
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}