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
package run

import (
	"dam/driver/db"
	"dam/driver/logger"
	"dam/driver/storage"
	"dam/driver/validate"
)

type AddRepoSettings struct {
	Name   string
	Server string
	Default bool
	Username string
	Password string
}

var AddRepoFlags = new(AddRepoSettings)

func AddRepo(){
	validate.RepoName(AddRepoFlags.Name)
	validate.RepoServer(AddRepoFlags.Server)
	validate.RepoUsername(AddRepoFlags.Username)
	validate.RepoPassword(AddRepoFlags.Password)

	repo  := new(storage.Repo)
	repo.Default = AddRepoFlags.Default
	repo.Name = AddRepoFlags.Name
	repo.Server = AddRepoFlags.Server
	repo.Username = AddRepoFlags.Username
	repo.Password = AddRepoFlags.Password

	for _, repoDB := range db.RDriver.GetRepos() {
		if repoDB.Name == repo.Name {
			logger.Fatal("Repository name already exist in DB")
		}
	}

	db.RDriver.NewRepo(repo)
}
