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
	"bufio"
	"dam/driver/db/files/apps/internal"
	"dam/driver/logger"
	"dam/driver/structures"
)

func (p *provider) GetApps() []*structures.App {
	var apps []*structures.App

	p.connect()
	fileScanner := bufio.NewScanner(p.client)
	defer p.close()

	for fileScanner.Scan() {
		newLine := fileScanner.Text()
		apps = append(apps, internal.Str2app(newLine))
	}
	return apps
}

func (p *provider) GetAppById(id int) *structures.App {
	apps := p.GetApps()
	for _, app := range apps {
		if app.Id == id {
			return app
		}
	}
	return nil
}

func (p *provider) NewApp(app *structures.App) {
	apps := p.GetApps()
	app.Id = internal.GetNewAppID(apps)

	newApps := append(apps, app)
	internal.SaveApps(newApps)
}

func (p *provider) ExistFamily(family string) bool {
	apps := p.GetApps()
	for _, a := range apps {
		if a.Family == family {
			return true
		}
	}
	return false
}

func (p *provider) RemoveAppById(id int) {
	newApps := make([]*structures.App, 0)

	apps := p.GetApps()
	for _, a := range apps {
		if a.Id != id {
			newApps = append(newApps, a)
		}
	}
	if len(newApps) < len(apps) {
		internal.SaveApps(newApps)
	} else {
		logger.Fatal("Not found Id of the App in DB")
	}
}
