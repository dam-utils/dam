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
	"dam/driver/flag"
	"dam/driver/logger"
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
	flag.ValidateRepoID(arg)
	if ExistingMRFlags["--name"] {
		flag.ValidateRepoName(ModifyRepoFlags.Name)
	}
	if ExistingMRFlags["--server"] {
		flag.ValidateRepoServer(ModifyRepoFlags.Server)
	}
	if ExistingMRFlags["--username"] {
		flag.ValidateRepoUsername(ModifyRepoFlags.Username)
	}
	if ExistingMRFlags["--password"] {
		flag.ValidateRepoPassword(ModifyRepoFlags.Password)
	}
	logger.Debug("Flags validated with success")

	logger.Debug("Preparing options ...")
	ID, err := strconv.Atoi(arg)
	if err != nil {
		logger.Fatal("Internal error. Command argument is not ID")
	}
	repo := db.RDriver.GetRepoById(ID)
	if repo == nil {
		logger.Fatal("Internal error. Cannot get repo for ID '%v'", ID)
	}
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
				logger.Fatal("Cannot modificate official repository with the flag '%s'. Except flag '--default'", noModFlag)
			}
		}
	}

	logger.Debug("Modifying checked repo ...")
	db.RDriver.ModifyRepo(repo)
}
