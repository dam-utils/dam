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
	"strconv"

	"dam/driver/db"
	"dam/driver/logger"
	"dam/driver/validate"
)

type ModifyRepoSettings struct {
	Default bool
	Name   string
	Server string
	Username string
	Password string
}

var ModifyRepoFlags = new(ModifyRepoSettings)

var ExistingMRFlags = make(map[string]bool)

func ModifyRepo(arg string) {
	validate.RepoID(arg)
	validate.RepoName(ModifyRepoFlags.Name)
	validate.RepoServer(ModifyRepoFlags.Server)
	validate.RepoUsername(ModifyRepoFlags.Username)
	validate.RepoPassword(ModifyRepoFlags.Password)

	ID, err := strconv.Atoi(arg)
	if err != nil {
		logger.Fatal("Internal error. Command argument is not ID. See 'help modifyrepo'")
	}
	repo := db.RDriver.GetRepoById(ID)
	if ExistingMRFlags["--default"] && repo.Default != ModifyRepoFlags.Default {
		repo.Default = ModifyRepoFlags.Default
	}
	if ModifyRepoFlags.Name != "" {
		repo.Name = ModifyRepoFlags.Name
	}
	if ModifyRepoFlags.Server != "" {
		repo.Server = ModifyRepoFlags.Server
	}
	if ExistingMRFlags["--username"] && ModifyRepoFlags.Username != "" {
		repo.Username = ModifyRepoFlags.Username
	}
	if ExistingMRFlags["--password"] && ModifyRepoFlags.Password != "" {
		repo.Password = ModifyRepoFlags.Password
	}

	if repo.Id == 1 {
		noModifyOfficialRepoFlags := []string{"--name", "--server", "--username", "--password"}
		for _, noModFlag := range noModifyOfficialRepoFlags {
			if ExistingMRFlags[noModFlag] {
				logger.Fatal("Cannot use flag '%s'for official repository. Except flag '--default'", noModFlag)
			}
		}
	}
	logger.Debug("Repo for modify: '%v'", *repo)
	db.RDriver.ModifyRepo(repo)
}
