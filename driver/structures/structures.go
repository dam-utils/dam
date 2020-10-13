package structures

import "dam/config"

type Repo struct {
	Id int
	Default bool
	Name string
	Server string
	Username string
	Password string
}

var OfficialRepo Repo

func init(){
	OfficialRepo.Id = 1
	OfficialRepo.Name = config.OFFICIAL_REGISTRY_NAME
	OfficialRepo.Default = true
	OfficialRepo.Server=config.OFFICIAL_REGISTRY_URL
	OfficialRepo.Username=""
	OfficialRepo.Password=""
}

type App struct {
	Id int
	DockerID string
	ImageName string
	ImageVersion string
	RepoID int
	MultiVersion bool
	Family string
}

