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
	"dam/driver/decorate"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/driver/registry"
	"dam/driver/sort"
)

func Search(arg string) {
	flag.ValidateSearchMask(arg)
	logger.Debug("Flags validated with success")

	logger.Debug("Getting default repo ...")
	repo := db.RDriver.GetDefaultRepo()
	if repo == nil {
		logger.Fatal("Internal error. Not found default repo")
	}
	logger.Debug("Getting registry info ...")
	registry.CheckRepository(repo)
	appList := registry.GetAppNamesByMask(repo, arg)
	logger.Debug("Sorting app list ...")
	appSortedList := sort.SortAppNames(appList)
	logger.Debug("Printing app list ...")
	for _, app := range *appSortedList {
		decorate.PrintSearchedApp(app)
		vers := registry.GetVersions(repo, app)
		versSortedList := sort.SortVersions(vers)
		decorate.PrintSearchedVersions(*versSortedList)
	}
}
