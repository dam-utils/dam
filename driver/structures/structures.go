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

var OfficialRepo = Repo{
	Id : 1,
	Default: true,
	Name: config.OFFICIAL_REGISTRY_NAME,
	Server: config.OFFICIAL_REGISTRY_URL,
	Username: "",
	Password: "",
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

