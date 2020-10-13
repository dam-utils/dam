package run

import (
	"dam/driver/db"
	"dam/driver/flag"
	"dam/driver/logger"
	"dam/driver/structures"
)

type AddRepoSettings struct {
	Name   string
	Server string
	Default bool
	Username string
	Password string
}

var AddRepoFlags = new(AddRepoSettings)

func AddRepo(){
	flag.ValidateRepoName(AddRepoFlags.Name)
	flag.ValidateRepoServer(AddRepoFlags.Server)
	flag.ValidateRepoUsername(AddRepoFlags.Username)
	flag.ValidateRepoPassword(AddRepoFlags.Password)
	logger.Debug("Flags validated with success")

	repo  := new(structures.Repo)
	repo.Default = AddRepoFlags.Default
	repo.Name = AddRepoFlags.Name
	repo.Server = AddRepoFlags.Server
	repo.Username = AddRepoFlags.Username
	repo.Password = AddRepoFlags.Password

	logger.Debug("Starting db.RDriver.GetRepos() ...")
	for _, repoDB := range db.RDriver.GetRepos() {
		if repoDB.Name == repo.Name {
			logger.Fatal("Repository name already exist in DB")
		}
	}

	logger.Debug("Creating new repo ...")
	db.RDriver.NewRepo(repo)
}
