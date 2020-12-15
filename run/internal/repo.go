package internal

import (
	"dam/driver/db"
	"dam/driver/structures"
)

func GetPrefixRepo() string {
	var prefixRepo string
	repo := db.RDriver.GetDefaultRepo()
	if repo.Id != structures.OfficialRepo.Id {
		prefixRepo = repo.Server + "/"
	}
	return prefixRepo
}
