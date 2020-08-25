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
	"dam/driver/db/storage"
	"dam/driver/flag"
	"dam/driver/logger"
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
	flag.ValidateRepoName(AddRepoFlags.Name)
	flag.ValidateRepoServer(AddRepoFlags.Server)
	flag.ValidateRepoUsername(AddRepoFlags.Username)
	flag.ValidateRepoPassword(AddRepoFlags.Password)
	logger.Debug("Flags validated with success")

	repo  := new(storage.Repo)
	repo.Default = AddRepoFlags.Default
	repo.Name = AddRepoFlags.Name
	repo.Server = AddRepoFlags.Server
	repo.Username = AddRepoFlags.Username
	repo.Password = AddRepoFlags.Password

	logger.Debug("Starting db.RDriver.GetRepos() ...")
	for _, repoDB := range db.RDriver.GetRepos() {
		if repoDB.Name == repo.Name {
			logger.Fatal("Repository name already exist in DB")
		}
	}

	logger.Debug("Creating new repo ...")
	db.RDriver.NewRepo(repo)
}
