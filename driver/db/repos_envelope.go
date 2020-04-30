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
package db

import (
	"dam/config"
	d_log "dam/decorate/log"
	files_db "dam/driver/db/files"
	"dam/driver/storage"
)

type Provider interface {
	GetRepos() *[]storage.Repo
	GetRepoById(id int) *storage.Repo
	GetDefaultRepo() *storage.Repo
	NewRepo(repo *storage.Repo)
	ModifyRepo(repo *storage.Repo)
	RemoveRepoById(id int)
	GetRepoIdByName(name *string) int
	ClearRepos()
}

var Driver Provider

func Init() {
	switch config.DB_TYPE {
	case "files":
		Driver = files_db.NewProvider()
	default:
		dbConfigureIsBad()
	}
}

func dbConfigureIsBad() {
	d_log.Fatal("Config option UTIL_NAME='" + config.UTIL_NAME + "' not valid. DB type is bad.")
}