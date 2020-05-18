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
	"dam/driver/storage/validate"
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
	// Convert pass to base64
	repo  := new(storage.Repo)
	repo.Default = AddRepoFlags.Default

	err := validate.CheckRepoName(AddRepoFlags.Name,)
	if err != nil {
		logger.Fatal(err.Error())
	} else {
		repo.Name = AddRepoFlags.Name
	}

	repos :=  db.RDriver.GetRepos()
	for _, repoDB := range *repos {
		if repoDB.Name == repo.Name {
			logger.Fatal("Repository name already exist in DB")
		}
	}

	err = validate.CheckServer(AddRepoFlags.Server)
	if err != nil {
		logger.Fatal(err.Error())
	} else {
		repo.Server = AddRepoFlags.Server
	}
	err = validate.CheckLogin(AddRepoFlags.Username)
	if err != nil {
		logger.Fatal(err.Error())
	} else {
		repo.Username = AddRepoFlags.Username
	}
	repo.Username = AddRepoFlags.Username
	repo.Password = AddRepoFlags.Password

	db.RDriver.NewRepo(repo)
}
