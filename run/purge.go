package run

import (
	"dam/driver/conf/option"
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
			_, allApps[id] = engine.VDriver.GetImageLabel(id, option.Config.ReservedEnvs.GetAppFamilyEnv())
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
