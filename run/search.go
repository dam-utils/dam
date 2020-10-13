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
