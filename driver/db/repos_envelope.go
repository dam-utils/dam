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
	files_db "dam/driver/db/files"
	"dam/driver/storage"
)



func GetRepos() *[]storage.Repo {
	switch config.DB_TYPE {
	case "files":
		return files_db.GetRepos()
	default:
		dbConfigureIsBad()
	}
	return nil
}

func GetRepoById(id int) *storage.Repo {
	switch config.DB_TYPE {
	case "files":
		return files_db.GetRepoById(id)
	default:
		dbConfigureIsBad()
	}
	return nil
}

func GetDefaultRepo() *storage.Repo {
	switch config.DB_TYPE {
	case "files":
		return files_db.GetDefaultRepo()
	default:
		dbConfigureIsBad()
	}
	return nil
}

func NewRepo(repo *storage.Repo) {
	switch config.DB_TYPE {
	case "files":
		files_db.NewRepo(repo)
	default:
		dbConfigureIsBad()
	}
}

func ModifyRepo(repo *storage.Repo) {
	switch config.DB_TYPE {
	case "files":
		files_db.ModifyRepo(repo)
	default:
		dbConfigureIsBad()
	}
}

func RemoveRepoById(id int) {
	switch config.DB_TYPE {
	case "files":
		files_db.RemoveRepoById(id)
	default:
		dbConfigureIsBad()
	}
}

func GetRepoIdByName(name *string) int {
	switch config.DB_TYPE {
	case "files":
		return files_db.GetRepoIdByName(name)
	default:
		dbConfigureIsBad()
	}
	return 0
}

func ClearRepos() {
	switch config.DB_TYPE {
	case "files":
		files_db.ClearRepos()
	default:
		dbConfigureIsBad()
	}
}
