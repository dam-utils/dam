package run

import (
	"dam/driver/conf/option"
	"dam/driver/db"
	"dam/driver/decorate"
	"dam/driver/engine"
	"dam/driver/logger"
	"dam/driver/structures"
)

type PurgeSettings struct {
	All bool
}

var PurgeFlags = new(PurgeSettings)

func Purge() {
	allApps := make(map[string]bool)
	if PurgeFlags.All {
		logger.Debug("Finding all apps ... ")
		for _, id := range *engine.VDriver.Images() {
			allApps[id] = true
		}
	} else {
		logger.Debug("Finding all apps with label FAMILY ... ")
		for _, id := range *engine.VDriver.Images() {
			logger.Debug("Check image id with family '%s'", id)
			_, allApps[id] = engine.VDriver.GetImageLabel(id, option.Config.ReservedEnvs.GetAppFamilyEnv())
		}
	}

	logger.Debug("All candidates for removing: '%v'", allApps)

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
