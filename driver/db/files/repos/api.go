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
package repos

import (
	"dam/driver/structures"
)

type provider struct {
	//GetRepos() []*storage.Repo
	//GetRepoById(id int) *storage.Repo
	//GetDefaultRepo() *storage.Repo
	//NewRepo(repo *storage.Repo)
	//ModifyRepo(repo *storage.Repo)
	//RemoveRepoById(id int)
	//GetRepoIdByName(name *string) int
	//ClearRepos()
}

func NewProvider() *provider {
	return &provider{}
}

func (p *provider) GetRepos() []*structures.Repo {
	return GetRepos()
}

func (p *provider) GetRepoById(id int) *structures.Repo {
	return GetRepoById(id)

}

func (p *provider) GetDefaultRepo() *structures.Repo {
	return GetDefaultRepo()
}

func (p *provider) NewRepo(repo *structures.Repo) {
	NewRepo(repo)
}

func (p *provider) ModifyRepo(repo *structures.Repo) {
	ModifyRepo(repo)
}

func (p *provider) RemoveRepoById(id int) {
	RemoveRepoById(id)
}

func (p *provider) GetRepoIdByName(name *string) int {
	return GetRepoIdByName(name)
}

func (p *provider) ClearRepos() {
	ClearRepos()
}