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

import "dam/driver/structures"

type RProvider interface {
	GetRepos() []*structures.Repo
	GetRepoById(id int) *structures.Repo
	GetDefaultRepo() *structures.Repo
	NewRepo(repo *structures.Repo)
	ModifyRepo(repo *structures.Repo)
	RemoveRepoById(id int)
	GetRepoIdByName(name *string) int
	ClearRepos()
}

type AProvider interface {
	GetApps() []*structures.App
	NewApp(app *structures.App)
	GetAppById(id int) *structures.App
	ExistFamily(family string) bool
	RemoveAppById(id int)
}