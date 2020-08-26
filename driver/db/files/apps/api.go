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
package apps

import (
	"dam/driver/structures"
)

type provider struct {
	//GetApps() []*storage.App
	//NewApp(app *storage.App)
	//GetAppById(id int) *storage.App
	//ExistFamily(family string) bool
	//RemoveAppById(id int)
}

func NewProvider() *provider {
	return &provider{}
}

func (p *provider) GetApps() []*structures.App {
	return GetApps()
}

func (p *provider) NewApp(app *structures.App) {
	NewApp(app)
}

func (p *provider) GetAppById(id int) *structures.App {
	return GetAppById(id)
}

func (p *provider) ExistFamily(family string) bool {
	return ExistFamily(family)
}

func (p *provider) RemoveAppById(id int) {
	RemoveAppById(id)
}
