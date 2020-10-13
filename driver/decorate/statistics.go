package decorate

import (
	"fmt"

	"dam/driver/structures"
)

func PrintGarbageStatistic(stats structures.Stats) {
	fmt.Printf(`
    Statistic:

Skip docker images: %v
Removed docker images: %v
Cannot removed docker images: %v
All docker images: %v

`,
	stats.Skip, stats.Deleted, stats.CanNotDeleted, stats.All)
}