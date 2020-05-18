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