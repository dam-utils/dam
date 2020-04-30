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
	d_log "dam/driver/logger"
	"strconv"
)

type RemoveRepoSettings struct {
	Force    bool
}

var RemoveRepoFlags = new(RemoveRepoSettings)

func RemoveRepo(arg string) {
	var repoId int
	repoId, err := strconv.Atoi(arg)
	if err != nil {
		// Maybe It is Name
		repoId = db.Driver.GetRepoIdByName(&arg)
	}
	if repoId == 1 {
		d_log.Fatal("Command argument '"+arg+"' is not Id or Name of Repository")
	}

	defRepo := db.Driver.GetDefaultRepo()
	if !RemoveRepoFlags.Force && repoId == defRepo.Id {
		d_log.Fatal("Repository with Id '" + strconv.Itoa(repoId) + "' is default. Use '--skip' flag for removing.")
	}
	db.Driver.RemoveRepoById(repoId)
}