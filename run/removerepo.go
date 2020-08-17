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

type RemoveRepoSettings struct {
	Force    bool
}

var RemoveRepoFlags = new(RemoveRepoSettings)

func RemoveRepo(arg string) {
	var repoId int
	if isRepoID(arg) {
		flag.ValidateRepoName(arg)
		repoId = db.RDriver.GetRepoIdByName(&arg)
	} else {
		flag.ValidateRepoID(arg)
		repoId, _ = strconv.Atoi(arg)
	}
	logger.Debug("Flags validated with success")

	logger.Debug("Getting default repo ...")
	defRepo := db.RDriver.GetDefaultRepo()
	if defRepo == nil {
		logger.Fatal("Internal error. Not found default repo")
	}
	if !RemoveRepoFlags.Force && repoId == defRepo.Id {
		logger.Fatal("Repository with Id '%v' is default. Use '--skip' flag for removing", repoId)
	}

	logger.Debug("Removing from DB ...")
	db.RDriver.RemoveRepoById(repoId)
}

func isRepoID(id string) bool {
	_, err := strconv.Atoi(id )

	return err != nil
}