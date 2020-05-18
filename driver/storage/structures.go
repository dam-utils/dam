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
package storage

import "dam/config"

type Repo struct {
	Id int
	Default bool
	Name string
	Server string
	Username string
	Password string
}

var OfficialRepo Repo

func init(){
	OfficialRepo.Id = 1
	OfficialRepo.Name = config.OFFICIAL_REGISTRY_NAME
	OfficialRepo.Default = true
	OfficialRepo.Server=config.OFFICIAL_REGISTRY_URL
	OfficialRepo.Username=""
	OfficialRepo.Password=""
}

type App struct {
	Id int
	DockerID string
	ImageName string
	ImageVersion string
	RepoID int
	MultiVersion bool
	Family string
}