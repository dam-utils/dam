package run

import (
	"dam/driver/conf/option"
	"strconv"

	"dam/driver/db"
	"dam/driver/flag"
	"dam/driver/logger"
)

type RemoveRepoSettings struct {
	Force    bool
}

var RemoveRepoFlags = new(RemoveRepoSettings)

func RemoveRepo(arg string) {
	var repoId int
	if isRepoID(arg) {
		flag.ValidateRepoName(arg)
		repoId = db.RDriver.GetRepoIdByName(&arg)
	} else {
		flag.ValidateRepoID(arg)
		repoId, _ = strconv.Atoi(arg)
	}
	logger.Debug("Flags validated with success")

	logger.Debug("Getting default repo ...")
	defRepo := db.RDriver.GetDefaultRepo()
	if defRepo == nil {
		logger.Fatal("Internal error. Not found default repo")
	}
	if !RemoveRepoFlags.Force && repoId == defRepo.Id {
		logger.Fatal("Repository with Id '%v' is default. Use '--force' flag for removing", repoId)
	}

	logger.Debug("Removing from DB ...")
	db.RDriver.RemoveRepoById(repoId)
	db.ADriver.ChangeRepoID(repoId, option.Config.DefaultRepo.GetUnknownRepoID())
}

func isRepoID(id string) bool {
	_, err := strconv.Atoi(id )

	return err != nil
}