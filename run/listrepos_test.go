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
	db.RDriver.NewRepo(&storage.Repo{Id: 2, Name: "test1", Server: "test.com"})
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
	db.RDriver.NewRepo(&storage.Repo{Id: 2, Name: "test1", Server: "test.com"})
	db.RDriver.NewRepo(&storage.Repo{Id: 5, Name: "test2", Server: "test2.com", Username: "user", Password: "user"})
	db.RDriver.NewRepo(&storage.Repo{Default: true, Name: "test2", Server: "test2.com", Username: "user", Password: "user"})
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}