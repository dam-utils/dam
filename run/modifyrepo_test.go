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

func TestModifyRepoNameAndServer(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	run.ListReposFlags.Raw = true
	run.ModifyRepoFlags.Default = true
	run.ModifyRepoFlags.Name = "NewName"
	run.ModifyRepoFlags.Server = "localhost:5000"
	run.ModifyRepoFlags.Username = ""
	run.ModifyRepoFlags.Password = ""
	run.ExistingMRFlags["--default"] = true
	run.ExistingMRFlags["--username"] = false
	run.ExistingMRFlags["--password"] = false

	var buf bytes.Buffer
	stdout := `1||official|https://registry-1.docker.io/|
2|*|NewName|localhost:5000|
`
	db.RDriver.NewRepo(&db.Repo{Id: 2, Default: true, Name: "test", Server: "test.com"})
	run.ModifyRepo("2")
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestModifyRepoUsername(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	run.ListReposFlags.Raw = true
	run.ModifyRepoFlags.Default = true
	run.ModifyRepoFlags.Name = "NewName"
	run.ModifyRepoFlags.Server = "localhost:5000"
	run.ModifyRepoFlags.Username = "user"
	run.ModifyRepoFlags.Password = ""
	run.ExistingMRFlags["--default"] = true
	run.ExistingMRFlags["--username"] = true
	run.ExistingMRFlags["--password"] = false

	var buf bytes.Buffer
	stdout := `1||official|https://registry-1.docker.io/|
2|*|NewName|localhost:5000|user
`
	db.RDriver.NewRepo(&db.Repo{Id: 2, Default: true, Name: "test", Server: "test.com"})
	run.ModifyRepo("2")
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestModifyRepoOneRepoToNotDefault(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	run.ListReposFlags.Raw = true
	run.ModifyRepoFlags.Default = false
	run.ModifyRepoFlags.Name = ""
	run.ModifyRepoFlags.Server = ""
	run.ModifyRepoFlags.Username = ""
	run.ModifyRepoFlags.Password = ""
	run.ExistingMRFlags["--default"] = true
	run.ExistingMRFlags["--username"] = false
	run.ExistingMRFlags["--password"] = false

	var buf bytes.Buffer
	stdout := `1|*|official|https://registry-1.docker.io/|
2||test|localhost:5000|
`
	db.RDriver.NewRepo(&db.Repo{Id: 1, Default: true, Name: "test", Server: "localhost:5000"})
	run.ModifyRepo("2")
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestModifyRepoSecondRepoToNotDefault(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	run.ListReposFlags.Raw = true

	db.RDriver.NewRepo(&db.Repo{Id: 1, Name: "test", Server: "localhost:5000"})
	db.RDriver.NewRepo(&db.Repo{Id: 2, Default: true, Name: "test2", Server: "localhost:5000"})

	run.ModifyRepo("2")

	run.ModifyRepoFlags.Default = false
	run.ModifyRepoFlags.Name = ""
	run.ModifyRepoFlags.Server = ""
	run.ModifyRepoFlags.Username = ""
	run.ModifyRepoFlags.Password = ""
	run.ExistingMRFlags["--default"] = true
	run.ExistingMRFlags["--username"] = false
	run.ExistingMRFlags["--password"] = false
	run.ModifyRepo("3")

	var buf bytes.Buffer
	stdout := `1|*|official|https://registry-1.docker.io/|
2||test|localhost:5000|
3||test2|localhost:5000|
`
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}

func TestModifyRepoOfficialRepoToNotDefault(t *testing.T) {
	setDefaultConfig()
	db.RDriver.ClearRepos()
	run.ListReposFlags.Raw = true

	db.RDriver.NewRepo(&db.Repo{Id: 1, Name: "test", Server: "localhost:5000"})

	run.ModifyRepoFlags.Default = false
	run.ModifyRepoFlags.Name = ""
	run.ModifyRepoFlags.Server = ""
	run.ModifyRepoFlags.Username = ""
	run.ModifyRepoFlags.Password = ""
	run.ExistingMRFlags["--default"] = true
	run.ExistingMRFlags["--username"] = false
	run.ExistingMRFlags["--password"] = false
	run.ModifyRepo("1")

	var buf bytes.Buffer
	stdout := `1|*|official|https://registry-1.docker.io/|
2||test|localhost:5000|
`
	log.SetOutput(&buf)
	run.ListRepos()
	log.SetOutput(os.Stdout)
	printFailInfo(t, stdout, buf.String())

	dropTestDB(t)
}
