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
	"dam/config"
	"dam/driver/db"
	"dam/driver/decorate"
	"dam/driver/engine"
	"dam/driver/structures"
)

type PurgeSettings struct {
	All bool
}

var PurgeFlags = new(PurgeSettings)

func Purge() {
	allApps := make(map[string]bool)
	if PurgeFlags.All {
		for _, id := range *engine.VDriver.Images() {
			allApps[id] = true
		}
	} else {
		for _, id := range *engine.VDriver.Images() {
			_, allApps[id] = engine.VDriver.GetImageLabel(id, config.APP_FAMILY_ENV)
		}
	}

	for _, a := range db.ADriver.GetApps() {
		allApps[a.DockerID] = false
	}

	var stats structures.Stats
	stats.All = len(allApps)

	for key, value := range allApps {
		if value {
			if engine.VDriver.ImageRemove(key) {
				stats.Deleted++
			} else {
				stats.CanNotDeleted++
			}
		} else {
			stats.Skip++
		}
	}

	decorate.PrintGarbageStatistic(stats)
}
